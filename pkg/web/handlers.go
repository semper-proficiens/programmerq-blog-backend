package web

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/semper-proficiens/programmerq-blog-backend/blog/internal/entities"
	"github.com/semper-proficiens/programmerq-blog-backend/blog/internal/interfaces"
)

// BlogPostHandler handles HTTP requests related to blog posts
type BlogPostHandler struct {
	PostInteractor interfaces.BlogPostInteractor
}

// NewBlogPostHandler creates a new BlogPostHandler
func NewBlogPostHandler(postInteractor interfaces.BlogPostInteractor) *BlogPostHandler {
	return &BlogPostHandler{
		PostInteractor: postInteractor,
	}
}

// CreatePost handles the creation of a new blog post
func (h *BlogPostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post entities.BlogPost
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.PostInteractor.CreatePost(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetPost handles the retrieval of a blog post by ID
func (h *BlogPostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]

	post, err := h.PostInteractor.GetPostByID(postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(post)
}