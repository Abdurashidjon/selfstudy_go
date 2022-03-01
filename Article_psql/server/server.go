package main

import (
	"fmt"
	"log"
	"net/http"

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
	var article models.Article
	err := c.BindJSON(&article)
	if err != nil {
		c.JSON(400, err.Error())
	}
	query := `INSERT INTO author(
		firstname,
		lastname,
		) 
	VALUES ($1, $2)`
	query2 := `INSERT INTO article(
		id,
		title,
		body
		) 
		VALUES ($1, %2)`

	_,err = db.Exec(query)
	if err != nil {
		panic(err)
	}
	_,err = db.Exec(query2)
	if err != nil {
		panic(err)
	}
	c.JSON(200, "Succes created")
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
		`SELECT 
		ar.id, 
		ar.title, 
		ar.body, 
		au.firstname, 
		au.lastname, 
		ar.created_at FROM article AS ar JOIN author AS au ON ar.author_id = au.id`,
	)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer rows.Close()
	var arr []models.Article
	for rows.Next() {
		var a models.Article
		err = rows.Scan(
			&a.ID,
			&a.Body,
			&a.Title,
			&a.Author.Firstname,
			&a.Author.Lastname,
			&a.CreatedAt)
		arr = append(arr, a)
		if err != nil {
			log.Panic(err)
		}
	}
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
	var arr models.Article
	id := c.Param("id")
	rows, err := db.Query(
		`SELECT
				ar.id,
				ar.title,
				ar.body,
				au.firstname,
				au.lastname,
				ar.created_at
				FROM article AS ar  JOIN author AS au ON (ar.author_id = au.id) WHERE ar.id=$1`, id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&arr.ID,
			&arr.Body,
			&arr.Title,
			&arr.Author.Firstname,
			&arr.Author.Lastname,
			&arr.CreatedAt)
		if err != nil {
			log.Panic(err)
		}
	}
	if arr.ID == 0 && arr.CreatedAt == nil {
		c.JSON(http.StatusNotFound, "Invalid id, not found")
	} else {
		c.JSON(http.StatusOK, arr)
	}

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
	delete, err := db.Exec(`
	DELETE from article WHERE id=$1`, id)
	delete1, err := db.Exec(`
	DELETE FROM author WHERE id=$1`, id)
	if err != nil {
		c.JSON(400, "Error")
	}
	delet, err := delete.RowsAffected()
	delet1, err := delete1.RowsAffected()
	if err != nil {
		c.JSON(422, "Server error")
	}
	if delet == 0 && delet1 == 0 {
		c.JSON(http.StatusNotFound, "Not found no id")
	} else {
		c.JSON(http.StatusOK, "Delete success")
	}

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
	//id := c.Param("id")

	var article models.Article
	err := c.BindJSON(&article)
	if err != nil {
		c.JSON(400, "Bad request")
	}
	query := (`UPDATE article SET title=$1,body=$2,updated_at=NOW() WHERE id=$3`)
	rows, err := db.Exec(
		query,
		article.Title,
		article.Body,
		article.ID,
	)
	if err != nil {
		c.JSON(422, "Error")
	}
	updated, err := rows.RowsAffected()
	if err != nil {
		panic(err)
	}
	query1 := (`UPDATE author SET firstname=$1,lastname=$2,updated_at=NOW() WHERE id=$3`)
	rows1, err := db.Exec(
		query1,
		article.Author.Firstname,
		article.Author.Lastname,
		article.ID,
	)
	if err != nil {
		c.JSON(422, "Error")
	}
	updated1, err := rows1.RowsAffected()
	if err != nil {
		panic(err)
	}
	if updated == 0 && updated1 == 0 {
		c.JSON(http.StatusNotFound, "Not found, No id")
	} else {
		c.JSON(http.StatusOK, "Succes update")
	}
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
