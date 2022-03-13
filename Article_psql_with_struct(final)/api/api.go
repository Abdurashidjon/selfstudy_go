package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"gitlab.com/Udevs/Article/api/docs"
	_ "gitlab.com/Udevs/Article/api/docs"
	"gitlab.com/Udevs/Article/api/handlers"
	"gitlab.com/Udevs/Article/config"
)

// @description this is a sample article demo
// @termsOfService https://udevs.io
func SetUpAPI(r *gin.Engine, h handlers.Handler, cfg config.Config) {
	// programmaticaly set swagger info 
	docs.SwaggerInfo_swagger.Title = cfg.App
	docs.SwaggerInfo_swagger.Version = cfg.Version
	docs.SwaggerInfo_swagger.Host = cfg.ServiceHost + cfg.HTTPPort
	docs.SwaggerInfo_swagger.Schemes = []string{"http","https"}

	//router := gin.Default()
	r.POST("/articles", h.CreateArticle)
	r.GET("/articles", h.GetListArticle)
	r.GET("/articles/:id", h.GetArticleById)
	r.PUT("/articles", h.UpdateById)
	r.DELETE("/articles/:id", h.DeleteById)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
