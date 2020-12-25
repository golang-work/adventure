package sms

import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)

type Sender interface {
    SendVcode(params ...string) error
}

type Vendor func() Sender

var adapters = make(map[string]Vendor)

func NewSenderVendor(adapterName string) (adapter Sender, err error) {
    handler, ok := adapters[adapterName]
    if !ok {
        err = fmt.Errorf("unknown sms adapter name %q (forgot to import?)", adapterName)
        return
    }
    adapter = handler()
    return
}

func Register(name string, adapter Vendor) {
    if adapter == nil {
        panic("register sms adapter is nil")
    }
    if _, ok := adapters[name]; !ok {
        adapters[name] = adapter
    }
}

func GenerateVcode(length int) string {
    dic := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    rand.Seed(time.Now().UnixNano())

    var res strings.Builder
    for i := 0; i < length; i++ {
        _, _ = fmt.Fprintf(&res, "%d", dic[rand.Intn(len(dic))])
    }
    return res.String()
}
