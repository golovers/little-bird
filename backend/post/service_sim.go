package post

import "gitlab.com/koffee/micro/config"

type simpleService struct {
	db PostDatabase
}

func newSimplePostService() PostService {
	var cfg Cfg
	config.LoadEnvConfig(&cfg)
	db, err := newMongoDB(cfg)
	if err != nil {
		panic(err)
	}
	return &simpleService{
		db: db,
	}
}

// List returns a list of posts, ordered by title.
func (s *simpleService) List() ([]*Post, error) {
	return s.db.List()
}

// ListByCreatedBy returns a list of posts, ordered by title, filtered by
// the user who created the post entry.
func (s *simpleService) ListByCreatedBy(userID string) ([]*Post, error) {
	return s.db.ListByCreatedBy(userID)
}

// Get retrieves a post by its ID.
func (s *simpleService) Get(id int64) (*Post, error) {
	return s.db.Get(id)
}

// Add saves a given post, assigning it a new ID.
func (s *simpleService) Add(b *Post) (id int64, err error) {
	return s.db.Add(b)
}

// Delete removes a given post by its ID.
func (s *simpleService) Delete(id int64) error {
	return s.db.Delete(id)
}

// Update updates the entry for a given post.
func (s *simpleService) Update(b *Post) error {
	return s.db.Update(b)
}
