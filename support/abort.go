package support

import (
	"fmt"
)

type Abort struct {
	Level   string
	Code    int
	Message string
}

var LevelMap = map[string]int{
	"error":           500,
	"notice":          200,
	"404":             404,
	"warning":         404,
	"unauthenticated": 401,
	"unauthorized":    403,
	"invalidParam":   422,
}

func Throw(error string, extra ...string) Abort {

	conf := Config["abort"].Sub(error)
	if conf == nil {
		conf = Config["abort"].Sub("paramsError")
	}
	abort := Abort{
		Level: conf.GetString("level"),
		Code: conf.GetInt("code"),
		Message: conf.GetString("message"),
	}
	abort.Message = Padding(abort.Message, extra...)
	return abort
}

func (abort Abort) Error() string {
	return fmt.Sprintf("异常：%d - %s", abort.Code, abort.Message)
}
