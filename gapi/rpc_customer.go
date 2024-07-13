package gapi

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/longIdt2502/pharmago_be/common"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *ServerGRPC) CustomerList(ctx context.Context, req *pb.CustomerListRequest) (*pb.CustomerListResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	customers, err := server.store.ListCustomer(ctx, db.ListCustomerParams{
		Company: req.GetCompany(),
		Search: sql.NullString{
			String: req.GetSearch(),
			Valid:  req.Search != nil,
		},
		Page: sql.NullInt32{
			Int32: req.GetPage(),
			Valid: req.Page != nil,
		},
		Limit: sql.NullInt32{
			Int32: req.GetLimit(),
			Valid: req.Limit != nil,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get customer: %e", err)
	}

	var customersPb []*pb.Customer
	for _, value := range customers {
		dataPb := &pb.Customer{
			Id:       value.ID.Int32,
			Code:     value.Code.String,
			FullName: value.FullName.String,
			Company:  value.Company.Int32,
			Phone:    value.Phone.String,
			Email:    &value.Email.String,
			Revenue:  float32(value.TotalRevenue),
			Orders:   value.TotalOrders,
		}
		customersPb = append(customersPb, dataPb)
	}

	count, err := server.store.CountCustomer(ctx, req.GetCompany())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to count customer: %e", err)
	}

	return &pb.CustomerListResponse{
		Code:    200,
		Message: "success",
		Details: customersPb,
		Count:   int32(count),
	}, nil
}

func (server *ServerGRPC) CustomerCreate(ctx context.Context, req *pb.CustomerCreateRequest) (*pb.CustomerCreateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	code := fmt.Sprintf("CUSTOMER-%s-%d", utils.RandomString(6), utils.RandomInt(100, 999))
	if req.Code != nil {
		code = req.GetCode()
	}

	var addressId int32
	if req.Address != nil {
		address, err := server.store.CreateAddress(ctx, db.CreateAddressParams{
			Lat: float64(req.Address.GetLat()),
			Lng: float64(req.Address.GetLng()),
			Province: sql.NullString{
				String: req.Address.GetProvince(),
				Valid:  true,
			},
			District: sql.NullString{
				String: req.Address.GetDistrict(),
				Valid:  true,
			},
			Ward: sql.NullString{
				String: req.Address.GetWard(),
				Valid:  req.Address.Ward != nil,
			},
			Title:       req.Address.GetTitle(),
			UserCreated: tokenPayload.UserID,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to record address: %e", err)
		}

		addressId = address.ID
	}

	var contactAddressId *int32
	if req.ContactAddress != nil {
		address, err := server.store.CreateAddress(ctx, db.CreateAddressParams{
			Lat: float64(req.ContactAddress.GetLat()),
			Lng: float64(req.ContactAddress.GetLng()),
			Province: sql.NullString{
				String: req.ContactAddress.GetProvince(),
				Valid:  true,
			},
			District: sql.NullString{
				String: req.ContactAddress.GetDistrict(),
				Valid:  true,
			},
			Ward: sql.NullString{
				String: req.ContactAddress.GetWard(),
				Valid:  req.ContactAddress.Ward != nil,
			},
			Title:       req.ContactAddress.GetTitle(),
			UserCreated: tokenPayload.UserID,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to record address: %e", err)
		}

		contactAddressId = &address.ID
	}

	customer, err := server.store.CreateCustomer(ctx, db.CreateCustomerParams{
		FullName: req.GetName(),
		Code:     code,
		Company:  req.GetCompany(),
		Address: sql.NullInt32{
			Int32: addressId,
			Valid: req.Address != nil,
		},
		Email: sql.NullString{},
		Phone: sql.NullString{
			String: req.GetPhone(),
			Valid:  true,
		},
		License: sql.NullString{},
		Birthday: sql.NullTime{
			Time:  time.Unix(req.GetBirthday().GetSeconds(), 0),
			Valid: req.Birthday != nil,
		},
		UserUpdated: sql.NullInt32{
			Int32: tokenPayload.UserID,
			Valid: true,
		},
		UserCreated: tokenPayload.UserID,
		Group: sql.NullInt32{
			Int32: req.GetGroup(),
			Valid: req.Group != nil,
		},
		Title: sql.NullString{
			String: req.GetTitle(),
			Valid:  req.Title != nil,
		},
		LicenseDate: sql.NullTime{
			Time:  time.Unix(req.GetLicenseDate().GetSeconds(), 0),
			Valid: req.Title != nil,
		},
		ContactName: sql.NullString{
			String: req.GetContactName(),
			Valid:  req.ContactName != nil,
		},
		ContactTitle: sql.NullString{
			String: req.GetContactTitle(),
			Valid:  req.ContactTitle != nil,
		},
		ContactPhone: sql.NullString{
			String: req.GetContactPhone(),
			Valid:  req.ContactPhone != nil,
		},
		ContactEmail: sql.NullString{
			String: req.GetContactEmail(),
			Valid:  req.ContactEmail != nil,
		},
		ContactAddress: sql.NullInt32{
			Int32: *contactAddressId,
			Valid: contactAddressId != nil,
		},
		AccountNumber: sql.NullString{
			String: req.GetAccountNumber(),
			Valid:  req.AccountNumber != nil,
		},
		BankName: sql.NullString{
			String: req.GetBankName(),
			Valid:  req.BankName != nil,
		},
		BankBranch: sql.NullString{
			String: req.GetBankBranch(),
			Valid:  req.BankBranch != nil,
		},
	})
	if err != nil {
		errLog := common.ErrDB(err)
		return &pb.CustomerCreateResponse{
			Code:    int32(errLog.StatusCode),
			Message: errLog.Message,
		}, nil
	}

	return &pb.CustomerCreateResponse{
		Code:    200,
		Message: "success",
		Details: customer.ID,
	}, nil
}

func (server *ServerGRPC) CustomerDetail(ctx context.Context, req *pb.CustomerDetailRequest) (*pb.CustomerDetailResponse, error) {
	customer, err := server.store.DetailCustomer(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "customer not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get customer: %e", err)
	}

	customerPb, err := mapper.CustomerDetailMapper(ctx, server.store, customer)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to mapper customer: %e", err)
	}

	return &pb.CustomerDetailResponse{
		Code:    200,
		Message: "success",
		Details: customerPb,
	}, nil
}

