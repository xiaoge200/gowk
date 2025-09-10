package gowk

import "strconv"

type UserIMAddr struct {
	TcpAddr string `json:"tcp_addr"`
	WsAddr  string `json:"ws_addr"`
	WssAddr string `json:"wss_addr"`
}

// 获取用户IM地址
func (g *GoWk) GetUserIMAddr(intranet int) (*UserIMAddr, error) {
	var result UserIMAddr
	resp, err := g.restyClient.R().
		SetResult(&result).
		Get("/route?intranet=" + strconv.Itoa(intranet))

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
func (g *GoWk) BatchGetUserIMAddr(intranet int, userId ...string) ([]BatchUserIMAddr, error) {
	var result []BatchUserIMAddr
	resp, err := g.restyClient.R().
		SetBody(userId).
		SetResult(&result).
		Post("/route/batch?intranet=" + strconv.Itoa(intranet))
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return result, nil
}
