package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"gitlab.com/Udevs/Article/models"
	_ "gitlab.com/Udevs/Article/server/docs"
	"gitlab.com/Udevs/Article/storage"
	"gitlab.com/Udevs/Article/storage/postgres"
)

//var articleStorage storage.ArticleStorage
var db *sqlx.DB
var articleRepo storage.ArticleRepoI

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

	articleRepo = postgres.NewArticleRepo(db)

	// gin.SetMode(gin.DebugMode)
	// r := gin.New()
	// r.Use(gin.Logger(), gin.Recovery())

	router := gin.Default()
	router.POST("/articles", createArticle)
	router.GET("/articles", getListArticle)
	router.GET("/articles/:id", GetArticleById)
	router.PUT("/articles", updateById)
	router.DELETE("/articles/:id", deleteById)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}

// Bu funksiya esa create qilish uchun
// @Router /articles [POST]
// @Summary Create article
// @Description API for creating article
// @Tags article
// @Accept json
// @Produce json
// @Param article body models.ArticleCreate true "article"
// @Success 201 {array} models.SuccessResponse
// @Failure 400,404  {object} models.ErrorResponse
// @Failure default {object} models.DefaultError
func createArticle(c *gin.Context) {
	var entity models.Article
	err := c.ShouldBind(&entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: err.Error(),
		})
		return
	}
	resp := articleRepo.CreateArticle(entity)
	c.JSON(200, resp)
}

// Funksiya ma'lumotlarni list qilib oladi array ichiga
// @Router /articles [GET]
// @Summary Get All
// @Description API for get all article
// @Tags article
// @Param offset path int false "offset"
// @Param limit path int false "limit"
// @Param search path string false "search"
// @Accept json
// @Produce json
// @Success 201 {array} models.Article
// @Failure 400,404 {object} models.DefaultError
// @Failure 500,503 {object} models.DefaultError
func getListArticle(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", "0")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	resp, err := articleRepo.GetListArticle(models.Query{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
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
// @Failure 400,404 {object} models.ErrorResponse
// @Failure 500,503 {object} models.DefaultError
func GetArticleById(c *gin.Context) {
	id := c.Param("id")
	id1, _ := strconv.Atoi(id)
	resp, err := articleRepo.GetByIdArticle(id1)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}
	c.JSON(200, resp)
}

// Id bo'yicha ma'lumotni o'chiruvchi funksiya
// @Router /articles/{id} [DELETE]
// @Summary Delete By Id
// @Description API for delete by id article
// @Tags article
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {string}  models.SuccessResponse
// @Failure 400,404 {object} models.ErrorResponse
// @Failure 500,503 {object} models.DefaultError
func deleteById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
	}

	resp := articleRepo.DeleteByIdArticle(id)
	if err != nil {
		c.JSON(400, models.DefaultError{
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, resp)
}

// id bo'yicha berilgan ma'lumotlarni yangilash
// @Router /articles [PUT]
// @Summary Update Article
// @Description API for updating articles
// @Tags article
// @Accept json
// @Produce json
// @Param article body models.ArticleReq true "article"
// @Success 200 {string} models.SuccessResponse
// @Failure 400,404 {object} models.DefaultError
// @Failure 500,503 {object} models.DefaultError
func updateById(c *gin.Context) {
	var entity models.Article
	err := c.ShouldBind(&entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: err.Error(),
		})
		return
	}
	resp, err := articleRepo.UpdateByIdArticle(entity)
	if err != nil {
		c.JSON(400,"Bad error")
	}
	c.JSON(200,resp)
}
