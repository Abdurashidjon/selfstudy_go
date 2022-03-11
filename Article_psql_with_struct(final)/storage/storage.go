package storage

import "gitlab.com/Udevs/Article/models"

type StorageI interface {
	Article() ArticleRepoI
	Author() AuthorRepoI
}

type ArticleRepoI interface {
	CreateArticle(entity models.Article) (err error)
	GetListArticle(entity models.Query) (resp []models.Article, err error)
	GetByIdArticle(ID int) (resp models.Article, err error)
	UpdateByIdArticle(entity models.Article) (resp models.Article, err error)
	DeleteByIdArticle(ID int) (err error)
}

type AuthorRepoI interface {
	CreateArticle(entity models.Person) (err error)
	GetListArticle(entity models.Query) (resp []models.Person, err error)
	GetByIdArticle(ID int) (resp models.Person, err error)
	UpdateByIdArticle(entity models.Person) (resp models.Person, err error)
	DeleteByIdArticle(ID int) (err error)
}
