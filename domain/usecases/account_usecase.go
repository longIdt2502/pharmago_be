package usecases

import (
	"github.com/gin-gonic/gin"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/domain/repositories"
	"time"
)

type CreateAccountRequest struct {
	Username string `json:"username" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type AccountResponse struct {
	Username         string    `json:"username" binding:"required"`
	FullName         string    `json:"full_name" binding:"required"`
	Email            string    `json:"email" binding:"required,email"`
	PasswordChangeAt time.Time `json:"password_change_at"`
	CreateAt         time.Time `json:"create_at"`
}

func NewAccountResponse(account *db.Account) AccountResponse {
	return AccountResponse{
		Username:         account.Username,
		FullName:         account.FullName,
		Email:            account.Email,
		PasswordChangeAt: account.PasswordChangedAt,
		CreateAt:         account.CreatedAt,
	}
}

type AccountUseCase interface {
	GetAccount(ctx *gin.Context)
	GetAccountByUsername(ctx *gin.Context, username string) (*db.Account, error)
	CreateAccount(ctx *gin.Context) (*db.Account, error)
}

type accountUseCase struct {
	accountRepository repositories.AccountRepository
}

func (a accountUseCase) GetAccount(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a accountUseCase) GetAccountByUsername(ctx *gin.Context, username string) (*db.Account, error) {
	return a.accountRepository.GetAccountByUsername(ctx, username)
}

func (a accountUseCase) CreateAccount(ctx *gin.Context) (*db.Account, error) {
	panic("implement me")
}

func NewAccountUseCase(accountRepository repositories.AccountRepository) AccountUseCase {
	return &accountUseCase{
		accountRepository: accountRepository,
	}
}
