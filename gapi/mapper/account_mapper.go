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
		AccountType:       "",
		CompanyName:       "",
		Role:              &account.Role.Int32,
		Gender:            (*string)(&account.Gender.Gender),
		Licence:           &account.Licence.String,
		Dob:               timestamppb.New(account.Dob.Time),
		PasswordChangedAt: timestamppb.New(account.PasswordChangedAt),
		CreatedAt:         timestamppb.New(account.CreatedAt),
	}
}

func ListAccountRowMapper(account db.ListAccountRow) *pb.Account {
	return &pb.Account{
		Id:                account.ID,
		Username:          account.Username,
		FullName:          account.FullName,
		Email:             account.Email,
		IsActive:          account.IsVerify,
		AccountType:       account.Title.String,
		CompanyName:       account.Name.String,
		Role:              &account.Role.Int32,
		Gender:            (*string)(&account.Gender.Gender),
		Licence:           &account.Licence.String,
		Dob:               timestamppb.New(account.Dob.Time),
		PasswordChangedAt: timestamppb.New(account.PasswordChangedAt),
		CreatedAt:         timestamppb.New(account.CreatedAt),
	}
}
