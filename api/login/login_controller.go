package login

import (
	"github.com/gin-gonic/gin"
	"github.com/longIdt2502/pharmago_be/domain/usecases"
	"github.com/longIdt2502/pharmago_be/utils"
	"net/http"
	"time"
)

type LoginController struct {
	LoginUseCase usecases.LoginUseCase
}

func (lc *LoginController) Login(ctx *gin.Context) {
	var req usecases.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	account, err := lc.LoginUseCase.GetUserByEmail(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
		return
	}

	err = lc.LoginUseCase.CheckPassword(req.Password, account.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.ErrorResponse(err))
		return
	}

	timeDuration, err := time.ParseDuration("24h")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	token, err := lc.LoginUseCase.CreateAccessToken(req.Username, timeDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}

	res := usecases.LoginResponse{
		AccessToken:  token,
		RefreshToken: "",
		Account:      usecases.NewAccountResponse(account),
	}

	ctx.JSON(http.StatusOK, res)
	return
}
