package posts

import (
	"context"
	"encoding/json"
	"review-go/blog/models"
	"review-go/blog/services/db"
)

func (self *PostsService) GetPostById(ctx context.Context, id string) (*models.Post, error) {
	var raw []byte
	if err := db.PrepStmts.Statements["getPostById"].QueryRowContext(ctx, id).Scan(&raw); err != nil {
		return nil, err
	}

	var post models.Post
	if err := json.Unmarshal(raw, &post); err != nil {
		return nil, err
	}
	return &post, nil
}
