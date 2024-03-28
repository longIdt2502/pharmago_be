package register

// import (
// 	"github.com/gin-gonic/gin"
// 	db "github.com/longIdt2502/pharmago_be/db/sqlc"
// 	"github.com/longIdt2502/pharmago_be/domain/entities"
// 	"github.com/longIdt2502/pharmago_be/domain/usecases"
// 	"github.com/longIdt2502/pharmago_be/utils"
// 	"net/http"
// 	"time"
// )

// type RegisterController struct {
// 	RegisterUseCase usecases.RegisterUseCase
// 	store           *db.Store
// }

// func (rc *RegisterController) Register(ctx *gin.Context) {
// 	var req usecases.RegisterRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
// 		return
// 	}

// 	if !rc.RegisterUseCase.CheckUniqueUsername(ctx, req.Username) {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"err": "username already exist"})
// 		return
// 	}

// 	hashedPassword, err := utils.HashedPassword(req.Password)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
// 		return
// 	}

// 	token, err := rc.RegisterUseCase.CreateAccessToken(req.Username, time.Duration(60*60*24))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
// 		return
// 	}

// 	account, err := rc.RegisterUseCase.CreateAccount(ctx, req, hashedPassword)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
// 		return
// 	}

// 	accountEntity := entities.NewAccountEntity(rc.store, ctx, account)

// 	res := usecases.RegisterResponse{
// 		AccessToken: token,
// 		Account:     accountEntity,
// 	}

// 	ctx.JSON(http.StatusCreated, res)
// 	return
// }
