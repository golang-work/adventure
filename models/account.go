package models

import (
	"github.com/golang-work/adventure/support"
)

type Account struct {
	support.Model
	Username string    `json:"username" gorm:"comment:用户名或手机号"`
	Password string    `json:"password" gorm:"comment:密码"`
}
