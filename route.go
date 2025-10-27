package gowk

import (
	"context"
	"strconv"
)

type UserIMAddr struct {
	TcpAddr string `json:"tcp_addr"`
	WsAddr  string `json:"ws_addr"`
	WssAddr string `json:"wss_addr"`
}

// 获取用户IM地址
func (g *GoWk) GetUserIMAddr(ctx context.Context, intranet int) (*UserIMAddr, error) {
	var result UserIMAddr
	resp, err := g.restyClient.R().
		SetContext(ctx).
		SetResult(&result).
		SetQueryParam("intranet", strconv.Itoa(intranet)).
		Get("/route")

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

type BatchUserIMAddr struct {
	Uids []string `json:"uids"`
	UserIMAddr
}

// 批量获取用户IM地址
func (g *GoWk) BatchGetUserIMAddr(ctx context.Context, intranet int, userId ...string) ([]BatchUserIMAddr, error) {
	var result []BatchUserIMAddr
	resp, err := g.restyClient.R().
		SetContext(ctx).
		SetBody(userId).
		SetResult(&result).
		SetQueryParam("intranet", strconv.Itoa(intranet)).
		Post("/route/batch")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return result, nil
}
