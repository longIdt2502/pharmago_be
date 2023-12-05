package repositories

import (
	"github.com/gin-gonic/gin"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
)

type AccountRepository interface {
	GetAccount(ctx *gin.Context)
	GetAccountByUsername(ctx *gin.Context, username string) (*db.Account, error)
	CreateAccount(ctx *gin.Context, arg db.CreateAccountParams) (db.Account, error)
}

type AccountRepositoryImpl struct {
	store *db.Store
}

func (a *AccountRepositoryImpl) GetAccount(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a *AccountRepositoryImpl) GetAccountByUsername(ctx *gin.Context, username string) (*db.Account, error) {
	account, err := a.store.GetAccountByUseName(ctx, username)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (a *AccountRepositoryImpl) CreateAccount(ctx *gin.Context, arg db.CreateAccountParams) (db.Account, error) {
	return a.store.CreateAccount(ctx, arg)
}

func NewAccountRepositoryImpl(store *db.Store) AccountRepository {
	return &AccountRepositoryImpl{
		store: store,
	}
}
