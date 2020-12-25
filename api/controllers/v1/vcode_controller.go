package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-work/adventure/api/domain"
	"github.com/golang-work/adventure/api/protocol"
	"github.com/golang-work/adventure/api/requests"
	"github.com/golang-work/adventure/support"
	"go.uber.org/zap"
)

type vcodeController struct{}

func VcodeController() *vcodeController {
	return &vcodeController{}
}

func (c *vcodeController) Send(ctx *gin.Context) {
	request := &requests.VcodeSend{}

	if err := request.BindValid(ctx); err != nil {
		protocol.Response(ctx).Abort(err).Json()
		return
	}

	if request.AppId == "" {
		request.AppId = support.Config["vcode"].GetString("default")
	}

	conf := support.Config["vcode"].Sub(request.AppId + "." + request.Category)
	if conf == nil {
		protocol.Response(ctx).Abort("paramsError").Json()
		return
	}

	err := domain.Vcode(ctx).Send(conf, request.Handle)
	if err != nil {
		support.Log.Error("send sms error", zap.Any("msg", err.Error()))
		protocol.Response(ctx).Abort(err).Json()
		return
	}

	protocol.Response(ctx).Success().Json()
	return
}
