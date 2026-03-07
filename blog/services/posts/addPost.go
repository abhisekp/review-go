package posts

import (
	"context"
	"errors"
	"fmt"
	appError "review-go/blog/errors"
	"review-go/blog/models"
)

type AddPostOptions struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (self *PostsService) AddPost(ctx context.Context, options AddPostOptions) (*models.Post, error) {
	var id int64
	err := self.db.QueryRowContext(ctx, "INSERT INTO public.posts (title, content) VALUES ($1, $2) RETURNING id", options.Title, options.Content).Scan(&id)
	if err != nil {
		fmt.Println(errors.Join(appError.ErrAddPostInternalServerError, err))
		return nil, appError.ErrAddPostInternalServerError
	}

	post := &models.Post{
		ID:      id,
		Title:   options.Title,
		Content: options.Content,
	}

	return post, nil
}
