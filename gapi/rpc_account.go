package gapi

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/gapi/config"
	"github.com/longIdt2502/pharmago_be/gapi/mapper"
	"github.com/longIdt2502/pharmago_be/pb"
	"github.com/longIdt2502/pharmago_be/utils"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	accountPb := mapper.AccountMapper(account)

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
		Company: req.GetCompany(),
		Search: sql.NullString{
			String: req.GetSearch(),
			Valid:  true,
		},
		Type: sql.NullInt32{
			Int32: req.GetType(),
			Valid: req.Type != nil,
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
		return nil, status.Errorf(codes.Internal, "failed to get list account: %e", err)
	}

	var accountsPb []*pb.Account
	for _, item := range accountsDb {
		accountsPb = append(accountsPb, mapper.AccountMapper(db.Account{
			ID:        item.ID,
			Username:  item.Username,
			FullName:  item.FullName,
			Email:     item.Email,
			Type:      item.Type,
			Role:      item.Role,
			CreatedAt: item.CreatedAt,
		}))
	}

	return &pb.AccountListResponse{
		Code:    200,
		Message: "success",
		Details: accountsPb,
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

	_, err = server.store.CreateAccountCompany(ctx, db.CreateAccountCompanyParams{
		Account: employee.ID,
		Company: req.GetCompany(),
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
