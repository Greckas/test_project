package model

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID          bson.ObjectId `bson:"_id" json:"-"`
	Login       string        `bson:"login" json:"login"`
	AccessLevel string        `bson:"accessLevel" json:"accessLevel"`
	Name        string        `bson:"name" json:"name"`
	Password    string        `bson:"password" json:"password"`
	Position    string        `bson:"position" json:"position"`
	Type        string        `bson:"type" json:"type"`
}

const usersDB = "users"

func (m *Model) Login(username string) (*User, error) {
	user := User{}
	err := m.db.C(usersDB).Find(bson.M{"login": username}).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
