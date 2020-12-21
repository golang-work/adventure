package models

import (
	"github.com/golang-work/adventure/support"
)

type SubAccount struct {
	support.Model
	MasterId   uint   `json:"masterId" gorm:"type:uint,comment:主账号ID"`
	Username   string `json:"username" gorm:"comment:子账号名"`
	Password   string `json:"password" gorm:"comment:子账号密码"`
	GroupId    uint   `json:"groupId" gorm:"comment:区服ID"`
	GroupName  string `json:"groupName" gorm:"comment:区服ID"`
	Avatar     string `json:"avatar" gorm:"comment:角色封面资源路径"`
	Online     bool   `json:"online" gorm:"comment:是否在线"`
	OnlineName string `json:"onlineName" gorm:"comment:角色封面"`
	DestroyAt *support.Time `json:"destroyAt" gorm:"comment:删除时间,type:timestamp"`
}
