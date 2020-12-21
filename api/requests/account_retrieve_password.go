package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-work/adventure/support"
)

type AccountRetrievePassword struct {
	Password        string `json:"password" validate:"required"`
	PassWordConfirm string `json:"passWordConfirm" validate:"required"`
	Code            string `json:"code" validate:"required"`
}

func (r *AccountRetrievePassword) BindValid(ctx *gin.Context) error {
	_ = ctx.ShouldBindJSON(r)

	if err := validator.New().Struct(r); err != nil {
		return support.Throw("invalid_param", "message", err.Error())
	}

	return nil
}
