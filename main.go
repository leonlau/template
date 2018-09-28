package main

import (
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"template/config"
	v "template/pkg/version" // 记得修改
)

var (
	cfg     = pflag.StringP("config", "c", "", "config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	pflag.Parse()

	// version
	if *version {
		v.Version()
	}

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	log.Info("successfully.")
	// Set gin mode.
	//gin.SetMode(viper.GetString("runmode"))

}
