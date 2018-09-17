package comments

import (
	"fmt"

	"gitlab.com/koffee/little-bird/backend/core"
	"gitlab.com/koffee/micro/config"
	"gopkg.in/mgo.v2"
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
	config.LoadEnvConfig(&cfg)

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
	//TODO implement me
	return []*core.Comment{}, nil
}

func (db *mongoDB) Create(b *core.Comment) (id string, err error) {
	//TODO implement me
	return "", nil
}
func (db *mongoDB) Delete(id string) error {
	//TODO implement me
	return nil
}

func (db *mongoDB) Update(b *core.Comment) error {
	//TODO implement me
	return nil
}

func (db *mongoDB) Close() {
	//TODO implement me
}
