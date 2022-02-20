package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/Udevs/Article/models"
	"gitlab.com/Udevs/Article/storage"
)

var articleStorage storage.ArticleStorage

func main() {
	articleStorage = storage.NewArticleStorage()

	router := gin.Default()
	router.POST("/articles", createArticle)
	router.GET("/articles", getArticle)
	router.GET("/articles/id",getArticleById)

	router.Run(":8080")
}

func createArticle(c *gin.Context) {
	var article models.Article
	t := time.Now()
	article.CreatedAt = &t

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