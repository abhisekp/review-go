package posts

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	appErrors "review-go/blog/errors"
	"review-go/blog/handlers/utils"
	"review-go/blog/models"
	postsService "review-go/blog/services/posts"
)

type AddPostOptions struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (self *PostsHandler) AddPost(w http.ResponseWriter, r *http.Request) {
	addPostContext, cancelAddPost := context.WithCancel(r.Context())
	_, _ = addPostContext, cancelAddPost

	postData := models.Post{}

	err := json.NewDecoder(r.Body).Decode(&postData)
	if err != nil {
		fmt.Println(appErrors.ErrAddPostBadRequest(err))
		utils.SendErrorResponse(w, appErrors.ErrAddPostBadRequest())
		return
	}

	err = validateAddPostOptions(addPostContext, postData)
	if err != nil {
		utils.SendErrorResponse(w, appErrors.ErrAddPostBadRequest(err))
		return
	}

	postResult, err := self.postService.AddPost(addPostContext, postsService.AddPostOptions{
		Title:   postData.Title,
		Content: postData.Content,
	})

	_ = utils.WriteJsonResponse(w, utils.WriteJSONResponseOptions[*models.Post]{
		Data:       postResult,
		StatusCode: http.StatusCreated,
	})
}

var ErrTitleRequired = errors.New("post Title is required")

func validateAddPostOptions(ctx context.Context, options models.Post) error {
	if options.Title == "" {
		return ErrTitleRequired
	}
	return nil
}
