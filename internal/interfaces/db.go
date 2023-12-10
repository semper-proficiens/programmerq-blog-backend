package interfaces

import "github.com/yourusername/blog/internal/entities"

// PostRepository defines the interface for interacting with the blog post repository
type PostRepository interface {
	CreatePost(post *entities.BlogPost) error
	GetPostByID(id string) (*entities.BlogPost, error)
}