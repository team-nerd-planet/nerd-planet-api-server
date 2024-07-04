package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kataras/iris/v12"
	"github.com/team-nerd-planet/api-server/infra/config"
	"github.com/team-nerd-planet/api-server/infra/router/handler"
	"github.com/team-nerd-planet/api-server/infra/router/middleware"
	"github.com/team-nerd-planet/api-server/internal/controller/rest"

	"github.com/iris-contrib/swagger"              // swagger middleware for Iris
	"github.com/iris-contrib/swagger/swaggerFiles" // swagger embed files

	"github.com/team-nerd-planet/api-server/third_party/docs"
)

type Router struct {
	app  *iris.Application
	conf config.Config
}

func NewRouter(
	conf config.Config,
	itemCtrl rest.ItemController,
	tagCtrl rest.TagController,
	subscriptionCtrl rest.SubscriptionController,
	feedCtrl rest.FeedController,
) Router {
	if conf.Rest.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	app := iris.Default()
	app.UseRouter(middleware.CorsHandler())

	docs.SwaggerInfo.Host = conf.Swagger.Host
	docs.SwaggerInfo.BasePath = conf.Swagger.BasePath

	swaggerUI := swagger.Handler(swaggerFiles.Handler,
		swagger.URL("/v1/docs/swagger.json"),
		swagger.DeepLinking(true),
		swagger.Prefix("/v1/docs"),
	)
	v1 := app.Party("/v1")
	{
		v1.Get("/docs", swaggerUI)
		v1.Get("/docs/{any:path}", swaggerUI)

		item := v1.Party("/item")
		{
			item.Get("/", func(ctx iris.Context) { handler.ListItems(ctx, itemCtrl) })
			item.Get("/next", func(ctx iris.Context) { handler.FindNextItems(ctx, itemCtrl) })
			item.Post("/view_increase", func(ctx iris.Context) { handler.IncreaseViewCount(ctx, itemCtrl) })
			item.Post("/like_increase", func(ctx iris.Context) { handler.IncreaseLikeCount(ctx, itemCtrl) })
		}

		tag := v1.Party("/tag")
		{
			tag.Get("/job", func(ctx iris.Context) { handler.ListJobTags(ctx, tagCtrl) })
			tag.Get("/skill", func(ctx iris.Context) { handler.ListSkillTags(ctx, tagCtrl) })
		}

		subscription := v1.Party("/subscription")
		{
			subscription.Post("/apply", func(ctx iris.Context) { handler.Apply(ctx, subscriptionCtrl) })
			subscription.Post("/approve", func(ctx iris.Context) { handler.Approve(ctx, subscriptionCtrl) })
		}

		feed := v1.Party("/feed")
		{
			feed.Get("/search", func(ctx iris.Context) { handler.SearchFeedName(ctx, feedCtrl) })
		}
	}

	return Router{
		app:  app,
		conf: conf,
	}
}

func (r Router) Run() error {
	return r.app.Listen(fmt.Sprintf(":%d", r.conf.Rest.Port))
}
