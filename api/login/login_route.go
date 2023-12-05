package login

import (
	"github.com/gin-gonic/gin"
	db "github.com/longIdt2502/pharmago_be/db/sqlc"
	"github.com/longIdt2502/pharmago_be/domain/repositories"
	"github.com/longIdt2502/pharmago_be/domain/usecases"
	"github.com/longIdt2502/pharmago_be/token"
)

func NewLoginRouter(group *gin.RouterGroup, store *db.Store, token token.Maker) {
	arp := repositories.NewAccountRepositoryImpl(store)
	lu := usecases.NewLoginUseCase(arp, token)
	lc := &LoginController{
		LoginUseCase: lu,
	}
	group.POST("/accounts/v1/login/", lc.Login)
}
