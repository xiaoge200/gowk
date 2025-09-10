package gowk

type ConnInfo struct {
	UID    string `json:"uid"`               // 必传，用户 ID
	ConnID int64  `json:"conn_id"`           // 必传，连接 ID
	NodeID *int64 `json:"node_id,omitempty"` // 可选，指定节点 ID
}

// 移除连接
func (g *GoWk) RemoveConn(req ConnInfo) (*StatusResponse, error) {
	var result StatusResponse
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/conn/remove")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

// 踢出连接
func (g *GoWk) KickConn(req ConnInfo) (*StatusResponse, error) {
	var result StatusResponse
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/conn/kick")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}
