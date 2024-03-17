package bookstore

import "slices"

type Book struct {
	Author string
	Name   string
	Copies int32
	ID     int32
}

func GetAllBooks(catalog []Book) []Book {
	return catalog
}

func GetBook(catalog []Book, id int32) Book {
	foundId := slices.IndexFunc(catalog, func(v Book) bool {
		return v.ID == 3
	})
	return catalog[foundId]
}

func GetMap() map[string]Book {
	var bookstoreMap map[string]Book
	bookstoreMap["hi"] = Book{}
	return bookstoreMap
}
