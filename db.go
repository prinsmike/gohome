package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Mongo struct {
	URL      string
	DbStr    string
	Session  *mgo.Session
	Database *mgo.Database
}

// A convenient alias for bson.M so we don't have to import bson everywhere.
type M bson.M

func (mongo *Mongo) GetSession() *mgo.Session {
	if mongo.Session == nil {
		var err error
		mongo.Session, err = mgo.Dial(mongo.URL)
		if err != nil {
			panic(err)
		}
	}
	return mongo.Session.Clone()
}

func (mongo *Mongo) GetDb() *mgo.Database {
	return mongo.Session.DB(mongo.DbStr)
}

func NewMongoConnection(url, db string) (session *mgo.Session, database *mgo.Database) {
	mongo := new(Mongo)
	mongo.URL = url
	mongo.DbStr = db
	session = mongo.GetSession()
	database = mongo.GetDb()
	return
}
