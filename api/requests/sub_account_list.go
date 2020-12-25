package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-work/adventure/support"
)

type SubAccountList struct {
	GroupId uint `json:"groupId" validate:"required"`
}

func (r *SubAccountList) BindValid(ctx *gin.Context) error {
	_ = ctx.ShouldBindJSON(r)

	if err := validator.New().Struct(r); err != nil {
		return support.Throw("invalidParam", "message", err.Error())
	}

	return nil
}
