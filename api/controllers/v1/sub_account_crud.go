package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-work/adventure/api/domain"
	"github.com/golang-work/adventure/api/protocol"
	"github.com/golang-work/adventure/api/requests"
	"github.com/golang-work/adventure/models"
	"github.com/golang-work/adventure/support"
)

type subAccountCrud struct{}

func SubAccountCrud() *subAccountCrud {
	return &subAccountCrud{}
}

func (c *subAccountCrud) List(ctx *gin.Context) {
	request := &requests.SubAccountList{}
	masterId := domain.Auth(ctx).User().ID

	if err := request.BindValid(ctx); err != nil {
		protocol.Response(ctx).Abort(err).Json()
		return
	}

	list := domain.SubAccount(ctx).List(masterId, request.GroupId)
	protocol.Response(ctx).Result(list).Json()
	return
}

func (c *subAccountCrud) Store(ctx *gin.Context) {
	request := &requests.SubAccountStore{}
	masterId := domain.Auth(ctx).User().ID

	if err := domain.SubAccount(ctx).CheckLimit(masterId); err != nil {
		protocol.Response(ctx).Abort(err).Json()
		return
	}

	if err := request.BindValid(ctx); err != nil {
		protocol.Response(ctx).Abort(err).Json()
		return
	}

	subAccount, err := domain.SubAccount(ctx).Create(masterId,
		support.String(25), support.String(25))
	if err != nil {
		protocol.Response(ctx).Abort(err).Json()
		return
	}

	subAccount.GroupId = request.GroupId
	subAccount.GroupName = request.GroupName
	subAccount.Avatar = request.Avatar
	subAccount.Online = false
	subAccount.OnlineName = request.OnlineName

	support.DB.Create(subAccount)
	protocol.Response(ctx).Result(subAccount).Json()
	return
}

func (c *subAccountCrud) Destroy(ctx *gin.Context) {
	request := &requests.SubAccountDestroy{}
	masterId := domain.Auth(ctx).User().ID

	if err := request.BindValid(ctx); err != nil {
		protocol.Response(ctx).Abort(err).Json()
		return
	}

	domain.SubAccount(ctx).Destroy(masterId, request.Username)
	protocol.Response(ctx).Success().Json()
	return
}

func (c *subAccountCrud) Recover(ctx *gin.Context) {
	request := &requests.SubAccountRecover{}
	masterId := domain.Auth(ctx).User().ID

	if err := request.BindValid(ctx); err != nil {
		protocol.Response(ctx).Abort(err).Json()
		return
	}

	support.DB.Model(&models.SubAccount{}).
		Where("master_id = ? and username = ? and destroy_at is not null and destroy_at >= ?",
		masterId, request.Username, support.Now().Format("2006-01-02 15:04:05")).
		Update("destroy_at", nil)

	protocol.Response(ctx).Success().Json()
	return
}
