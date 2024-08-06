package gapi

import (
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
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *ServerGRPC) AccountDetail(ctx context.Context, _ *pb.AccountDetailRequest) (*pb.AccountDetailResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	account, err := server.store.GetAccount(ctx, tokenPayload.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "account not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get account: %e", err)
	}

	accountPb := mapper.AccountRowMapper(account)

	companies, _ := server.store.GetCompanies(ctx, db.GetCompaniesParams{
		Owner: sql.NullInt32{
			Int32: account.ID,
			Valid: true,
		},
	})

	var companiesPb []*pb.Company
	for _, value := range companies {
		dataPb := mapper.CompanyMapper(ctx, server.store, value)
		companiesPb = append(companiesPb, dataPb)
	}

	return &pb.AccountDetailResponse{
		Code:    200,
		Message: "success",
		Details: &pb.AccountDetailResponseDetail{
			Account: accountPb,
			Company: companiesPb,
		},
	}, nil
}

func (server *ServerGRPC) AccountInactive(ctx context.Context, req *pb.AccountInactiveRequest) (*pb.AccountInactiveResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	var accountId int32

	if req.Id != nil {
		account, err := server.store.GetAccount(ctx, req.GetId())
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, status.Errorf(codes.NotFound, "account not exists")
			}
			return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to get account: %v", err))
		}
		accountId = account.ID

		if account.Type == 3 && req.GetId() != tokenPayload.UserID {
			return nil, status.Errorf(codes.PermissionDenied, "permission denied")
		}
	} else {
		accountId = tokenPayload.UserID
	}

	_, err = server.store.UpdateAccount(ctx, db.UpdateAccountParams{
		ID: sql.NullInt32{
			Int32: accountId,
			Valid: true,
		},
		IsVerify: sql.NullBool{
			Bool:  req.GetStatus(),
			Valid: true,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to inactive user: %v", err))
	}

	return &pb.AccountInactiveResponse{
		Code:    200,
		Message: "success",
	}, nil
}

func (server *ServerGRPC) AccountList(ctx context.Context, req *pb.AccountListRequest) (*pb.AccountListResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	accountsDb, err := server.store.ListAccount(ctx, db.ListAccountParams{
		Company:  sql.NullInt32{Int32: req.GetCompany(), Valid: true},
		Search:   sql.NullString{String: req.GetSearch(), Valid: true},
		IsVerify: sql.NullBool{Bool: req.GetActive(), Valid: req.Active != nil},
		Type:     sql.NullInt32{Int32: req.GetType(), Valid: req.Type != nil},
		Role:     sql.NullInt32{Int32: req.GetRole(), Valid: req.Role != nil},
		Page:     sql.NullInt32{Int32: req.GetPage(), Valid: req.Page != nil},
		Limit:    sql.NullInt32{Int32: req.GetLimit(), Valid: req.Limit != nil},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get list account: %e", err)
	}

	var accountsPb []*pb.Account
	for _, item := range accountsDb {
		accountsPb = append(accountsPb, mapper.ListAccountRowMapper(item))
	}

	counts, err := server.store.CountAccountByStatus(ctx, req.Company)
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.AccountListResponse{
			Code:         int32(errApp.StatusCode),
			Message:      errApp.Message,
			MessageTrans: errApp.MessageTrans,
			Log:          errApp.Log,
		}, nil
	}

	countsPb := pb.AccountListCount{Active: 0, UnActive: 0}
	for _, item := range counts {
		if item.IsVerify {
			countsPb.Active = int32(item.Count)
		} else {
			countsPb.UnActive = int32(item.Count)
		}
	}

	return &pb.AccountListResponse{
		Code:    200,
		Message: "success",
		Details: accountsPb,
		Counts:  &countsPb,
	}, nil
}

func (server *ServerGRPC) CreateEmployee(ctx context.Context, req *pb.CreateEmployeeRequest) (*pb.CreateEmployeeResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	hashPass, err := utils.HashedPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password")
	}

	var address db.Address
	if req.Address != nil {
		address, err = server.store.CreateAddress(ctx, db.CreateAddressParams{
			Lat: float64(req.Address.Lat),
			Lng: float64(req.Address.Lng),
			Province: sql.NullString{
				String: req.Address.Province.Code,
				Valid:  true,
			},
			Ward: sql.NullString{
				String: req.Address.Ward.Code,
				Valid:  true,
			},
			District: sql.NullString{
				String: req.Address.District.Code,
				Valid:  true,
			},
			Title:       req.Address.Title,
			UserCreated: tokenPayload.UserID,
		})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to create address")
		}
	}

	accountType, err := server.store.GetAccountType(ctx, db.GetAccountTypeParams{
		ID: sql.NullInt32{},
		Code: sql.NullString{
			String: req.AccountType,
			Valid:  true,
		},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get account type")
	}

	employee, err := server.store.CreateAccount(ctx, db.CreateAccountParams{
		Username:       req.GetUsername(),
		HashedPassword: hashPass,
		FullName:       req.FullName,
		Email:          req.Email,
		Type:           accountType.ID,
		Role: sql.NullInt32{
			Int32: req.GetRole(),
			Valid: req.Role != nil,
		},
		Gender: db.NullGender{
			Gender: db.GenderNam,
			Valid:  req.Gender != nil,
		},
		Licence: sql.NullString{
			String: req.GetLicence(),
			Valid:  req.Licence != nil,
		},
		Dob: sql.NullTime{
			Time:  time.Unix(req.GetDob().GetSeconds(), 0),
			Valid: req.Dob.IsValid(),
		},
		Address: sql.NullInt32{
			Int32: address.ID,
			Valid: req.Address != nil,
		},
		IsVerify: req.GetActive(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to create employee: %e", err))
	}

	companyDb, _ := server.store.GetCompanyById(ctx, req.GetCompany())
	companyParent := req.GetCompany()
	if companyDb.Parent.Valid {
		companyParent = companyDb.Parent.Int32
	}
	_, err = server.store.CreateAccountCompany(ctx, db.CreateAccountCompanyParams{
		Account:       employee.ID,
		Company:       sql.NullInt32{Int32: req.GetCompany(), Valid: true},
		CompanyParent: sql.NullInt32{Int32: companyParent, Valid: true},
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("failed to create account company: %e", err))
	}

	return &pb.CreateEmployeeResponse{
		Code:    200,
		Message: "success",
		Details: employee.ID,
	}, nil
}

func (server *ServerGRPC) UpdateEmployee(ctx context.Context, req *pb.EmployeeUpdateRequest) (*pb.EmployeeUpdateResponse, error) {
	tokenPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	employee, err := server.store.GetAccount(ctx, req.GetId())
	if err != nil {
		errApp := common.ErrDB(err)
		return &pb.EmployeeUpdateResponse{
			Code:         int32(errApp.StatusCode),
			Message:      errApp.Message,
			MessageTrans: errApp.MessageTrans,
			Log:          errApp.Log,
		}, nil
	}

	hashPass, err := utils.HashedPassword(req.GetNewPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password")
	}

	var newAddressId int32
	if req.Address != nil {
		address, err := server.store.GetAddress(ctx, employee.Address.Int32)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				address, err := server.store.CreateAddress(ctx, db.CreateAddressParams{
					Lat:         float64(req.GetAddress().GetLat()),
					Lng:         float64(req.GetAddress().GetLng()),
					Province:    sql.NullString{String: req.GetAddress().GetProvince().GetCode(), Valid: true},
					District:    sql.NullString{String: req.GetAddress().GetDistrict().GetCode(), Valid: true},
					Ward:        sql.NullString{String: req.GetAddress().GetWard().GetCode(), Valid: true},
					Title:       req.GetAddress().GetTitle(),
					UserCreated: tokenPayload.UserID,
				})
				if err != nil {
					errApp := common.ErrDB(err)
					return &pb.EmployeeUpdateResponse{
						Code:         int32(errApp.StatusCode),
						Message:      errApp.Message,
						MessageTrans: errApp.MessageTrans,
						Log:          errApp.Log,
					}, nil
				}
				newAddressId = address.ID
			} else {
				errApp := common.ErrDB(err)
				return &pb.EmployeeUpdateResponse{
					Code:         int32(errApp.StatusCode),
					Message:      errApp.Message,
					MessageTrans: errApp.MessageTrans,
					Log:          errApp.Log,
				}, nil
			}
		} else {
			_, err = server.store.UpdateAddress(ctx, db.UpdateAddressParams{
				Lat:      sql.NullFloat64{Float64: float64(req.GetAddress().GetLat()), Valid: true},
				Lng:      sql.NullFloat64{Float64: float64(req.GetAddress().GetLng()), Valid: true},
				Province: sql.NullString{String: req.GetAddress().GetProvince().GetCode(), Valid: true},
				District: sql.NullString{String: req.GetAddress().GetDistrict().GetCode(), Valid: true},
				Ward:     sql.NullString{String: req.GetAddress().GetWard().GetCode(), Valid: true},
				Title:    sql.NullString{String: req.GetAddress().GetTitle(), Valid: true},
				ID:       address.ID,
			})
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to update address")
			}
		}
	}

	_, err = server.store.UpdateAccount(ctx, db.UpdateAccountParams{
		IsVerify: sql.NullBool{Bool: req.GetActive(), Valid: req.Active != nil},
		Password: sql.NullString{String: hashPass, Valid: req.NewPassword != nil},
		FullName: sql.NullString{String: req.GetFullName(), Valid: req.FullName != nil},
		Email:    sql.NullString{String: req.GetEmail(), Valid: req.Email != nil},
		Type:     sql.NullInt32{},
		Role:     sql.NullInt32{Int32: req.GetRole(), Valid: req.Role != nil},
		Gender:   db.NullGender{Gender: db.Gender(req.GetGender()), Valid: req.Gender != nil},
		Licence:  sql.NullString{String: req.GetLicence(), Valid: req.Licence != nil},
		Dob:      sql.NullTime{Time: time.Unix(req.GetDob().GetSeconds(), 0), Valid: req.Dob.IsValid()},
		Address:  sql.NullInt32{Int32: newAddressId, Valid: newAddressId != 0},
		ID:       sql.NullInt32{Int32: req.Id, Valid: true},
	})
	if err != nil {
		errApp := common.ErrDBWithMsg(err, "Cập nhật thông tin nhân viên thất bại")
		return &pb.EmployeeUpdateResponse{
			Code:         int32(errApp.StatusCode),
			Message:      errApp.Message,
			MessageTrans: errApp.MessageTrans,
			Log:          errApp.Log,
		}, nil
	}

	return &pb.EmployeeUpdateResponse{
		Code:    200,
		Message: "success",
	}, nil
}

