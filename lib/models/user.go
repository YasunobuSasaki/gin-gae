package models

type User struct {
	Id 		string `datastore:"-" goon:"id"`
	Name     string `datastore:"Name,noindex"`
}

