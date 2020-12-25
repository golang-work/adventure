package foundation

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"path"
	"strings"
)

func Config() map[string]*viper.Viper {
	conf := make(map[string]*viper.Viper)
	confDirName := "config"
	confFileExt := ".yaml"

	dir, err := ioutil.ReadDir(confDirName)
	if err != nil {
		panic(fmt.Errorf("ergodic config dir error: %s \n", err))
	}

	for _, fi := range dir {
		if !fi.IsDir() && strings.HasSuffix(fi.Name(), confFileExt) {
			v := viper.New()
			v.SetConfigFile(path.Join(confDirName, fi.Name()))
			err := v.ReadInConfig()
			if err != nil {
				panic(fmt.Errorf("loading config file error: %s \n", err))
			}
			conf[strings.TrimSuffix(fi.Name(), confFileExt)] = v
		}
	}
	return conf
}
