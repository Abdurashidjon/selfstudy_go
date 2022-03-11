package postgres

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"gitlab.com/Udevs/Article/models"
	"gitlab.com/Udevs/Article/storage"
)

type articleRepo struct {
	db *sqlx.DB
}

func NewArticleRepo(db *sqlx.DB) storage.ArticleRepoI {
	return &articleRepo{db: db}
}

func (r articleRepo) CreateArticle(entity models.Article) (err error) {
	query := `WITH author AS(INSERT INTO author(firstname,lastname) VALUES($1,$2) RETURNING id)
				INSERT INTO article(title,body,author_id,created_at) VALUES($3,$4,(SELECT id FROM author),NOW())`

	rows, err := r.db.Query(
		query,
		entity.Title,
		entity.Body,
		entity.Author.Firstname,
		entity.Author.Lastname,
	)
	defer rows.Close()
	if err != nil {
		return err
	}
	return err
}

func (r articleRepo) GetListArticle(entity models.Query) (resp []models.Article, err error) {
	rows, err := r.db.Query(
		`SELECT
		ar.id,
		ar.title,
		ar.body,
		au.firstname,
		au.lastname,
		ar.created_at FROM article AS ar JOIN author AS au ON ar.author_id = au.id OFFSET $1 LIMIT $2`,
		entity.Offset,
		entity.Limit,
		//entity.Search,
	)
	if err != nil {
		return resp, err
	}

	defer rows.Close()
	for rows.Next() {
		var a models.Article
		err = rows.Scan(
			&a.ID,
			&a.Body,
			&a.Title,
			&a.Author.Firstname,
			&a.Author.Lastname,
			&a.CreatedAt)
		resp = append(resp, a)
		if err != nil {
			return resp, err
		}
	}
	return resp, err
}

func (r articleRepo) GetByIdArticle(ID int) (resp models.Article, err error) {
	rows, err := r.db.Query(
		`SELECT
				ar.id,
				ar.title,
				ar.body,
				au.firstname,
				au.lastname,
				ar.created_at
				FROM article AS ar  JOIN author AS au ON (ar.author_id = au.id) WHERE ar.id=$1`, ID)
	if err != nil {
		return resp, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&resp.ID,
			&resp.Body,
			&resp.Title,
			&resp.Author.Firstname,
			&resp.Author.Lastname,
			&resp.CreatedAt)
		if err != nil {
			return resp, err
		}
	}
	if resp.ID == 0 && resp.CreatedAt == nil {
		return resp, err
	} else {
		return resp, nil
	}
}

func (r articleRepo) UpdateByIdArticle(entity models.Article) (resp models.Article, err error) {
	query := `WITH article AS(UPDATE article 	
				SET title=$1,body=$2,updated_at=NOW() WHERE id=$3 RETURNING author_id)
					UPDATE author SET firstname=$4,lastname=$5,updated_at=NOW()	
					WHERE id = (SELECT author_id FROM article)`
	res, err := r.db.Exec(
		query,
		entity.Title,
		entity.Body,
		entity.ID,
		entity.Author.Firstname,
		entity.Author.Lastname)
	if err != nil {
		return resp, err
	}
	if i, _ := res.RowsAffected(); i == 0 {
		return resp, sql.ErrNoRows
	}

	resp.Title = entity.Title
	resp.Body = entity.Body
	resp.Author = entity.Author
	resp.ID = entity.ID
	resp.CreatedAt = entity.CreatedAt

	return resp, nil
}

func (r articleRepo) DeleteByIdArticle(ID int) (err error) {
	delete, err := r.db.Exec(`
		WITH article AS(DELETE FROM article WHERE id=$1 RETURNING author_id)
			DELETE FROM author WHERE id=(SELECT author_id from article)`, ID)
	if err != nil {
		return err
	}
	if i, _ := delete.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}
	return nil

}
