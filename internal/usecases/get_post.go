package usecases

import (
	"programmerq-blog-backend/internal/entities"
	"programmerq-blog-backend/internal/interfaces"
)

// GetPostUseCase handles the use case for retrieving a blog post
type GetPostUseCase struct {
	Repository interfaces.PostRepository
}

// Execute retrieves a blog post by ID
func (uc *GetPostUseCase) Execute(id string) (*entities.BlogPost, error) {
	return uc.Repository.GetPostByID(id)
}