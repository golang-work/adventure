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
	"invalid_param":   422,
}

var e = map[string]Abort{
	"server_error":            {"error", 20001, "服务器错误"},
	"params_error":            {"notice", 20002, "参数错误"},
	"account_exists":          {"notice", 20003, "账号已存在"},
	"account_not_find":        {"notice", 20004, "账号不存在"},
	"jwt_make_token_fail":     {"notice", 20005, "生成token失败"},
	"unauthenticated":         {"unauthenticated", 20006, "请登录后操作"},
	"unauthorized":            {"unauthorized", 20007, "请登录后操作"},
	"token_occupied":          {"unauthenticated", 20008, "您的帐户异地登陆或令牌失效"},
	"token_expired":           {"unauthenticated", 20009, "授权已过期"},
	"token_invalid":           {"unauthenticated", 20010, "token无效 {{reason}}"},
	"sub_account_count_limit": {"notice", 20011, "子账号数已达到上限"},
	"invalid_param":           {"invalid_param", 20012, "参数错误：{{message}}"},
	"different_new_password":  {"notice", 20013, "新密码与确认新密码不一致"},
	"different_password":      {"notice", 20014, "密码与确认密码不一致"},
	"sms_code_error":          {"notice", 20015, "验证码无效"},
	"set_sign_in_state_fail":  {"notice", 20016, "设置登录状态失败"},
	"add_token_to_black_fail": {"notice", 20017, "token加入黑名单失败"},
	"password_fail":           {"notice", 20018, "密码错误"},
}

func Throw(error string, extra ...string) Abort {
	abort, ok := e[error]
	if !ok {
		abort = e["params_error"]
	}
	abort.Message = Padding(abort.Message, extra...)
	return abort
}

func (abort Abort) Error() string {
	return fmt.Sprintf("异常：%d - %s", abort.Code, abort.Message)
}
