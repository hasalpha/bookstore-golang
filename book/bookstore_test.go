package bookstore_test

import (
	"bookstore"
	"github.com/google/go-cmp/cmp"
	"testing"
)

// func TestGetAllBooks(t *testing.T) {
// 	t.Parallel()
// 	catalog := []bookstore.Book{
// 		{
// 			Author: "Mohammed Arshad Hassan",
// 			Name:   "Be focused",
// 			Copies: 3,
// 			ID:     1,
// 		},
// 		{
// 			Author: "Mohammed Arshad Hassan",
// 			Name:   "Be disciplined",
// 			Copies: 20,
// 			ID:     2,
// 		},
// 		{
// 			Author: "Mohammed Arshad Hassan",
// 			Name:   "Be hardworking",
// 			Copies: 5,
// 			ID:     3,
// 		},
// 	}
// 	want := catalog
// 	got := bookstore.GetAllBooks(catalog)
// 	if !cmp.Equal(want, got) {
// 		t.Errorf(cmp.Diff(want, got))
// 		return
// 	}
// }
//
// func TestGetBook(t *testing.T) {
// 	t.Parallel()
// 	catalog := []bookstore.Book{
// 		{
// 			Author: "Mohammed Arshad Hassan",
// 			Name:   "Be focused",
// 			Copies: 3,
// 			ID:     1,
// 		},
// 		{
// 			Author: "Mohammed Arshad Hassan",
// 			Name:   "Be disciplined",
// 			Copies: 20,
// 			ID:     2,
// 		},
// 		{
// 			Author: "Mohammed Arshad Hassan",
// 			Name:   "Be hardworking",
// 			Copies: 5,
// 			ID:     3,
// 		},
// 	}
// 	want := catalog[len(catalog)-1]
// 	got := bookstore.GetBook(catalog, 3)
// 	if !cmp.Equal(want, got) {
// 		t.Errorf(cmp.Diff(want, got))
// 	}
// }

func TestGetMap(t *testing.T) {
	t.Parallel()
	bookstore.GetMap()
	cmp.Equal(1, 2)
}
