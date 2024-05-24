package router

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
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
	store  *persistence.InMemoryStore
	conf   *config.Config
}

func NewRouter(conf *config.Config, itemCtrl rest.ItemController, tabCtrl rest.TagController) Router {
	if conf.Rest.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(middleware.CorsHandler())
	s := persistence.NewInMemoryStore(time.Hour)

	docs.SwaggerInfo.BasePath = "/"
	v1 := r.Group("/v1")
	{
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

		item := v1.Group("/item")
		{
			item.GET("/", cache.CachePage(s, time.Hour, func(c *gin.Context) { handler.ListItems(c, itemCtrl) }))
		}

		tag := v1.Group("/tag")
		{
			tag.GET("/job", cache.CachePage(s, time.Hour, func(c *gin.Context) { handler.ListJobTags(c, tabCtrl) }))
			tag.GET("/skill", cache.CachePage(s, time.Hour, func(c *gin.Context) { handler.ListSkillTags(c, tabCtrl) }))
		}
	}

	return Router{
		router: r,
		store:  s,
		conf:   conf,
	}
}

func (r Router) Run() error {
	return r.router.Run(fmt.Sprintf(":%d", r.conf.Rest.Port))
}
