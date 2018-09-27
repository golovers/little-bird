package comments

import (
	"fmt"

	"github.com/golovers/little-bird/backend/core"
	"github.com/rs/xid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var _ Repository = &mongoDB{}

type mongoDB struct {
	conn *mgo.Session
	c    *mgo.Collection
}

type mongCfg struct {
	DbURI      string `envconfig:"COMMENT_DB_URI"`
	DbName     string `envconfig:"COMMENT_DB_NAME"`
	DbUser     string `envconfig:"COMMENT_DB_USERNAME"`
	DbPassword string `envconfig:"COMMENT_DB_PASSWORD"`
}

// NewMongoDB creates a new article repositoryy backed by a given Mongo server,
// authenticated with given credentials.
func newMongoDB() (*mongoDB, error) {
	var cfg mongCfg
	core.LoadEnvConfig(&cfg)

	conn, err := mgo.Dial(cfg.DbURI)
	if err != nil {
		return nil, fmt.Errorf("mongo: could not dial: %v", err)
	}
	if cfg.DbUser != "" {
		cred := &mgo.Credential{
			Username: cfg.DbUser,
			Password: cfg.DbPassword,
		}
		if cred != nil {
			if err := conn.Login(cred); err != nil {
				return nil, err
			}
		}
	}
	return &mongoDB{
		conn: conn,
		c:    conn.DB(cfg.DbName).C("comments"),
	}, nil
}

func (db *mongoDB) ListByArticle(articleID string) ([]*core.Comment, error) {
	var rs []*core.Comment
	err := db.c.Find(bson.D{{Name: "articleid", Value: articleID}}).All(&rs)
	if err != nil {
		return rs, err
	}
	return rs, nil
}

func (db *mongoDB) Create(c *core.Comment) (id string, err error) {
	c.ID = randomID()
	if err := db.c.Insert(c); err != nil {
		return "", fmt.Errorf("mongodb: could not create new comment: %v", err)
	}
	return c.ID, nil
}
func (db *mongoDB) Delete(id string) error {
	return db.c.Remove(bson.D{{Name: "id", Value: id}})
}

func (db *mongoDB) Update(c *core.Comment) error {
	return db.c.Update(bson.D{{Name: "id", Value: c.ID}}, c)
}

func (db *mongoDB) Get(id string) (*core.Comment, error) {
	c := &core.Comment{}
	if err := db.c.Find(bson.D{{Name: "id", Value: id}}).One(c); err != nil {
		return nil, err
	}
	return c, nil
}

func (db *mongoDB) Close() {
	//TODO implement me
}

// randomID returns a positive number that fits within an int64.
func randomID() string {
	return xid.New().String()
}
