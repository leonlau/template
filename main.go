package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"template/config"
	v "template/pkg/version" // 记得修改
	"template/router"
	"template/router/middleware"
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
	gin.SetMode(viper.GetString("runmode"))

	// create the Gin engine
	g := gin.New()

	router.Load(
		g,
		middleware.Logging(),
	)
	// Listen and serve on 0.0.0.0:8080
	g.Run(":8080")
}
