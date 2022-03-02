package storage

import "gitlab.com/Udevs/Article/models"

type ArticleRepoI interface {
	CreateArticle(entity models.Article ) (err error)
	GetListArticle(entity models.Query) (resp []models.Article,err error)
	GetByIdArticle(ID int) (resp models.Article,err error)
	UpdateByIdArticle(entity models.Article) (resp models.Article,err error)
	DeleteByIdArticle(ID int) (err error)
}