package post

type PostService interface {
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
}

var postService = newSimplePostService()

// List returns a list of posts, ordered by title.
func List() ([]*Post, error) {
	return postService.List()
}

// ListByCreatedBy returns a list of posts, ordered by title, filtered by
// the user who created the post entry.
func ListByCreatedBy(userID string) ([]*Post, error) {
	return postService.ListByCreatedBy(userID)
}

// Get retrieves a post by its ID.
func Get(id int64) (*Post, error) {
	return postService.Get(id)
}

// Add saves a given post, assigning it a new ID.
func Add(b *Post) (id int64, err error) {
	return postService.Add(b)
}

// Delete removes a given post by its ID.
func Delete(id int64) error {
	return postService.Delete(id)
}

// Update updates the entry for a given post.
func Update(b *Post) error {
	return postService.Update(b)
}
