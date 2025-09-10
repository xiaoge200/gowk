package gowk

import (
	"fmt"
)

type APIError struct {
	StatusCode int    `json:"-"`
	Status     int    `json:"status"`
	Msg        string `json:"msg"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Status, e.Msg)
}
