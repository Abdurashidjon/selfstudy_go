package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gitlab.com/Udevs/Article/models"
	_ "gitlab.com/Udevs/Article/server/docs"
	"gitlab.com/Udevs/Article/storage"
)

var articleStorage storage.ArticleStorage
var db *sqlx.DB

// @title Swager Example API
// @version 1.1
// @description this is a sample article demo
// @termsOfService https://udevs.io
// @host localhost:8080
func main() {

	psqlConnString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		5432,
		"postgres",
		"test1234",
		"bootcamp",
	)

	var err error
	db, err = sqlx.Connect("postgres", psqlConnString)
	if err != nil {
		log.Panic(err)
	}

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
	Code    int    `json:"code"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
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
	rows, err := db.Query(
		`SELECT ar.id, ar.title, ar.body, au.firstname, au.lastname, ar.created_at FROM article AS ar JOIN author AS au ON ar.author_id=au.id`,
	)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(rows)
	defer rows.Close()
	var arr []models.Article
	for rows.Next() {
		var a models.Article
		err = rows.Scan(
			&a.ID, &a.Body,
			&a.Title,
			&a.Author.Firstname,
			&a.Author.Lastname,
			&a.CreatedAt)
		arr = append(arr, a)
		if err != nil {
			log.Panic(err)
		}
	}
	fmt.Println(arr)
	c.JSON(200, arr)
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
