package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"template/handler"
	"template/handler/user"
	"template/pkg/check"
	"template/router/middleware"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	// Middlewares
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(mw...)

	g.NoRoute(handler.NotFound)

	// swagger api docs
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// pprof router
	pprof.Register(g)

	// api for authentication functionalities
	g.POST("/login", user.Login)

	v1 := g.Group("/v1/")
	v1.Use(middleware.AuthMiddleware())
	{
		v1.GET("/", user.Get)
	}
	// ping
	g.GET("/ping", check.HealthCheck)

	return g
}
