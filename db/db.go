package db

import (
	"strings"

	mgo "gopkg.in/mgo.v2"
)

type Query struct {
	Filter  map[string]interface{} `msgpack:"f,omitempty"`
	OrderBy []string               `msgpack:"o,omitempty"` // +|-{fieldName}
	Limit   int                    `msgpack:"l,omitempty"`
	Cursor  string                 `msgpack:"-"`
}

type Cursor struct {
	IsPrev bool        `msgpack:"p,omitempty"`
	ID     interface{} `msgpack:"i,omitempty"`
	Query  Query       `msgpack:"q,omitempty"`
}

type PagingCursor struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}

type Paging struct {
	Total        int `json:"total"`
	PagingCursor `json:"cursor"`
}

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

func (q *Query) GetLimit() int {
	if q == nil {
		return 0
	}
	return q.Limit
}

func (q *Query) GetFilter() map[string]interface{} {
	if q == nil {
		return nil
	}
	return q.Filter
}

// OrderByMaps transforms OrderBy array to array of map
// [-field1, +field2] => [{field1: "-"}, {field2: "+"}]
func (q *Query) OrderByMaps() []map[string]string {
	if q == nil || q.OrderBy == nil {
		return nil
	}
	arr := []map[string]string{}
	fmap := map[string]string{}
	for _, o := range q.OrderBy {
		o = strings.TrimSpace(o)
		if o == "" {
			continue
		}
		if _, exist := fmap[o]; exist {
			continue
		}
		order := "+"
		if od := string(o[0]); od == "-" || od == "+" {
			order = od
			o = string(o[1:])
			if o == "" {
				continue
			}
		}
		fmap[o] = order
		arr = append(arr, map[string]string{o: order})
	}
	return arr
}
