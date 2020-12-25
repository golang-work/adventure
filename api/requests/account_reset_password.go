package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-work/adventure/support"
)

type AccountResetPassword struct {
	Password           string `json:"password" validate:"required"`
	NewPassWord        string `json:"newPassWord" validate:"required"`
	NewPassWordConfirm string `json:"newPassWordConfirm" validate:"required"`
}

func (r *AccountResetPassword) BindValid(ctx *gin.Context) error {
	_ = ctx.ShouldBindJSON(r)

	if err := validator.New().Struct(r); err != nil {
		return support.Throw("invalidParam", "message", err.Error())
	}

	return nil
}
