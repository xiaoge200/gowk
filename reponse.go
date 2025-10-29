package gowk

import (
	"fmt"
	"net/http"
)

type StatusResponse struct {
	StatusCode int    `json:"-"`
	Status     int    `json:"status"` // 操作状态，成功时返回 "ok"
	Msg        string `json:"msg"`
	Data       any    `json:"data"`
}

func (r *StatusResponse) IsOk() bool {
	return r.Status == http.StatusOK
}

func (r *StatusResponse) Error() string {
	return fmt.Sprintf("[%d] %s", r.Status, r.Msg)
}
