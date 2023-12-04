package entities

import (
	"github.com/gin-gonic/gin"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"time"
)

type AccountEntity struct {
	Id                int64             `json:"id"`
	Username          string            `json:"username"`
	FullName          string            `json:"full_name"`
	Email             string            `json:"email"`
	Type              AccountTypeEntity `json:"type"`
	Media             *string           `json:"media"`
	PasswordChangedAt time.Time         `json:"password_changed_at"`
	CreatedAt         time.Time         `json:"created_at"`
}

func NewAccountEntity(store *db.Store, ctx *gin.Context, account db.Account) AccountEntity {
	accountType, _ := store.GetAccountType(ctx, account.Type)

	return AccountEntity{
		Id:       account.ID,
		Username: account.Username,
		FullName: account.FullName,
		Email:    account.Email,
		Type: AccountTypeEntity{
			Id:    accountType.ID,
			Code:  accountType.Code,
			Title: accountType.Title,
		},
		Media:             nil,
		PasswordChangedAt: account.PasswordChangedAt,
		CreatedAt:         account.CreatedAt,
	}
}
