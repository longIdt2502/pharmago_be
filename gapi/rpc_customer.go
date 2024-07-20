package gapi

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/longIdt2502/pharmago_be/common"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/utils"
	"github.com/thoas/go-funk"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PayloadZaloOAS struct {
	Page   int      `json:"page"`
	Limit  int      `json:"limit"`
	Phones []string `json:"phones"`
}

type ResponseZaloOAS struct {
	Data ResponseDataZaloOAS `json:"data"`
}

type ResponseDataZaloOAS struct {
	Items []pb.Conversation `json:"items"`
}

func formatPhoneNumber(phone string) string {
	if strings.HasPrefix(phone, "0") {
		return "84" + phone[1:]
	}
	return phone
}

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

	client := &http.Client{}

	url := fmt.Sprintf("%s/v1/zalo-oas/6/", server.config.WezoloServerAdress)

	phones := funk.Map(customers, func(value db.ListCustomerRow) string {
		return formatPhoneNumber(value.Phone.String)
	})

	payload := PayloadZaloOAS{
		Page:   1,
		Limit:  10,
		Phones: phones.([]string),
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	reqHttp, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln("Error creating HTTP request:", err))
	}

	// Thiết lập các header
	reqHttp.Header.Set("Content-Type", "application/json")
	reqHttp.Header.Set("Authorization", "Token 1e21c13f941d67507d9d1099150866b6759d9336")

	// Gửi request và nhận response
	resp, err := client.Do(reqHttp)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln("Error sending HTTP request:", err))
	}

	// Đảm bảo rằng response body sẽ được đóng sau khi hoàn tất
	defer resp.Body.Close()

	// Đọc response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln("Error reading response body:", err))
	}

	var responseData ResponseZaloOAS
	if err := json.Unmarshal(body, &responseData); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln("Lỗi giải mã JSON:", err))
	}

	var customersPb []*pb.Customer
	for _, value := range customers {
		conversation := &pb.Conversation{}
		for index := range responseData.Data.Items {
			if responseData.Data.Items[index].GetPhone() == formatPhoneNumber(value.Phone.String) {
				conversation = &responseData.Data.Items[index]
			}
		}
		dataPb := &pb.Customer{
			Id:           value.ID,
			Code:         value.Code,
			FullName:     value.FullName,
			Company:      value.Company,
			Phone:        value.Phone.String,
			Email:        &value.Email.String,
			Revenue:      float32(value.TotalRevenue.Float64),
			Orders:       value.TotalOrders.Int32,
			Conversation: conversation,
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

	var contactAddressId int32
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

		contactAddressId = address.ID
	}

	customer, err := server.store.CreateCustomer(ctx, db.CreateCustomerParams{
		FullName: req.GetName(),
		Code:     code,
		Company:  req.GetCompany(),
		Address: sql.NullInt32{
			Int32: addressId,
			Valid: req.Address != nil,
		},
		Email: sql.NullString{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
		Phone: sql.NullString{
			String: req.GetPhone(),
			Valid:  true,
		},
		License: sql.NullString{
			String: req.GetLicense(),
			Valid:  req.License != nil,
		},
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
		IssuedBy: sql.NullString{
			String: req.GetIssuedBy(),
			Valid:  req.IssuedBy != nil,
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
			Int32: contactAddressId,
			Valid: req.ContactAddress != nil,
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
			Code:         int32(errLog.StatusCode),
			Message:      errLog.Message,
			MessageTrans: "Lỗi tạo khách hàng",
			Log:          errLog.Log,
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

func (server *ServerGRPC) MedicalRecordCreate(ctx context.Context, req *pb.MedicalRecordCreateRequest) (*pb.MedicalRecordCreateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	var asUuid uuid.UUID
	if req.AppointmentSchedule != nil {
		asUuid, err = uuid.Parse(req.GetAppointmentSchedule())
		if err != nil {
			errApp := common.ErrInternalWithMsg(err, "lịch hẹn không tồn tại")
			return &pb.MedicalRecordCreateResponse{
				Code:         int32(errApp.StatusCode),
				Message:      errApp.Message,
				MessageTrans: errApp.MessageTrans,
				Log:          errApp.Log,
			}, nil
		}
	}

	var res *pb.MedicalRecordCreateResponse

	for _, item := range req.GetFiles() {
		file, _ := utils.NewFileFromFile(item.GetFile(), item.GetName())
		_, err = server.b2Bucket.UploadFile(file.Name, file.Meta, file.File)
		if err != nil {
			errApp := common.ErrInternal(err)
			return &pb.MedicalRecordCreateResponse{
				Code:         int32(errApp.StatusCode),
				Message:      errApp.Message,
				MessageTrans: errApp.MessageTrans,
				Log:          errApp.Log,
			}, nil
		}
		title := file.Name
		url, _ := server.b2Bucket.FileURL(file.Name)

		record, err := server.store.CreateMedicalRecordLink(ctx, db.CreateMedicalRecordLinkParams{
			Uuid:                uuid.New(),
			Type:                db.MedicalRecordLinkType(req.GetType().String()),
			Title:               sql.NullString{String: title, Valid: true},
			Url:                 url,
			Customer:            sql.NullInt32{Int32: req.GetCustomer(), Valid: true},
			AppointmentSchedule: uuid.NullUUID{UUID: asUuid, Valid: req.AppointmentSchedule != nil},
			UserCreated:         sql.NullInt32{Int32: tokenPayload.UserID, Valid: true},
		})
		if err != nil {
			errApp := common.ErrDBWithMsg(err, "Tải lên file thất bại")
			res = &pb.MedicalRecordCreateResponse{
				Code:         int32(errApp.StatusCode),
				Message:      errApp.Message,
				MessageTrans: errApp.MessageTrans,
				Log:          errApp.Log,
			}
		} else {
			res = &pb.MedicalRecordCreateResponse{
				Code:    200,
				Message: "success",
				Details: record.ID,
			}
		}
	}

	return res, nil
}

func (server *ServerGRPC) MedicalRecordCreateStream(req *pb.MedicalRecordCreateRequest, stream pb.Pharmago_MedicalRecordCreateStreamServer) error {
	ctx := stream.Context()

	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return config.UnauthenticatedError(err)
	}

	var asUuid uuid.UUID
	if req.AppointmentSchedule != nil {
		asUuid, err = uuid.Parse(req.GetAppointmentSchedule())
		if err != nil {
			errApp := common.ErrInternalWithMsg(err, "lịch hẹn không tồn tại")
			return errApp
		}
	}

	for _, item := range req.GetFiles() {
		file, _ := utils.NewFileFromFile(item.GetFile(), item.GetName())
		_, err = server.b2Bucket.UploadFile(file.Name, file.Meta, file.File)
		if err != nil {
			return common.ErrInternal(err)
		}
		title := file.Name
		url, _ := server.b2Bucket.FileURL(file.Name)

		var res *pb.MedicalRecordCreateResponse
		record, err := server.store.CreateMedicalRecordLink(ctx, db.CreateMedicalRecordLinkParams{
			Uuid:                uuid.New(),
			Type:                db.MedicalRecordLinkType(req.GetType().String()),
			Title:               sql.NullString{String: title, Valid: true},
			Url:                 url,
			Customer:            sql.NullInt32{Int32: req.GetCustomer(), Valid: true},
			AppointmentSchedule: uuid.NullUUID{UUID: asUuid, Valid: req.AppointmentSchedule != nil},
			UserCreated:         sql.NullInt32{Int32: tokenPayload.UserID, Valid: true},
		})
		if err != nil {
			errApp := common.ErrDBWithMsg(err, "Tải lên file thất bại")
			res = &pb.MedicalRecordCreateResponse{
				Code:         int32(errApp.StatusCode),
				Message:      errApp.Message,
				MessageTrans: errApp.MessageTrans,
				Log:          errApp.Log,
			}
		} else {
			res = &pb.MedicalRecordCreateResponse{
				Code:    200,
				Message: "success",
				Details: record.ID,
			}
		}
		stream.Send(res)
	}

	stream.Send(&pb.MedicalRecordCreateResponse{
		Code:    200,
		Message: "stream end",
	})

	return nil
}

func (server *ServerGRPC) MedicalRecordList(ctx context.Context, req *pb.MedicalRecordListRequest) (*pb.MedicalRecordListResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	var asUuid uuid.UUID
	if req.AppointmentSchedule != nil {
		asUuid, err = uuid.Parse(req.GetAppointmentSchedule())
		if err != nil {
			errApp := common.ErrInternalWithMsg(err, "lịch hẹn không tồn tại")
			return &pb.MedicalRecordListResponse{
				Code:         int32(errApp.StatusCode),
				Message:      errApp.Message,
				MessageTrans: errApp.MessageTrans,
				Log:          errApp.Log,
			}, nil
		}
	}

	medicalRecords, err := server.store.ListMedicalRecordLink(ctx, db.ListMedicalRecordLinkParams{
		Customer: sql.NullInt32{Int32: req.GetCustomer(), Valid: req.Customer != nil},
		TypeMrl: db.NullMedicalRecordLinkType{
			MedicalRecordLinkType: db.MedicalRecordLinkType(req.GetType().String()),
			Valid:                 req.Type != nil,
		},
		Schedule: uuid.NullUUID{UUID: asUuid, Valid: req.AppointmentSchedule != nil},
	})
	if err != nil {
		errApp := common.ErrDBWithMsg(err, "Dữ liệu tài liệu lỗi")
		return &pb.MedicalRecordListResponse{
			Code:         int32(errApp.StatusCode),
			Message:      errApp.Message,
			MessageTrans: errApp.MessageTrans,
			Log:          errApp.Log,
		}, nil
	}

	var medicalRecordsPb []*pb.MedicalRecordLink
	for _, item := range medicalRecords {
		var as_uuid *string
		if item.AppointmentSchedule.Valid {
			s := item.AppointmentSchedule.UUID.String()
			as_uuid = &s
		}
		medicalRecordsPb = append(medicalRecordsPb, &pb.MedicalRecordLink{
			Id:                  item.ID,
			Uuid:                item.Uuid.String(),
			Type:                convertDBTypeToPBType(item.Type),
			Title:               item.Title.String,
			Url:                 item.Url,
			Customer:            item.Customer.Int32,
			AppointmentSchedule: as_uuid,
			UserCreated:         item.UserCreated.Int32,
			CreatedAt:           timestamppb.New(item.CreatedAt),
		})
	}

	return &pb.MedicalRecordListResponse{
		Code:    200,
		Message: "success",
		Details: medicalRecordsPb,
	}, nil
}

func convertDBTypeToPBType(dbType db.MedicalRecordLinkType) pb.MedicalRecordType {
	switch dbType {
	case db.MedicalRecordLinkTypeTest:
		return pb.MedicalRecordType_test
	case db.MedicalRecordLinkTypePatient:
		return pb.MedicalRecordType_patient
	case db.MedicalRecordLinkTypeDiagnostic:
		return pb.MedicalRecordType_diagnostic
	default:
		return pb.MedicalRecordType_test
	}
}

func (server *ServerGRPC) CreateMannyMediaRecord(ctx context.Context, list []*pb.FileItem, typeMR string, account, customer int32, as_uuid *uuid.UUID) ([]db.MedicalRecordLink, error) {
	var results []db.MedicalRecordLink

	for _, item := range list {
		file, _ := utils.NewFileFromFile(item.GetFile(), item.GetName())
		_, err := server.b2Bucket.UploadFile(file.Name, file.Meta, file.File)
		if err != nil {
			return nil, common.ErrInternal(err)
		}
		title := file.Name
		url, _ := server.b2Bucket.FileURL(file.Name)

		record, err := server.store.CreateMedicalRecordLink(ctx, db.CreateMedicalRecordLinkParams{
			Uuid:                uuid.New(),
			Type:                db.MedicalRecordLinkType(typeMR),
			Title:               sql.NullString{String: title, Valid: true},
			Url:                 url,
			Customer:            sql.NullInt32{Int32: customer, Valid: true},
			AppointmentSchedule: uuid.NullUUID{UUID: *as_uuid, Valid: as_uuid != nil},
			UserCreated:         sql.NullInt32{Int32: account, Valid: true},
		})
		if err != nil {
			return nil, common.ErrDBWithMsg(err, "Tải lên file thất bại")
		} else {
			results = append(results, record)
		}
	}

	return results, nil
}
