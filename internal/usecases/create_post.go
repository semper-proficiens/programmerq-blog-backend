package usecases

import (
	"programmerq-blog-backend/internal/entities"
	"programmerq-blog-backend/internal/interfaces"
)

// CreatePostUseCase handles the use case for creating a blog post
type CreatePostUseCase struct {
	Repository interfaces.PostRepository
}

// Execute creates a new blog post
func (uc *CreatePostUseCase) Execute(post *entities.BlogPost) error {
	return uc.Repository.CreatePost(post)
}