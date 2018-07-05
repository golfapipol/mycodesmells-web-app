package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Mongo struct {
	Session *mgo.Session
	dbName  string
}

func NewMongo(dbURI, dbName string) (Mongo, error) {
	session, err := mgo.Dial(dbURI)

	return Mongo{
		Session: session,
		dbName:  dbName,
	}, err
}

func (m Mongo) db() *mgo.Database {
	session := m.Session.Copy()
	return session.DB(m.dbName)
}

func (m Mongo) Countries() ([]Country, error) {
	var cs []Country
	err := m.db().C("countries").Find(bson.M{}).All(&cs)
	return cs, err
}