func (server *ServerGRPC) CustomerUpdate(ctx context.Context, req *pb.CustomerUpdateRequest) (*pb.CustomerUpdateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	customerDb, err := server.store.DetailCustomer(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "customer not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get customer: %e", err)
	}

	if req.Address != nil {
		if customerDb.Address.Valid {
			_, err = server.store.UpdateAddress(ctx, db.UpdateAddressParams{
				Lat: sql.NullFloat64{
					Float64: float64(req.Address.GetLat()),
					Valid:   true,
				},
				Lng: sql.NullFloat64{
					Float64: float64(req.Address.GetLng()),
					Valid:   true,
				},
				Province: sql.NullString{
					String: req.Address.GetProvince(),
					Valid:  true,
				},
				District: sql.NullString{
					String: req.Address.GetDistrict(),
					Valid:  true,
				},
				Ward: sql.NullString{
					String: req.Address.GetWard(),
					Valid:  req.Address.Ward != nil,
				},
				Title: sql.NullString{
					String: req.Address.GetTitle(),
					Valid:  true,
				},
				ID: customerDb.Address.Int32,
			})
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to update address: %e", err)
			}
		}
		// else {
		// 	// TODO: create new address
		// }
	}

	if req.ContactAddress != nil {
		if customerDb.ContactAddress.Valid {
			_, err = server.store.UpdateAddress(ctx, db.UpdateAddressParams{
				Lat: sql.NullFloat64{
					Float64: float64(req.ContactAddress.GetLat()),
					Valid:   true,
				},
				Lng: sql.NullFloat64{
					Float64: float64(req.ContactAddress.GetLng()),
					Valid:   true,
				},
				Province: sql.NullString{
					String: req.ContactAddress.GetProvince(),
					Valid:  true,
				},
				District: sql.NullString{
					String: req.ContactAddress.GetDistrict(),
					Valid:  true,
				},
				Ward: sql.NullString{
					String: req.ContactAddress.GetWard(),
					Valid:  req.ContactAddress.Ward != nil,
				},
				Title: sql.NullString{
					String: req.ContactAddress.GetTitle(),
					Valid:  true,
				},
				ID: customerDb.ContactAddress.Int32,
			})
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to update contact address address: %e", err)
			}
		}
		// else {
		// 	// TODO: create new address
		// }
	}

	_, err = server.store.UpdateCustomer(ctx, db.UpdateCustomerParams{
		FullName: sql.NullString{
			String: req.GetName(),
			Valid:  true,
		},
		Code: sql.NullString{
			String: req.GetCode(),
			Valid:  true,
		},
		Email: sql.NullString{},
		Phone: sql.NullString{
			String: req.GetPhone(),
			Valid:  true,
		},
		License: sql.NullString{},
		Birthday: sql.NullTime{
			Time:  time.Unix(req.GetBirthday().GetSeconds(), 0),
			Valid: req.Birthday != nil,
		},
		UserUpdated: sql.NullInt32{
			Int32: tokenPayload.UserID,
			Valid: true,
		},
		ID: req.GetId(),
		Title: sql.NullString{
			String: req.GetTitle(),
			Valid:  req.Title != nil,
		},
		LicenseDate: sql.NullTime{
			Time:  time.Unix(req.GetLicenseDate().GetSeconds(), 0),
			Valid: req.Title != nil,
		},
		ContactName: sql.NullString{
			String: req.GetContactName(),
			Valid:  req.ContactName != nil,
		},
		ContactTitle: sql.NullString{
			String: req.GetContactTitle(),
			Valid:  req.ContactTitle != nil,
		},
		ContactPhone: sql.NullString{
			String: req.GetContactPhone(),
			Valid:  req.ContactPhone != nil,
		},
		ContactEmail: sql.NullString{
			String: req.GetContactEmail(),
			Valid:  req.ContactEmail != nil,
		},
		AccountNumber: sql.NullString{
			String: req.GetAccountNumber(),
			Valid:  req.AccountNumber != nil,
		},
		BankName: sql.NullString{
			String: req.GetBankName(),
			Valid:  req.BankName != nil,
		},
		BankBranch: sql.NullString{
			String: req.GetBankBranch(),
			Valid:  req.BankBranch != nil,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update customer: %e", err)
	}

	return &pb.CustomerUpdateResponse{
		Code:    200,
		Message: "success",
	}, nil
}

func (server *ServerGRPC) CustomerGroupList(ctx context.Context, req *pb.CustomerGroupListRequest) (*pb.CustomerGroupListResponse, error) {
	customerGroup, err := server.store.ListCustomerGroup(ctx, db.ListCustomerGroupParams{
		Search: sql.NullString{
			String: req.GetSearch(),
			Valid:  req.Search != nil,
		},
		Page: sql.NullInt32{
			Int32: req.GetPage(),
			Valid: req.Page != nil,
		},
		Limit: sql.NullInt32{
			Int32: req.GetLimit(),
			Valid: req.Limit != nil,
		},
		Company: req.GetCompany(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get list customer group: %e", err))
	}

	var customerGroupPb []*pb.SimpleData
	for _, value := range customerGroup {
		var userCreatedName *string
		if value.FullName.Valid {
			name := value.FullName.String
			userCreatedName = &name
		}

		var description *string
		if value.Note.Valid {
			data := value.Note.String
			description = &data
		}

		dataPb := &pb.SimpleData{
			Id:              value.ID,
			Name:            value.Name,
			Code:            value.Code,
			Description:     description,
			UserCreatedName: userCreatedName,
			CreatedAt:       timestamppb.New(value.CreatedAt),
		}
		customerGroupPb = append(customerGroupPb, dataPb)
	}

	return &pb.CustomerGroupListResponse{
		Code:    200,
		Message: "success",
		Details: customerGroupPb,
	}, nil
}

func (server *ServerGRPC) CustomerGroupCreate(ctx context.Context, req *pb.CustomerGroupCreateRequest) (*pb.CustomerGroupCreateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	code := fmt.Sprintf("CG-%s-%d", utils.RandomString(3), utils.RandomInt(100, 999))
	if req.Code != nil {
		code = req.GetCode()
	}
	data, err := server.store.CreateCustomerGroup(ctx, db.CreateCustomerGroupParams{
		Code: code,
		Name: req.GetName(),
		Note: sql.NullString{
			String: req.GetNote(),
			Valid:  req.Note != nil,
		},
		Company:     req.GetCompany(),
		UserCreated: tokenPayload.UserID,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to record customer group: %e", err)
	}

	for _, id := range req.GetCustomers() {
		server.store.UpdateCustomer(ctx, db.UpdateCustomerParams{
			ID: id,
			Group: sql.NullInt32{
				Int32: data.ID,
				Valid: true,
			},
		})
	}

	return &pb.CustomerGroupCreateResponse{
		Code:    200,
		Message: "success",
		Details: data.ID,
	}, nil
}

func (server *ServerGRPC) CustomerGroupDetail(ctx context.Context, req *pb.CustomerGroupDetailRequest) (*pb.CustomerGroupDetailResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	data, err := server.store.DetailCustomerGroup(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get detail customer group: %e", err)
	}

	var userCreatedName *string
	if data.FullName.Valid {
		name := data.FullName.String
		userCreatedName = &name
	}

	var userUpdatedName *string
	if data.FullName_2.Valid {
		nameUd := data.FullName_2.String
		userUpdatedName = &nameUd
	}

	var description *string
	if data.Note.Valid {
		data := data.Note.String
		description = &data
	}

	return &pb.CustomerGroupDetailResponse{
		Code:    200,
		Message: "success",
		Details: &pb.SimpleData{
			Id:              data.ID,
			Name:            data.Name,
			Code:            data.Code,
			UserCreatedName: userCreatedName,
			CreatedAt:       timestamppb.New(data.CreatedAt),
			UserUpdatedName: userUpdatedName,
			UpdatedAt:       timestamppb.New(data.UpdatedAt.Time),
			ValueExtra:      nil,
			Description:     description,
		},
	}, nil
}

func (server *ServerGRPC) CustomerGroupUpdate(ctx context.Context, req *pb.CustomerGroupUpdateRequest) (*pb.CustomerGroupUpdateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	data, err := server.store.UpdateCustomerGroup(ctx, db.UpdateCustomerGroupParams{
		Name: sql.NullString{
			String: req.GetName(),
			Valid:  true,
		},
		Code: sql.NullString{
			String: req.GetCode(),
			Valid:  req.Code != nil,
		},
		Note: sql.NullString{
			String: req.GetNote(),
			Valid:  req.Note != nil,
		},
		ID: req.GetId(),
		UserUpdated: sql.NullInt32{
			Int32: tokenPayload.UserID,
			Valid: true,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update customer group: %e", err)
	}

	return &pb.CustomerGroupUpdateResponse{
		Code:    200,
		Message: "success",
		Details: data.ID,
	}, nil
}

func (server *ServerGRPC) CustomerGroupDelete(ctx context.Context, req *pb.CustomerGroupDeleteRequest) (*pb.CustomerGroupDeleteResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	_, err = server.store.DeleteCustomerGroup(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "customer group not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to delete customer group: %e", err)
	}

	return &pb.CustomerGroupDeleteResponse{
		Code:    200,
		Message: "success",
	}, nil
}
