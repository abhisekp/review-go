package posts

import (
	"database/sql"
	"log"
)

type PostsService struct {
	db *sql.DB
}

type PostsServiceOptions struct {
	DB *sql.DB
	_  struct{}
}

func NewPostsService(options PostsServiceOptions) *PostsService {
	if options.DB == nil {
		log.Fatalln("DB connection is required for PostsService")
	}

	return &PostsService{
		db: options.DB,
	}
}
