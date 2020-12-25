package foundation

import (
	"fmt"
	"github.com/golang-work/adventure/models"
	"github.com/golang-work/adventure/support"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
)

func Gorm() *gorm.DB {
	dft := support.Config["database"].GetString("default")
	conf := support.Config["database"].Sub("connections." + dft)
	if conf == nil {
		panic("database connections by " + dft + " not find")
	}
	switch conf.GetString("driver") {
	case "mysql":
		return gormMysql(conf)
	default:
		return gormMysql(conf)
	}
}

func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		models.Account{},
		models.SubAccount{},
		models.JwtToken{},
	)
	if err != nil {
		support.Log.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	support.Log.Info("register table success")
}

func gormMysql(conf *viper.Viper) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		conf.GetString("username"),
		conf.GetString("password"),
		conf.GetString("path"),
		conf.GetString("dbName"),
		conf.GetString("config"))
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         255,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(conf)); err != nil {
		support.Log.Error("mysql failed to start", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(
			conf.GetInt("maxIdleConns"))
		sqlDB.SetMaxOpenConns(
			conf.GetInt("maxOpenConns"))
		return db
	}
}

func gormConfig(conf *viper.Viper) *gorm.Config {
	if conf.GetBool("logMode") {
		return &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: conf.GetString("tablePrefix"),
			},
		}
	} else {
		return &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: conf.GetString("tablePrefix"),
			},
		}
	}
}
