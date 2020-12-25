package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-work/adventure/support"
)

type SubAccountStore struct {
	GroupId    uint   `json:"groupId" validate:"required"`
	GroupName  string `json:"groupName" validate:"required"`
	Avatar     string `json:"avatar" validate:"required"`
	OnlineName string `json:"onlineName" validate:"required"`
}

func (r *SubAccountStore) BindValid(ctx *gin.Context) error {
	_ = ctx.ShouldBindJSON(r)

	if err := validator.New().Struct(r); err != nil {
		return support.Throw("invalidParam", "message", err.Error())
	}

	return nil
}
