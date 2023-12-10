package usecases

import (
	"github.com/yourusername/blog/internal/entities"
	"github.com/yourusername/blog/internal/interfaces"
)

// GetPostUseCase handles the use case for retrieving a blog post
type GetPostUseCase struct {
	Repository interfaces.PostRepository
}

// Execute retrieves a blog post by ID
func (uc *GetPostUseCase) Execute(id string) (*entities.BlogPost, error) {
	return uc.Repository.GetPostByID(id)
}