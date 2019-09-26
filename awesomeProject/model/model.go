package model

import (
	"gopkg.in/mgo.v2"
)

// Model -- constructor struct contains db connection
type Model struct {
	db *mgo.Database
}

// New -- model constructor
func New(conn *mgo.Session, name string) *Model {
	return &Model{
		db: conn.DB(name),
	}
}
