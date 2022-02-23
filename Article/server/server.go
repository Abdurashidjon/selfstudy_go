package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gitlab.com/Udevs/Article/models"
	_ "gitlab.com/Udevs/Article/server/docs"
	"gitlab.com/Udevs/Article/storage"
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
	router.GET("/articles/:id", getArticleById)
	router.DELETE("/articles/:id", deleteById)
	router.PUT("/articles", updateById)
	//router.GET("articles/search", getSearch)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}

type DefaultError struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code int `json:"code"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

// Bu funksiya esa create qilish uchun
// @Router /articles [POST]
// @Summary Create article
// @Description API for creating article
// @Tags article
// @Accept json
// @Produce json
// @Param article body models.ArticleReq true "article"
// @Success 201 {array} SuccessResponse
// @Failure 400,404  {object} ErrorResponse
// @Failure default {object} DefaultError
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
// @Router /articles [GET]
// @Summary Get All
// @Description API for get all article
// @Tags article
// @Accept json
// @Produce json
// @Success 201 {array} models.Article
// @Failure 400,404 {object} DefaultError
// @Failure 500,503 {object} DefaultError
func getArticle(c *gin.Context) {
	resp := articleStorage.GetList()
	if resp == nil {
		log.Println("Error not found dokument")
		c.JSON(404, "Not found")
	}
	c.JSON(200, resp)
}

// Id bo'yicha ma'lumot oluvchi funksiya
// @Router /articles/{id} [GET]
// @Summary Get By Id
// @Description API for get by id article
// @Tags article
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Article
// @Failure 400,404 {object} ErrorResponse
// @Failure 500,503 {object} DefaultError
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
// @Router /articles/{id} [DELETE]
// @Summary Delete By Id
// @Description API for delete by id article
// @Tags article
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {string}  SuccessResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500,503 {object} DefaultError
func deleteById(c *gin.Context) {
	id := c.Param("id")
	id1, _ := strconv.Atoi(id)
	res := articleStorage.Delete(id1)
	if res == storage.ErrorNotFound {
		c.JSON(http.StatusNotFound, "No document")
		return
	}
	c.JSON(http.StatusOK, res)
}

// id bo'yicha berilgan ma'lumotlarni yangilash
// @Router /articles [PUT]
// @Summary Update Article
// @Description API for updating articles
// @Tags article
// @Accept json
// @Produce json
// @Param article body models.ArticleReq true "article"
// @Success 200 {string} SuccessResponse
// @Failure 400,404 {object} DefaultError
// @Failure 500,503 {object} DefaultError
func updateById(c *gin.Context) {
	var update models.Article

	s := time.Now()
	update.CreatedAt = &s
	res := c.BindJSON(&update)
	if res != nil {
		c.JSON(422, res.Error())
		return
	}
	res = articleStorage.Update(update)
	if res != nil {
		c.JSON(200, res.Error())
		return
	}
	c.JSON(http.StatusOK, "message: Update completed")
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
