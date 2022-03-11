package handlers

import "gitlab.com/Udevs/Article/storage"

type Handler struct {
	strg storage.StorageI
}

func NewHandler(strg storage.StorageI) Handler {
	return Handler{
		strg: strg,
	}
}