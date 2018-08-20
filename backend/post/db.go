package api

// PostDatabase provides thread-safe access to a database of posts.
type PostDatabase interface {
	// List returns a list of posts, ordered by title.
	List() ([]*Post, error)

	// ListByCreatedBy returns a list of posts, ordered by title, filtered by
	// the user who created the post entry.
	ListByCreatedBy(userID string) ([]*Post, error)

	// Get retrieves a post by its ID.
	Get(id int64) (*Post, error)

	// Add saves a given post, assigning it a new ID.
	Add(b *Post) (id int64, err error)

	// Delete removes a given post by its ID.
	Delete(id int64) error

	// Update updates the entry for a given post.
	Update(b *Post) error

	// Close closes the database, freeing up any available resources.
	// TODO(cbro): Close() should return an error.
	Close()
}
