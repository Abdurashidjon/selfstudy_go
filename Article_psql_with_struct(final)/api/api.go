package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "gitlab.com/Udevs/Article/api/docs"
	"gitlab.com/Udevs/Article/api/handlers"
)


// @title Swager Example API
// @version 1.1
// @description this is a sample article demo
// @termsOfService https://udevs.io
// @host localhost:8080
func SetUpAPI(r *gin.Engine, h handlers.Handler) {
	//router := gin.Default()
	r.POST("/articles", h.CreateArticle)
	r.GET("/articles", h.GetListArticle)
	r.GET("/articles/:id", h.GetArticleById)
	r.PUT("/articles", h.UpdateById)
	r.DELETE("/articles/:id", h.DeleteById)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
