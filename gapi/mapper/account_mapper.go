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
		PasswordChangedAt: timestamppb.New(account.PasswordChangedAt),
		CreatedAt:         timestamppb.New(account.CreatedAt),
		IsActive:          account.IsVerify,
		Role:              &account.Role.Int32,
		Gender:            (*string)(&account.Gender.Gender),
		Licence:           &account.Licence.String,
		Dob:               timestamppb.New(account.Dob.Time),
	}
}
