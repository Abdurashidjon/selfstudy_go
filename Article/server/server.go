package main

import (
	"fmt"
	"log"

	_ "gitlab.com/Udevs/Article/server/docs"
	"github.com/gin-gonic/gin"
	"gitlab.com/Udevs/Article/models"
	"gitlab.com/Udevs/Article/storage"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var articleStorage storage.ArticleStorage

// @title Swager Example API
// @version 1.1
// @description this is a sample article demo
// @termsOfService https://udevs.io
// @host localhost:8080
func main() {
	articleStorage = storage.NewArticleStorage()

	router := gin.Default()
	router.POST("/articles", createArticle)
	router.GET("/articles", getArticle)
	router.GET("/articles/id",getArticleById)

	router.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}

type DefaultError struct {
	Message string `json:"message"`
}

// @Router /articles [POST]
// @Summary Create articels
// @Description API for creating articles
// @Tags article
// @Accept json
// @Produce json
// @Param article body models.Article true "normative"
// @Success 201 {object} string "ok"
// @Failure 400,404 {object} DefaultError
// @Failure default {object} DefaultError
// @Failure 503 {object} DefaultError
func createArticle(c *gin.Context) {
	var article models.Article
	// t := time.Now()
	// article.CreatedAt = &t

	if err := c.BindJSON(&article); err != nil {
		log.Println(err)
		c.JSON(422, err.Error())
		return
	}

	err := articleStorage.Add(article)
	if err != nil {
		log.Println(err)
		c.JSON(500, err.Error())
		return
	}
	c.JSON(201, "Succes")
}

// GetAllHandler godoc
// @Tags article
// @Summary Show All Article
// @Description it gets all artciles from memory
// @ID get-all-handler
// @Accept json
// @Produce json
// @Success 200 {array} models.Article
// @Failure default {object} DefaultError
// @Router /articles [get]
func getArticle(c *gin.Context) {
	resp := articleStorage.GetList()
	if resp == nil {
		log.Println("Error not found dokument")
		c.JSON(404, "Not found")
	}
	fmt.Printf("%v\n",resp)
	c.JSON(200, resp)
}

func getArticleById(c *gin.Context)  {
	// var id int
	// var res models.Article
	// id = c.Query()
		
}