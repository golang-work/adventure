package list

import (
    "encoding/json"
    "github.com/golang-work/adventure/models"
)

type account struct {
    source []*models.Account
    result []*struct{
        Username string    `json:"username"`
    }
}

func NewAccount(origin []*models.Account) *account {
    return &account{
        source: origin,
    }
}

func (l *account) Convert() interface{} {
    jsonByte, _ := json.Marshal(l.source)
    _ = json.Unmarshal(jsonByte, &l.result)

    return  l.result
}