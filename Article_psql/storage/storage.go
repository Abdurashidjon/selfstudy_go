package storage

import (
	"errors"

	"gitlab.com/Udevs/Article/models"
)

type ArticleStorage struct {
	data map[int]models.Article
}

// Errors
var ErrorAlreadyExists = errors.New("already exists")
var ErrorNotFound = errors.New("Not found ")
var ErrorUpdate = errors.New("Not update data by id")
var ErrorDelete = errors.New("Succes delete")
var ErrorNotUpdate = errors.New("Error not found id")

// Succesfull
var SuccessUpdate = errors.New("Update succesful")

func NewArticleStorage() ArticleStorage {
	return ArticleStorage{
		data: make(map[int]models.Article),
	}
}

func (storage *ArticleStorage) Add(entity models.Article) error {
	// code
	if _, ok := storage.data[entity.ID]; ok {
		return ErrorAlreadyExists
	}
	storage.data[entity.ID] = entity
	return nil
}

func (storage *ArticleStorage) GetByID(ID int) (resp models.Article, err error) {
	// code
	var ok bool
	if resp, ok = storage.data[ID]; !ok {
		return resp, ErrorNotFound
	}
	return resp, err
}

func (storage *ArticleStorage) GetList() []models.Article {
	// code
	var res []models.Article
	for _, v := range storage.data {
		res = append(res, storage.data[v.ID])
	}
	return res
}

func (storage *ArticleStorage) Search(str string) ([]models.Article, int) {
	//code
	var res []models.Article
	count := 0
	for _, v := range storage.data {
		if str == v.Body {
			//res = v  //bu hammasini chiqarib bermidi faqat 1-sini chiqarib bera oladi
			res = append(res, storage.data[v.ID]) // bu hammasini chiqarib beradi va array ichiga olIPadi
			count += 1
		}
	}
	return res, count
}

func (storage *ArticleStorage) Update(entity models.Article) (error) {
	for index, item := range storage.data {
		if item.ID == entity.ID {
			storage.data[index] = entity
			return SuccessUpdate
		}
	}
	return ErrorNotUpdate
}

func (storage *ArticleStorage) Delete(ID int) error {
	var ok bool
	if _, ok = storage.data[ID]; !ok {
		return ErrorNotFound
	}
	delete(storage.data, ID)
	return ErrorDelete
}
