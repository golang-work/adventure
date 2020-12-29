package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-work/adventure/provider"
	send "github.com/golang-work/adventure/provider/sms"
	"github.com/golang-work/adventure/support"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type vcode struct {
	ctx *gin.Context
}

func Vcode(ctx *gin.Context) *vcode {
	return &vcode{
		ctx: ctx,
	}
}

func (d *vcode) Send(conf *viper.Viper, handle string) error {
	vendor := conf.GetString("vendor")
	sender, err := send.NewSenderVendor(vendor)
	if err != nil {
		return err
	}
	// 不同供应商发送参数及顺序可能不同
	var params []string
	switch vendor {
	case provider.SENDER_BY_ALIYUN:
		params = append(params, handle,
			conf.GetString("sign." + vendor + ".signName"),
			conf.GetString("sign." + vendor + ".templateCode"),
			send.GenerateVcode(conf.GetInt("codeLength")))
	}

	err = sender.SendVcode(params...)
	if err != nil {
		support.Log.Error("send sms error", zap.Any("msg", err.Error()))
		return err
	}
	return nil
}
