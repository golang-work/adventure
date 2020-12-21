package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-work/adventure/api/domain"
	"github.com/golang-work/adventure/api/protocol"
	"github.com/golang-work/adventure/api/requests"
	"github.com/golang-work/adventure/support"
)

type accountController struct{}

func AccountController() *accountController {
	return &accountController{}
}

func (c *accountController) SignUp(ctx *gin.Context) {
	request := &requests.AccountSignUp{}

	if err := request.BindValid(ctx); err != nil {
		protocol.Response(ctx).Abort(err).Json()
		return
	}

	account, err := domain.Account(ctx).SignUp(request.Username, request.Password)
	if err != nil {
		protocol.Response(ctx).Abort(err).Json()
		return
	}

	protocol.Response(ctx).Result(account).Json()
	return
}

func (c *accountController) SignIn(ctx *gin.Context) {
	request := &requests.AccountSignIn{}

	if err := request.BindValid(ctx); err != nil {
		protocol.Response(ctx).Abort(err).Json()
		return
	}

	account, err := domain.Account(ctx).SignIn(request.Username, request.Password)
	if err != nil {
		protocol.Response(ctx).Abort(err).Json()
		return
	}

	token, expiresAt, err := domain.Auth(ctx).MakeToken(account)
	if err != nil {
		protocol.Response(ctx).Abort(err).Json()
		return
	}

	protocol.Response(ctx).Result(gin.H{
		"token":     token,
		"expiresAt": expiresAt,
	}).Json()
	return
}

func (c *accountController) ResetPassword(ctx *gin.Context) {
	request := &requests.AccountResetPassword{}

	if err := request.BindValid(ctx); err != nil {
		protocol.Response(ctx).Abort(err).Json()
		return
	}

	if request.NewPassWord != request.NewPassWordConfirm {
		protocol.Response(ctx).Abort("different_new_password").Json()
		return
	}

	masterUserName := domain.Auth(ctx).User().Username
	account, err := domain.Account(ctx).SignIn(masterUserName, request.Password)
	if err != nil {
		protocol.Response(ctx).Abort(err).Json()
		return
	}

	account.Password, _ = support.Bcrypt(request.NewPassWord, "")
	support.DB.Save(account)

	protocol.Response(ctx).Success().Json()
	return
}

func (c *accountController) RetrievePassword(ctx *gin.Context) {
	request := &requests.AccountRetrievePassword{}
	accountDomain := domain.Account(ctx)

	if err := request.BindValid(ctx); err != nil {
		protocol.Response(ctx).Abort(err).Json()
		return
	}

	if request.Password != request.PassWordConfirm {
		protocol.Response(ctx).Abort("different_password").Json()
		return
	}

	// todo:: 短信验证保持通过
	if request.Code != "10866" {
		protocol.Response(ctx).Abort("sms_code_error").Json()
		return
	}

	account, err := accountDomain.FindById(domain.Auth(ctx).User().ID)
	if err == nil {
		accountDomain.ModifyPassword(account, request.Password)
	}

	protocol.Response(ctx).Success().Json()
	return
}
