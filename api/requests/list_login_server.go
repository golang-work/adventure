package requests

import (
	"github.com/gin-gonic/gin"
)

type ListLoginServer struct {
	ChannelId string `json:"channelId"`
	Version string `json:"version"`
	Platform string `json:"platform"`
}

func (r *ListLoginServer) BindValid(ctx *gin.Context) error {
	_ = ctx.ShouldBindJSON(r)

	return nil
}
