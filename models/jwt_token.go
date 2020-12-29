package models

import (
	"github.com/golang-work/adventure/support"
)

type JwtToken struct {
	support.Model
	AccountId uint   `gorm:"comment:用户ID"`
	Username  string `gorm:"comment:用户名"`
	Token     string `gorm:"type:text;comment:token"`
}
