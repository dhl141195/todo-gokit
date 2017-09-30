package db

import mgo "gopkg.in/mgo.v2"

type Mongo struct {
	Conn *mgo.Session
	DB   *mgo.Database
}

func NewMongo(c *mgo.Session, dbname string) *Mongo {
	return &Mongo{
		Conn: c,
		DB:   c.DB(dbname),
	}
}
