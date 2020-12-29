package list

import (
    "encoding/json"
    "github.com/golang-work/adventure/models"
)

type subAccount struct {
    source []*models.SubAccount
    result []*struct{
        Username   string `json:"username"`
        Password   string `json:"password"`
        GroupId    uint   `json:"groupId"`
        GroupName  string `json:"groupName"`
        Avatar     string `json:"avatar"`
        Online     bool   `json:"online"`
        OnlineName string `json:"onlineName"`
    }
}

func NewSubAccount(origin []*models.SubAccount) *subAccount {
    return &subAccount{
        source: origin,
    }
}

func (l *subAccount) Convert() interface{} {
    jsonByte, _ := json.Marshal(l.source)
    _ = json.Unmarshal(jsonByte, &l.result)

    return  l.result
}