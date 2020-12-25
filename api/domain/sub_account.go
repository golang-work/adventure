package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-work/adventure/models"
	"github.com/golang-work/adventure/support"
	"time"
)

type subAccount struct {
	ctx *gin.Context
}

func SubAccount(ctx *gin.Context) *subAccount {
	return &subAccount{
		ctx: ctx,
	}
}

func (d *subAccount) Create(masterId uint,
	username string, password string) (*models.SubAccount, error) {
	account := models.SubAccount{}

	result := support.DB.Where("master_id = ? and username = ?",
		masterId, username).First(&account)
	if result.RowsAffected > 0 {
		return &account, support.Throw("account_exists")
	}
	account.MasterId = masterId
	account.Username = username
	account.Password = password
	return &account, nil
}

func (d *subAccount) CheckLimit(masterId uint) error {
	var count int64
	support.DB.Model(&models.SubAccount{}).Where("master_id = ?", masterId).Count(&count)

	if count >= support.Config["account"].GetInt64("subCountLimit") {
		return support.Throw("sub_account_count_limit")
	}
	return nil
}

func (d *subAccount) Destroy(masterId uint, username string) error {
	sa := &models.SubAccount{}
	res := support.DB.Where("master_id = ? and username = ?",
		masterId, username).First(sa)
	if res.RowsAffected <= 0 {
		return support.Throw("account_not_find")
	}

	if sa.DestroyAt == nil {
		t := support.Time(support.Now().Add(time.Duration(support.Config["account"].GetInt("recoverValidityHour")) * time.Hour))
		sa.DestroyAt = &t
		support.DB.Save(sa)
	}
	return nil
}

func (d *subAccount) List(masterId uint, groupId uint) []*models.SubAccount {
	var list []*models.SubAccount
	support.DB.Where("master_id = ? and group_id = ?",
		masterId, groupId).Find(&list)
	return list
}
