package relpaginator

import (
	"fmt"
	"net/http"
)

type PageError struct {
	PageNumber uint
}

func (c PageError) Error() string {
	return fmt.Sprintf("Page %v does not exists", c.PageNumber)
}

func (c PageError) StatusCode() int {
	return http.StatusBadRequest
}