func (server *ServerGRPC) DetailEmployee(ctx context.Context, req *pb.EmployeeDetailRequest) (*pb.EmployeeDetailResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	account, err := server.store.GetAccount(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "account not exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to get account")
	}

	accountPb := mapper.AccountRowMapper(account)

	var address *pb.Address
	if account.Address.Valid {
		addressDb, _ := server.store.GetAddress(ctx, account.Address.Int32)
		address = mapper.AddressMapper(ctx, server.store, addressDb)
	}

	var role *pb.Role
	if account.Role.Valid {
		roleDb, err := server.store.RoleDetail(ctx, account.Role.Int32)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get role")
		}
		role = &pb.Role{
			Id:              roleDb.ID,
			Code:            roleDb.Code,
			Title:           roleDb.Title,
			Company:         roleDb.Company.Int32,
			UserCreatedName: roleDb.CreatedName,
			UserUpdatedName: roleDb.UpdatedName.String,
			CreatedAt:       timestamppb.New(roleDb.CreatedAt),
			UpdatedAt:       timestamppb.New(roleDb.UpdatedAt.Time),
		}
	}

	accountPb.Address = address
	accountPb.RoleData = role

	return &pb.EmployeeDetailResponse{
		Code:    200,
		Message: "success",
		Details: accountPb,
	}, nil
}

func (server *ServerGRPC) AssignRoleEmployee(ctx context.Context, req *pb.AssignRoleEmployeeRequest) (*pb.AssignRoleEmployeeResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	for _, item := range req.GetAccounts() {
		_, err = server.store.UpdateAccount(ctx, db.UpdateAccountParams{
			Role: sql.NullInt32{Int32: req.GetRole(), Valid: true},
			ID:   sql.NullInt32{Int32: item, Valid: true},
		})
		if err != nil {
			errApp := common.ErrDB(err)
			return &pb.AssignRoleEmployeeResponse{
				Code:    int32(errApp.StatusCode),
				Message: errApp.Message,
				Log:     errApp.Log,
			}, nil
		}
	}

	return &pb.AssignRoleEmployeeResponse{
		Code:    200,
		Message: "success",
	}, nil
}

func (server *ServerGRPC) DeleteEmployee(ctx context.Context, req *pb.AccountDetailRequest) (*pb.AccountDetailResponse, error) {
	_, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, config.UnauthenticatedError(err)
	}

	_, err = server.store.DeleteEmployee(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete account: %e", err)
	}

	return &pb.AccountDetailResponse{
		Code:    200,
		Message: "success",
	}, nil
}
