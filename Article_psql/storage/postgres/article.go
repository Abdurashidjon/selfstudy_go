package postgres

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/Udevs/Article/models"
)

type articleRepo struct {
	db *sqlx.DB
}

func (r articleRepo) Create(entity models.Article) (err error) {
	return
}
