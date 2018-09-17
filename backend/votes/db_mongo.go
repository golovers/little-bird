package votes

import (
	"fmt"

	"github.com/go-mgo/mgo"
	"github.com/go-mgo/mgo/bson"
	"github.com/rs/xid"
	"gitlab.com/koffee/little-bird/backend/core"
)

var _ Repository = &mongoDB{}

type mongoDB struct {
	conn *mgo.Session
	c    *mgo.Collection
}

type mongCfg struct {
	DbURI      string `envconfig:"VOTE_DB_URI"`
	DbName     string `envconfig:"VOTE_DB_NAME"`
	DbUser     string `envconfig:"VOTE_DB_USERNAME"`
	DbPassword string `envconfig:"VOTE_DB_PASSWORD"`
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
		c:    conn.DB(cfg.DbName).C("votes"),
	}, nil
}

func (db *mongoDB) ListByArticle(articleID string) ([]*core.Vote, error) {
	var result []*core.Vote
	if err := db.c.Find(nil).Sort("-lastupdate").All(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (db *mongoDB) Create(v *core.Vote) (id string, err error) {
	id = randomID()

	v.ID = id
	if err := db.c.Insert(v); err != nil {
		return id, fmt.Errorf("mongodb: could not create vote: %v", err)
	}
	return id, nil
}

func (db *mongoDB) Get(id string) (*core.Vote, error) {
	var v core.Vote
	if err := db.c.FindId(bson.D{{Name: "id", Value: id}}).One(&v); err != nil {
		return &core.Vote{}, err
	}
	return &v, nil
}

func (db *mongoDB) GetByArticleAndUserID(articleID string, userID string) (*core.Vote, error) {
	var v core.Vote
	if err := db.c.Find(bson.D{{Name: "articleid", Value: articleID}, {Name: "createdbyid", Value: userID}}).One(&v); err != nil {
		return &core.Vote{}, err
	}
	return &v, nil
}

func (db *mongoDB) Close() {
	db.conn.Clone()
}

// randomID returns a positive number that fits within an int64.
func randomID() string {
	return xid.New().String()
}
