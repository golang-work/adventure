package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-work/adventure/api/protocol"
	"github.com/golang-work/adventure/api/requests"
)

type outsideController struct{}

func OutsideController() *outsideController {
	return &outsideController{}
}

func (b *outsideController) ListLoginServer(ctx *gin.Context) {
	request := &requests.ListLoginServer{}

	if err := request.BindValid(ctx); err != nil {
		protocol.Response(ctx).Abort(err).Json()
		return
	}

	protocol.Response(ctx).Result(nil).Json()
	return
}

