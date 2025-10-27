package gowk

import "context"

// SendEventRequest 发送事件请求
type SendEventRequest struct {
	ClientMsgNo string      `json:"client_msg_no"` // 必传，客户端消息编号，必须唯一
	ChannelID   string      `json:"channel_id"`    // 必传，目标频道ID
	ChannelType ChannelType `json:"channel_type"`  // 必传，频道类型 (1=个人, 2=群组)
	FromUid     string      `json:"from_uid"`      // 发送者UID
	Event       Event       `json:"event"`         // 必传，事件对象
}

// EventData 事件对象
type Event struct {
	Type      string  `json:"type"`                // 必传，事件类型
	ID        *string `json:"id,omitempty"`        // 可选，事件 ID
	Timestamp *int64  `json:"timestamp,omitempty"` // 可选，事件时间戳（毫秒）
	Data      string  `json:"data"`                // 必传，事件数据内容
}

// 发送事件
func (g *GoWk) SendEvent(ctx context.Context, force *string, req ConnInfo) (*StatusResponse, error) {
	var result StatusResponse
	r := g.restyClient.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&result)
	if force != nil {
		r.SetQueryParam("force", *force)
	}
	resp, err := r.Post("/event")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}
