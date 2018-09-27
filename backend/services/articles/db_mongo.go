package articles

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
	DbURI      string `envconfig:"ARTICLE_DB_URI"`
	DbName     string `envconfig:"ARTICLE_DB_NAME"`
	DbUser     string `envconfig:"ARTICLE_DB_USERNAME"`
	DbPassword string `envconfig:"ARTICLE_DB_PASSWORD"`
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
		c:    conn.DB(cfg.DbName).C("articles"),
	}, nil
}

// Close closes the database.
func (db *mongoDB) Close() {
	db.conn.Close()
}

// Get retrieves an article by its ID.
func (db *mongoDB) Get(id string) (*core.Article, error) {
	a := &core.Article{}
	if err := db.c.Find(bson.D{{Name: "id", Value: id}}).One(a); err != nil {
		return nil, err
	}
	return a, nil
}

// Create saves a given article, assigning it a new ID.
func (db *mongoDB) Create(a *core.Article) (id string, err error) {
	a.ID = randomID()
	if err := db.c.Insert(a); err != nil {
		return "", fmt.Errorf("mongodb: could not create new article: %v", err)
	}
	return a.ID, nil
}

// Delete removes a given article by its ID.
func (db *mongoDB) Delete(id string) error {
	return db.c.Remove(bson.D{{Name: "id", Value: id}})
}

// Update updates the entry for a given article.
func (db *mongoDB) Update(a *core.Article) error {
	return db.c.Update(bson.D{{Name: "id", Value: a.ID}}, a)
}

// List returns a list of articles, ordered by last update.
func (db *mongoDB) List(offset, limit int64) ([]*core.Article, error) {
	//TODO implement paging solution
	var result []*core.Article
	if err := db.c.Find(nil).Sort("-lastupdate").All(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListByCreatedBy returns a list of article, ordered by last update, filtered by
// the user who created the article.
func (db *mongoDB) ListByCreatedBy(userID string) ([]*core.Article, error) {
	var result []*core.Article
	if err := db.c.Find(bson.D{{Name: "createdbyid", Value: userID}}).Sort("-lastupdate").All(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// randomID returns a positive number that fits within an int64.
func randomID() string {
	return xid.New().String()
}
