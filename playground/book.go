package playground

import (
	"encoding/json"
	"io"
	"time"
)

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	CreateOn string `json:"-"`
	DeleteOn string `json:"-"`
	UpdateOn string `json:"-"`
}

type Books []*Book

func (p *Books) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetBooks() Books {
	return bookList
}

var bookList = []*Book{
	{
		ID:       1,
		Title:    "The Hitchhiker's Guide to the Galaxy",
		Author:   "Douglas Adams",
		CreateOn: time.Now().UTC().String(),
		UpdateOn: time.Now().UTC().String(),
	},
	{
		ID:       2,
		Title:    "Cloud Native Go",
		Author:   "M.-L. Reimer, K. M. Tsang",
		CreateOn: time.Now().UTC().String(),
		UpdateOn: time.Now().UTC().String(),
	},
}
