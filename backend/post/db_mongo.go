package post

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongoDB struct {
	conn *mgo.Session
	c    *mgo.Collection
}

// NewMongoDB creates a new PostDatabase backed by a given Mongo server,
// authenticated with given credentials.
func newMongoDB(cfg Cfg) (PostDatabase, error) {
	conn, err := mgo.Dial(cfg.DbURI)
	if err != nil {
		return nil, fmt.Errorf("mongo: could not dial: %v", err)
	}
	cred := &mgo.Credential{
		Username: cfg.DbUser,
		Password: cfg.DbPassword,
	}
	if cred != nil {
		if err := conn.Login(cred); err != nil {
			return nil, err
		}
	}

	return &mongoDB{
		conn: conn,
		c:    conn.DB(cfg.DbName).C("posts"),
	}, nil
}

// Close closes the database.
func (db *mongoDB) Close() {
	db.conn.Close()
}

// Get retrieves a post by its ID.
func (db *mongoDB) Get(id int64) (*Post, error) {
	b := &Post{}
	if err := db.c.Find(bson.D{{Name: "id", Value: id}}).One(b); err != nil {
		return nil, err
	}
	return b, nil
}

var maxRand = big.NewInt(1<<63 - 1)

// randomID returns a positive number that fits within an int64.
func randomID() (int64, error) {
	// Get a random number within the range [0, 1<<63-1)
	n, err := rand.Int(rand.Reader, maxRand)
	if err != nil {
		return 0, err
	}
	// Don't assign 0.
	return n.Int64() + 1, nil
}

// Add saves a given post, assigning it a new ID.
func (db *mongoDB) Add(b *Post) (id int64, err error) {
	id, err = randomID()
	if err != nil {
		return 0, fmt.Errorf("mongodb: could not assign an new ID: %v", err)
	}

	b.ID = id
	if err := db.c.Insert(b); err != nil {
		return 0, fmt.Errorf("mongodb: could not add post: %v", err)
	}
	return id, nil
}

// Delete removes a given post by its ID.
func (db *mongoDB) Delete(id int64) error {
	return db.c.Remove(bson.D{{Name: "id", Value: id}})
}

// Update updates the entry for a given post.
func (db *mongoDB) Update(b *Post) error {
	return db.c.Update(bson.D{{Name: "id", Value: b.ID}}, b)
}

// List returns a list of posts, ordered by title.
func (db *mongoDB) List() ([]*Post, error) {
	var result []*Post
	if err := db.c.Find(nil).Sort("-lastupdate").All(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListByCreatedBy returns a list of posts, ordered by title, filtered by
// the user who created the post entry.
func (db *mongoDB) ListByCreatedBy(userID string) ([]*Post, error) {
	var result []*Post
	if err := db.c.Find(bson.D{{Name: "createdbyid", Value: userID}}).Sort("title").All(&result); err != nil {
		return nil, err
	}
	return result, nil
}
