package foundation

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/golang-work/adventure/support"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	var config string
	flag.StringVar(&config, "c", "config/local.yaml", "custom config file")
	flag.Parse()
	fmt.Printf("use config file：%v\n", config)

	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("loading config file error: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file updated:", e.Name)
		if err := v.Unmarshal(&support.Config); err != nil {
			fmt.Println(err)
		}
	})

	if err := v.Unmarshal(&support.Config); err != nil {
		fmt.Println(err)
	}
	return v
}
