package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-work/adventure/models"
	"github.com/golang-work/adventure/support"
)

type account struct {
	ctx *gin.Context
}

func Account(ctx *gin.Context) *account {
	return &account{
		ctx: ctx,
	}
}

func (d *account) SignUp(username string, password string) (*models.Account, error) {
	account := models.Account{}
	rs := support.DB.Where("username = ?", username).First(&account)
	if rs.RowsAffected > 0 {
		return &account, support.Throw("accountExists")
	}
	account.Username = username
	hash, _ := support.Bcrypt(password, "")
	account.Password = hash

	support.DB.Create(&account)
	return &account, nil
}

func (d *account) SignIn(username string, password string) (*models.Account, error) {
	account := models.Account{}
	rs := support.DB.Where("username = ?", username).First(&account)
	if rs.RowsAffected <= 0 {
		return &account, support.Throw("accountNotFind")
	}
	_, err := support.Bcrypt(password, account.Password)
	if err != nil {
		return &account, support.Throw("passwordFail")
	}
	return &account, nil
}

func (d *account) FindById(id uint) (*models.Account, error) {
	var account models.Account
	if err := support.DB.Where("`id` = ?", id).First(&account).Error; err != nil {
		return &account, support.Throw("accountNotFind")
	}
	return &account, nil
}

func (d *account) ModifyPassword(account *models.Account, password string) {
	account.Password, _ = support.Bcrypt(password, "")
	support.DB.Save(&account)
}
