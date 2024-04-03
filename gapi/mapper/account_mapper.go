package mapper

import (
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func AccountMapper(account db.Account) *pb.Account {
	return &pb.Account{
		Username:          account.Username,
		FullName:          account.FullName,
		Email:             account.Email,
		OaId:              &account.OaID.String,
		PasswordChangedAt: timestamppb.New(account.PasswordChangedAt),
		CreatedAt:         timestamppb.New(account.CreatedAt),
	}
}
