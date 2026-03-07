package posts

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	appErrors "review-go/blog/errors"
	"review-go/blog/handlers/utils"
	"review-go/blog/models"
)

func (self *PostsHandler) GetById(w http.ResponseWriter, r *http.Request) {
	// Get id from the param
	postId := r.PathValue("id")

	getPostCtx, cancelPostCtx := context.WithCancel(r.Context())
	_, _ = getPostCtx, cancelPostCtx

	post, err := self.postService.GetPostById(getPostCtx, postId)
	if err != nil {
		if errors.As(sql.ErrNoRows, &err) {
			err = appErrors.ErrPostNotFound(postId)
		}
		utils.SendErrorResponse(w, err)
		return
	}

	_ = utils.WriteJsonResponse[*models.Post](w, utils.WriteJSONResponseOptions[*models.Post]{
		Data: post,
	})
}
