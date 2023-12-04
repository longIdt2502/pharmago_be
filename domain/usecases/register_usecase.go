package usecases

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/domain/entities"
	"github.com/longIdt2502/pharmago_be/domain/repositories"
	"github.com/longIdt2502/pharmago_be/token"
	"github.com/longIdt2502/pharmago_be/utils"
	"time"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type RegisterResponse struct {
	AccessToken string                 `json:"access_token"`
	Account     entities.AccountEntity `json:"account"`
}

type RegisterUseCase interface {
	CheckUniqueUsername(ctx *gin.Context, username string) bool
	HashedPassword(password string) (string, error)
	CreateAccount(ctx *gin.Context, request RegisterRequest, hashedPassword string) (db.Account, error)
	CreateAccessToken(username string, duration time.Duration) (string, error)
	CreateRefreshToken()
}

type registerUseCase struct {
	accountRepository repositories.AccountRepository
	tokenMaker        token.Maker
}

func (r *registerUseCase) CheckUniqueUsername(ctx *gin.Context, username string) bool {
	//TODO implement me
	_, err := r.accountRepository.GetAccountByUsername(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return true
		}
		return false
	}
	return false
}

func (r *registerUseCase) HashedPassword(password string) (string, error) {
	return utils.HashedPassword(password)
}

func (r *registerUseCase) CreateAccount(ctx *gin.Context, request RegisterRequest, hashedPassword string) (db.Account, error) {
	arg := db.CreateAccountParams{
		Username:       request.Username,
		HashedPassword: hashedPassword,
		FullName:       request.FullName,
		Email:          request.Email,
		Type:           1,
		Media:          sql.NullInt64{Int64: 0},
	}
	return r.accountRepository.CreateAccount(ctx, arg)
}

func (r *registerUseCase) CreateAccessToken(username string, duration time.Duration) (string, error) {
	return r.tokenMaker.CreateToken(username, duration)
}

func (r *registerUseCase) CreateRefreshToken() {
	//TODO implement me
	panic("implement me")
}

func NewRegisterUseCase(accountRepository repositories.AccountRepository, tokenMaker token.Maker) RegisterUseCase {
	return &registerUseCase{
		accountRepository,
		tokenMaker,
	}
}
