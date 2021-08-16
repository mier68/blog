package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/goprogramming/blog/internal/routers/api/v1"
)


func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	tag := v1.NewTag()
	article := v1.NewAriticle()
	apivl := r.Group("/api/v1")
	{
		apivl.POST("tags",tag.Create)
		apivl.DELETE("tags/:id",tag.Delete)
		apivl.PUT("/tags/:id",tag.Update)
		apivl.PATCH("tags/:id/state",tag.Update)
		apivl.GET("tags",tag.List)

		apivl.POST("/articles",article.Create)
		apivl.DELETE("/articles/:id",article.Delete)
		apivl.PUT("/articles/:id",article.Update)
		apivl.PATCH("/articles/:id/state",article.Update)
		apivl.GET("/articles/:id",article.Get)
		apivl.GET("/articles",article.List)
	}
	return r
}
