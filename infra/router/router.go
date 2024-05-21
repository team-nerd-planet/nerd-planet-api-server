package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/team-nerd-planet/api-server/docs"
	"github.com/team-nerd-planet/api-server/infra/config"
	"github.com/team-nerd-planet/api-server/infra/router/handler"
	"github.com/team-nerd-planet/api-server/infra/router/middleware"
	"github.com/team-nerd-planet/api-server/internal/controller/rest"
)

type Router struct {
	router *gin.Engine
	conf   *config.Config
}

func NewRouter(conf *config.Config, ctrl rest.ItemController) Router {
	if conf.Rest.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(middleware.CorsHandler())

	docs.SwaggerInfo.BasePath = "/"
	v1 := r.Group("/v1")
	{
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		item := v1.Group("/item")
		{
			item.GET("/", func(c *gin.Context) { handler.ListItems(c, ctrl) })
		}
	}

	return Router{
		router: r,
		conf:   conf,
	}
}

func (r Router) Run() error {
	return r.router.Run(fmt.Sprintf(":%d", r.conf.Rest.Port))
}
