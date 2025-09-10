package gowk

import "errors"

type StatusResponse struct {
	Status  string `json:"status"` // 操作状态，成功时返回 "ok"
	Message string `json:"message"`
}

func (r *StatusResponse) IsOk() bool {
	return r.Status == "ok"
}

func (r *StatusResponse) Error() error {
	if r.IsOk() {
		return nil
	}
	return errors.New(r.Message)
}
