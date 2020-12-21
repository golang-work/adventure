package bootstrap

import (
	"fmt"
	"github.com/golang-work/adventure/routes"
	"github.com/golang-work/adventure/support"
	"go.uber.org/zap"
	"time"
)

type server interface {
	ListenAndServe() error
}

func Run() {
	router := routes.Routers()
	//router.Static("/assets", "./storage/assets")
	//router.StaticFS(support.Config.Oss.Qiniu.ImgPath, http.Dir(support.Config.Oss.Qiniu.ImgPath))

	address := support.Config.System.Address
	s := initServer(address, router)
	// 保证文本顺序输出
	time.Sleep(10 * time.Microsecond)
	support.Log.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
API运行地址:http://127.0.0.1%s
`, address)
	support.Log.Error(s.ListenAndServe().Error())
}
