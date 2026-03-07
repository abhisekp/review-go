package utils

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	appErrors "review-go/blog/errors"
)

func SendErrorResponse(w http.ResponseWriter, err error) {
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		if err_, ok := errors.AsType[*appErrors.ErrAddPost](err); ok {
			w.WriteHeader(err_.StatusCode)
			if err = json.NewEncoder(w).Encode(struct {
				Message    string `json:"message"`
				StatusCode int    `json:"statusCode"`
			}{
				Message:    err_.Error(),
				StatusCode: err_.StatusCode,
			}); err != nil {
				fmt.Println(fmt.Errorf("could not write response: %w", err))
				return
			}
		} else if errors.As(err, &sql.ErrNoRows) {
			w.WriteHeader(http.StatusNotFound)
			if err = json.NewEncoder(w).Encode(struct {
				Message    string `json:"message"`
				StatusCode int    `json:"statusCode"`
			}{
				Message:    err.Error(),
				StatusCode: http.StatusNotFound,
			}); err != nil {
				fmt.Println(fmt.Errorf("could not write response: %w", err))
				return
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			if err = json.NewEncoder(w).Encode(struct {
				Message    string `json:"message"`
				StatusCode int    `json:"statusCode"`
			}{
				Message:    "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			}); err != nil {
				fmt.Println(fmt.Errorf("could not write response: %w", err))
				return
			}
		}
		return
	}
}

type WriteJSONResponseOptions[T any] struct {
	Data       T
	StatusCode int
}

func WriteJsonResponse[T any](w http.ResponseWriter, option WriteJSONResponseOptions[T]) error {
	w.Header().Set("Content-Type", "application/json")
	statusCode := option.StatusCode
	if statusCode == 0 {
		statusCode = http.StatusOK
	}
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(option.Data)
	if err != nil {
		fmt.Println(fmt.Errorf("could not write json data to response: %w", err))
		return err
	}
	return nil
}
