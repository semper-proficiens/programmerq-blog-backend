package usecases

import (
	"github.com/semper-proficiens/programmerq-blog-backend/blog/internal/entities"
	"github.com/semper-proficiens/programmerq-blog-backend/blog/internal/interfaces"
)

// GetPostUseCase handles the use case for retrieving a blog post
type GetPostUseCase struct {
	Repository interfaces.PostRepository
}

// Execute retrieves a blog post by ID
func (uc *GetPostUseCase) Execute(id string) (*entities.BlogPost, error) {
	return uc.Repository.GetPostByID(id)
}