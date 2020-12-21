package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-work/adventure/support"
)

type SubAccountRecover struct {
	Username string `json:"username" validate:"required"`
}

func (r *SubAccountRecover) BindValid(ctx *gin.Context) error {
	_ = ctx.ShouldBindJSON(r)

	if err := validator.New().Struct(r); err != nil {
		return support.Throw("invalid_param", "message", err.Error())
	}

	return nil
}
