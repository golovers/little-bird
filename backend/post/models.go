package api

import "time"

// Post holds metadata about a book.
type Post struct {
	ID         int64
	Title      string
	Content    string
	LastUpdate time.Time

	CreatedBy   string
	CreatedByID string
}

// CreatedByDisplayName returns a string appropriate for displaying the name of
// the user who created this book object.
func (b *Post) CreatedByDisplayName() string {
	if b.CreatedByID == "anonymous" {
		return "Anonymous"
	}
	return b.CreatedBy
}

// SetCreatorAnonymous sets the CreatedByID field to the "anonymous" ID.
func (b *Post) SetCreatorAnonymous() {
	b.CreatedBy = ""
	b.CreatedByID = "anonymous"
}
