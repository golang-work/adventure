package provider

import (
    "fmt"
    "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
    "github.com/golang-work/adventure/provider/sms"
    "github.com/golang-work/adventure/support"
    "go.uber.org/zap"
)

const SENDER_BY_ALIYUN = "aliyun"

type aliyun struct {
    accessKeyId string
    accessKeySecret string
}

func init()  {
    sms.Register(SENDER_BY_ALIYUN, func() sms.Sender {
        return NewAliyun(support.Config["vendor"].GetString("aliyun.accessKeyId"),
            support.Config["vendor"].GetString("aliyun.accessKeySecret"))
    })
}

func NewAliyun(accessKeyId, accessKeySecret string) *aliyun {
    return &aliyun{
        accessKeyId: accessKeyId,
        accessKeySecret: accessKeySecret,
    }
}

func (a *aliyun) SendVcode (params ...string) error {
    client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou",
        a.accessKeyId, a.accessKeySecret)

    request := dysmsapi.CreateSendSmsRequest()
    request.Scheme = "https"

    request.PhoneNumbers = params[0]
    request.SignName = params[1]
    request.TemplateCode = params[2]
    request.TemplateParam = fmt.Sprintf(`{"code":"%s"}`, params[3])

    response, err := client.SendSms(request)
    if err != nil {
        support.Log.Error("send sms by aliyun error", zap.Any("msg", err.Error()))
        return err
    }
    support.Log.Info("send sms by aliyun return", zap.Any("msg", fmt.Sprintf("%#v", response)))

    return nil
}
