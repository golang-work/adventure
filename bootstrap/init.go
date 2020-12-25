package bootstrap

import (
	"github.com/golang-work/adventure/foundation"
	"github.com/golang-work/adventure/support"
)

func init() {
	support.Config = foundation.Config()
	support.Log = foundation.Zap()
	support.DB = foundation.Gorm()
	support.Redis = foundation.Redis()

	foundation.MysqlTables(support.DB)
}
