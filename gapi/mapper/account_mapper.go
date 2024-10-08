package mapper

import (
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func AccountMapper(account db.Account) *pb.Account {
	return &pb.Account{
		Id:                account.ID,
		Username:          account.Username,
		FullName:          account.FullName,
		Email:             account.Email,
		IsActive:          account.IsVerify,
		Role:              &account.Role.Int32,
		Gender:            (*string)(&account.Gender.Gender),
		Licence:           &account.Licence.String,
		Dob:               timestamppb.New(account.Dob.Time),
		Address:           &pb.Address{},
		RoleData:          &pb.Role{},
		PasswordChangedAt: timestamppb.New(account.PasswordChangedAt),
		CreatedAt:         timestamppb.New(account.CreatedAt),
	}
}

func AccountRowMapper(account db.GetAccountRow) *pb.Account {
	return &pb.Account{
		Id:                account.ID,
		Username:          account.Username,
		FullName:          account.FullName,
		Email:             account.Email,
		IsActive:          account.IsVerify,
		Role:              &account.Role.Int32,
		Gender:            (*string)(&account.Gender.Gender),
		Licence:           &account.Licence.String,
		Dob:               timestamppb.New(account.Dob.Time),
		Address:           &pb.Address{},
		RoleData:          &pb.Role{},
		PasswordChangedAt: timestamppb.New(account.PasswordChangedAt),
		CreatedAt:         timestamppb.New(account.CreatedAt),
		CompanyName:       account.Name.String,
	}
}

func ListAccountRowMapper(account db.ListAccountRow) *pb.Account {
	return &pb.Account{
		Id:          account.ID,
		Username:    account.Username,
		FullName:    account.FullName,
		Email:       account.Email,
		VerifyId:    0,
		IsActive:    account.IsVerify,
		AccountType: account.Title.String,
		Role:        &account.Role.Int32,
		Gender:      (*string)(&account.Gender.Gender),
		Licence:     &account.Licence.String,
		Dob:         timestamppb.New(account.Dob.Time),
		RoleData: &pb.Role{
			Id:    account.Role.Int32,
			Code:  account.Code_3.String,
			Title: account.Title_2.String,
		},
		PasswordChangedAt: timestamppb.New(account.PasswordChangedAt),
		CreatedAt:         timestamppb.New(account.CreatedAt),
		CompanyName:       account.Name.String,
	}
}
