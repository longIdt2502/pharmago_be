package usecases

import (
	"github.com/gin-gonic/gin"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/domain/repositories"
	"github.com/longIdt2502/pharmago_be/token"
	"github.com/longIdt2502/pharmago_be/utils"
	"time"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string          `json:"access_token"`
	RefreshToken string          `json:"refresh_token"`
	Account      AccountResponse `json:"account"`
}

type LoginUseCase interface {
	GetUserByEmail(ctx *gin.Context, username string) (*db.Account, error)
	CheckPassword(password string, hashedPassword string) error
	CreateAccessToken(username string, duration time.Duration) (string, error)
}

type LoginUseCaseImpl struct {
	accountRepository repositories.AccountRepository
	tokenMaker        token.Maker
}

func (l *LoginUseCaseImpl) GetUserByEmail(ctx *gin.Context, username string) (*db.Account, error) {
	return l.accountRepository.GetAccountByUsername(ctx, username)
}

func (l *LoginUseCaseImpl) CheckPassword(password string, hashedPassword string) error {
	return utils.CheckPassword(password, hashedPassword)
}

func (l *LoginUseCaseImpl) CreateAccessToken(username string, duration time.Duration) (string, error) {
	return l.tokenMaker.CreateToken(username, duration)
}

func NewLoginUseCase(accountRepository repositories.AccountRepository, tokenMaker token.Maker) LoginUseCase {
	useCase := &LoginUseCaseImpl{
		accountRepository,
		tokenMaker,
	}
	return useCase
}
