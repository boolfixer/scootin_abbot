package http_error

import (
	"fmt"
)

type NotFoundError struct {
	ModelName string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s not found.", e.ModelName)
}
