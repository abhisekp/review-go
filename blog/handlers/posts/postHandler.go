package posts

import (
	"database/sql"
	postsService "review-go/blog/services/posts"
)

type PostsHandler struct {
	postService *postsService.PostsService
}

type PostsHandlerOptions struct {
	DB *sql.DB
	_  struct{}
}

func NewPostsHandler(options PostsHandlerOptions) *PostsHandler {
	return &PostsHandler{
		postService: postsService.NewPostsService(postsService.PostsServiceOptions{
			DB: options.DB,
		}),
	}
}
