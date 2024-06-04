package router

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/team-nerd-planet/api-server/infra/config"
	"github.com/team-nerd-planet/api-server/infra/router/handler"
	"github.com/team-nerd-planet/api-server/infra/router/middleware"
	"github.com/team-nerd-planet/api-server/internal/controller/rest"
	docs "github.com/team-nerd-planet/api-server/third_party/docs"
)

type Router struct {
	router *gin.Engine
	store  *persistence.InMemoryStore
	conf   *config.Config
}

func NewRouter(
	conf *config.Config,
	itemCtrl rest.ItemController,
	tagCtrl rest.TagController,
	subscriptionCtrl rest.SubscriptionController,
	feedCtrl rest.FeedController,
) Router {
	if conf.Rest.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(middleware.CorsHandler())
	s := persistence.NewInMemoryStore(time.Hour)

	docs.SwaggerInfo.Host = conf.Swagger.Host
	docs.SwaggerInfo.BasePath = conf.Swagger.BasePath
	v1 := r.Group("/v1")
	{
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

		item := v1.Group("/item")
		{
			item.GET("/", cache.CachePage(s, time.Hour, func(ctx *gin.Context) { handler.ListItems(ctx, itemCtrl) }))
		}

		tag := v1.Group("/tag")
		{
			tag.GET("/job", cache.CachePage(s, time.Hour, func(ctx *gin.Context) { handler.ListJobTags(ctx, tagCtrl) }))
			tag.GET("/skill", cache.CachePage(s, time.Hour, func(ctx *gin.Context) { handler.ListSkillTags(ctx, tagCtrl) }))
		}

		subscription := v1.Group("/subscription")
		{
			subscription.POST("/apply", func(ctx *gin.Context) { handler.Apply(ctx, subscriptionCtrl) })
			subscription.GET("/approve", func(ctx *gin.Context) { handler.ApproveGet(ctx, subscriptionCtrl) })
			subscription.POST("/approve", func(ctx *gin.Context) { handler.Approve(ctx, subscriptionCtrl) })
		}

		feed := v1.Group("/feed")
		{
			feed.GET("/search", func(ctx *gin.Context) { handler.SearchFeedName(ctx, feedCtrl) })
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
