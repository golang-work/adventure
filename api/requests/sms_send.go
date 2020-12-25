package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-work/adventure/support"
)

type VcodeSend struct {
	AppId string `json:"appId"`
	Category string `json:"category" validate:"required"`
	Handle string `json:"handle" validate:"required"`
}

func (r *VcodeSend) BindValid(ctx *gin.Context) error {
	_ = ctx.ShouldBindJSON(r)

	if err := validator.New().Struct(r); err != nil {
		return support.Throw("invalidParam", "message", err.Error())
	}

	return nil
}
