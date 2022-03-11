package postgres

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/Udevs/Article/models"
	"gitlab.com/Udevs/Article/storage"
)

type authorRepo struct {
	db *sqlx.DB
}

func NewAuthorRepo(db *sqlx.DB) storage.AuthorRepoI {
	return &authorRepo{db: db}
}

func (r *authorRepo) CreateArticle(entity models.Person) (err error) {
	return
}

func (r *authorRepo) GetListArticle(entity models.Query) (resp []models.Person, err error) {
	return
}

func (r *authorRepo) GetByIdArticle(ID int) (resp models.Person, err error) {
	return
}

func (r *authorRepo) UpdateByIdArticle(entity models.Person) (resp models.Person, err error) {
	return
}

func (r *authorRepo) DeleteByIdArticle(ID int) (err error) {
	return
}
