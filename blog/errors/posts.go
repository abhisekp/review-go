package errors

import (
	"fmt"
	"net/http"
)

type ErrAddPost struct {
	Message    string
	StatusCode int
}

func (self *ErrAddPost) Error() string {
	return self.Message
}

func ErrAddPostBadRequest(errs ...error) *ErrAddPost {
	if len(errs) > 0 {
		err := errs[0]
		err = fmt.Errorf("could not create post: %w", err)
		return &ErrAddPost{err.Error(), http.StatusBadRequest}
	}
	return &ErrAddPost{"could not create post", http.StatusBadRequest}

}

func ErrPostNotFound(postId string) *ErrAddPost {
	return &ErrAddPost{fmt.Errorf("post with id %s not found", postId).Error(), http.StatusNotFound}
}

var ErrAddPostInternalServerError = &ErrAddPost{"could not create post", http.StatusInternalServerError}
