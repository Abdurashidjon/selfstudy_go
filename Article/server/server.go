package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	router.GET("/articles/:id", getArticleById)
	router.DELETE("/articles/:id", deleteById)
	router.PUT("articles", updateById)
	//router.GET("articles/search", getSearch)

	router.Run(":8080")
}

// Bu funksiya esa create qilish uchun
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

// Funksiya ma'lumotlarni list qilib oladi array ichiga
func getArticle(c *gin.Context) {
	resp := articleStorage.GetList()
	if resp == nil {
		log.Println("Error not found dokument")
		c.JSON(404, "Not found")
	}
	c.JSON(200, resp)
}

// Id bo'yicha ma'lumot oluvchi funksiya
func getArticleById(c *gin.Context) {
	id := c.Param("id")
	id1, _ := strconv.Atoi(id)
	res, err := articleStorage.GetByID(id1)
	if err != nil {
		c.JSON(404, "Not found by Id")
		return
	}
	c.JSON(http.StatusOK, res)
}

// Id bo'yicha ma'lumotni o'chiruvchi funksiya
func deleteById(c *gin.Context) {
	id := c.Param("id")
	id1, _ := strconv.Atoi(id)
	res := articleStorage.Delete(id1)
	if res == storage.ErrorNotFound {
		c.JSON(http.StatusNotFound,"No document")
		return
	}
	c.JSON(http.StatusOK, res)
}

// id bo'yicha berilgan ma'lumotlarni yangilash
func updateById(c *gin.Context) {
	var update models.Article

	res:= articleStorage.Update(update)
	if err := c.BindJSON(&res); err != nil {
		log.Println(err)
		c.JSON(422, err.Error())
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, res.Error())
		return
	}

	fmt.Println("aaaaaa ", res)
	c.JSON(http.StatusOK, res)
}

// func getSearch(c *gin.Context) {
// 	search := c.Param("search")
// 	res, _ := articleStorage.Search(search)
// 	//err := articleStorage.Delete(id1)
// 	if res != nil {
// 		log.Println(res)
// 		c.JSON(404, "Not found by Id")
// 		return
// 	}
// 	c.JSON(200, res)
// }
