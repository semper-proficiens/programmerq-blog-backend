package interfaces

import "programmerq-blog-backend/internal/entities"

// BlogPostInteractor defines the interface for blog post use cases
type BlogPostInteractor interface {
	CreatePost(post *entities.BlogPost) error
	GetPostByID(id string) (*entities.BlogPost, error)
}
