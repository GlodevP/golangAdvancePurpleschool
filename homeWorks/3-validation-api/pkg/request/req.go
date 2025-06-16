package request

import (
	"3-validation-api/pkg/response"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := decodeBody[T](r.Body)
	if err != nil {
		response.Json(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	err = validateBody(body)
	if err != nil {
		response.Json(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	return body, nil
}

func decodeBody[T any](body io.ReadCloser) (*T, error) {
	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func validateBody[T any](payload T) error {
	validate := validator.New()
	err := validate.Struct(payload)
	return err
}
