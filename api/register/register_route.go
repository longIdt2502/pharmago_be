package register

import (
	"github.com/gin-gonic/gin"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/domain/repositories"
	"github.com/longIdt2502/pharmago_be/domain/usecases"
	"github.com/longIdt2502/pharmago_be/token"
)

func NewRegisterRouter(group *gin.RouterGroup, store *db.Store, token token.Maker) {
	arp := repositories.NewAccountRepositoryImpl(store)
	ru := usecases.NewRegisterUseCase(arp, token)
	rc := &RegisterController{
		RegisterUseCase: ru,
		store:           store,
	}
	group.POST("/accounts/v1/register/", rc.Register)
}
