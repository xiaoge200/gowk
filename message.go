package gowk

import (
	"context"
	"strconv"
)

type Header struct {
	NoPersist *int `json:"no_persist,omitempty"` // 是否持久化
	RedDot    *int `json:"red_dot,omitempty"`    // 小红点显示
	SyncOnce  *int `json:"sync_once,omitempty"`  // 是否仅同步一次
}

type Message struct {
	Payload     string      `json:"payload"`                 // 必传，Base64 编码的消息内容
	FromUID     string      `json:"from_uid"`                // 必传，发送者用户 ID
	ChannelID   string      `json:"channel_id"`              // 必传，目标频道 ID
	ChannelType ChannelType `json:"channel_type"`            // 必传，频道类型 (1=个人, 2=群组)
	Header      *Header     `json:"header,omitempty"`        // 可选，消息头部信息
	ClientMsgNo *string     `json:"client_msg_no,omitempty"` // 可选，客户端消息编号
	StreamNo    *string     `json:"stream_no,omitempty"`     // 可选，流消息编号
	Expire      *int64      `json:"expire,omitempty"`        // 可选，消息过期时间（秒），0 表示不过期
	Subscribers []string    `json:"subscribers,omitempty"`   // 可选，指定订阅者（CMD消息有效）
}

type SendMessageResponse struct {
	MessageID   int64  `json:"message_id"`    // 服务器生成的消息 ID
	MessageSeq  int64  `json:"message_seq"`   // 消息序列号
	ClientMsgNo string `json:"client_msg_no"` // 客户端消息编号（回显）
}

// 发送消息
func (g *GoWk) SendMessage(ctx context.Context, req Message) (*SendMessageResponse, error) {
	var result SendMessageResponse
	resp, err := g.restyClient.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&result).
		Post("/message/send")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

type BatchSendMessageRequst struct {
	Header      Header   `json:"header"`      // 消息头部信息
	FromUID     string   `json:"from_uid"`    // 发送者UID
	Subscribers []string `json:"subscribers"` // 订阅者 如果此字段有值，表示消息只发给指定的订阅者
	Payload     string   `json:"payload"`     // 消息内容
}

type BatchSendMessageResponse struct {
	FailUids []string `json:"fail_uids"` // 返回发送失败的用户列表
	Reason   []string `json:"reason"`    // 发送失败用户列表对应的失败原因列表，与fail_uids一一对应
}

// 批量发送消息
func (g *GoWk) BatchSendMessage(ctx context.Context, req BatchSendMessageRequst) (*BatchSendMessageResponse, error) {
	var result BatchSendMessageResponse
	resp, err := g.restyClient.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&result).
		Post("/message/sendbatch")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

type MessageSyncRequest struct {
	LoginUID        string      `json:"login_uid"`                   // 必传，当前登录用户 ID
	ChannelID       string      `json:"channel_id"`                  // 必传，频道 ID
	ChannelType     ChannelType `json:"channel_type"`                // 必传，频道类型 (1=个人, 2=群组)
	StartMessageSeq *int64      `json:"start_message_seq,omitempty"` // 可选，起始消息序号（包含）
	EndMessageSeq   *int64      `json:"end_message_seq,omitempty"`   // 可选，结束消息序号（不包含）
	Limit           *int        `json:"limit,omitempty"`             // 可选，返回消息数量限制，最大 10000
	PullMode        *PullMode   `json:"pull_mode,omitempty"`         // 可选，拉取模式 (0=向下拉取, 1=向上拉取)
}

type RecentItem struct {
	Header      Header      `json:"header"`        // 消息头
	Setting     uint8       `json:"setting"`       // // 消息设置 消息设置是一个 uint8的数字类型 为1个字节，完全由第三方自定义 比如定义第8位为已读未读回执标记，开启则为0000 0001 = 1
	MessageID   int64       `json:"message_id"`    // 消息 ID
	MessageSeq  int64       `json:"message_seq"`   // 消息序列号
	ClientMsgNo string      `json:"client_msg_no"` // 客户端消息编号（回显）
	FromUID     string      `json:"from_uid"`      // 发送者用户 ID
	ChannelID   string      `json:"channel_id"`    // 频道 ID
	ChannelType ChannelType `json:"channel_type"`  // 频道类型 (1=个人,2=群组)
	Timestamp   int64       `json:"timestamp"`     // 消息时间戳
	Payload     string      `json:"payload"`       // Base64 编码的消息内容
}

type MessageSyncResponse struct {
	StartMessageSeq uint32       `json:"start_message_seq"` // 查询的start_message_seq
	EndMessageSeq   uint32       `json:"end_message_seq"`   // 查询的end_message_seq
	More            int          `json:"more"`              // 是否有更多  0.无 1.有
	Messages        []RecentItem `json:"messages"`          // 最近N条消息
}

// 同步频道历史消息
func (g *GoWk) MessageSync(ctx context.Context, req MessageSyncRequest) (*MessageSyncResponse, error) {
	var result MessageSyncResponse
	resp, err := g.restyClient.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&result).
		Post("/channel/messagesync")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

type GetMaxMessageSeqResponse struct {
	MessageID   int64       `json:"message_id"`    // 消息 ID
	MessageSeq  int64       `json:"message_seq"`   // 消息序列号
	ClientMsgNo string      `json:"client_msg_no"` // 客户端消息编号（回显）
	FromUID     string      `json:"from_uid"`      // 发送者用户 ID
	ChannelID   string      `json:"channel_id"`    // 频道 ID
	ChannelType ChannelType `json:"channel_type"`  // 频道类型 (1=个人,2=群组)
	Timestamp   int64       `json:"timestamp"`     // 消息时间戳
	Payload     string      `json:"payload"`       // Base64 编码的消息内容
}

// 获取频道最大消息序号
func (g *GoWk) GetMaxMessageSeq(ctx context.Context, channelId string, channelType ChannelType) (*GetMaxMessageSeqResponse, error) {
	var result GetMaxMessageSeqResponse
	resp, err := g.restyClient.R().
		SetContext(ctx).
		SetResult(&result).
		SetQueryParam("channel_id", channelId).
		SetQueryParam("channel_type", strconv.Itoa(int(channelType))).
		Get("/channel/max_message_seq")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

type GetMessageResponse struct {
	MessageID   int64       `json:"message_id"`    // 服务器生成的消息 ID
	MessageSeq  int64       `json:"message_seq"`   // 消息序列号
	ClientMsgNo string      `json:"client_msg_no"` // 客户端消息编号（回显）
	FromUID     string      `json:"from_uid"`      // 发送者用户 ID
	ChannelID   string      `json:"channel_id"`    // 频道 ID
	ChannelType ChannelType `json:"channel_type"`  // 频道类型 (1=个人,2=群组)
	Timestamp   int64       `json:"timestamp"`     // 消息时间戳（Unix 时间戳）
	Payload     string      `json:"payload"`       // Base64 编码的消息内容
}

type GetMessageRequest struct {
	MessageId int64 `json:"message_id"` // 必传，消息 ID 列表
}

// 单条消息搜索
func (g *GoWk) GetMessageByID(ctx context.Context, req GetMessageRequest) (*GetMessageResponse, error) {
	var result GetMessageResponse
	resp, err := g.restyClient.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&result).
		Post("/message")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

type BatchGetMessageRequest struct {
	MessageIDs []int64 `json:"message_ids"` // 必传，消息 ID 列表
}

// 批量消息搜索
func (g *GoWk) BatchGetMessage(ctx context.Context, req BatchGetMessageRequest) ([]GetMessageResponse, error) {
	var result []GetMessageResponse
	resp, err := g.restyClient.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&result).
		Post("/messages")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return result, nil
}

type SearchUserMessagesRequest struct {
	UID          string         `json:"uid"`                     // 必传，当前用户 UID
	Payload      map[string]any `json:"payload,omitempty"`       // 可选，消息 payload，自定义字段
	PayloadTypes []int          `json:"payload_types,omitempty"` // 可选，消息类型搜索
	FromUID      *string        `json:"from_uid,omitempty"`      // 可选，发送者 UID
	ChannelID    *string        `json:"channel_id,omitempty"`    // 可选，频道 ID
	ChannelType  *ChannelType   `json:"channel_type,omitempty"`  // 可选，频道类型 (1=个人,2=群组)
	Topic        *string        `json:"topic,omitempty"`         // 可选，topic 搜索
	Limit        *int           `json:"limit,omitempty"`         // 可选，限制数量（默认 10）
	Page         *int           `json:"page,omitempty"`          // 可选，页码（默认 1）
	StartTime    *int64         `json:"start_time,omitempty"`    // 可选，消息时间（开始）
	EndTime      *int64         `json:"end_time,omitempty"`      // 可选，消息时间（结束，包含 end_time）
	Highlights   []string       `json:"highlights,omitempty"`    // 可选，需要高亮的字段
}

type SearchUserMessagesResponse struct {
	Total    int               `json:"total"`    // 消息总数量
	Limit    int               `json:"limit"`    // 查询数量限制
	Page     int               `json:"page"`     // 当前页码
	Messages []SearchedMessage `json:"messages"` // 消息列表
}

type SearchedMessage struct {
	MessageID    int64                  `json:"message_id"`      // 消息唯一 ID
	MessageIDStr string                 `json:"message_idstr"`   // 消息唯一 ID（字符串形式）
	MessageSeq   int64                  `json:"message_seq"`     // 消息序号
	ClientMsgNo  string                 `json:"client_msg_no"`   // 客户端消息唯一编号
	FromUID      string                 `json:"from_uid"`        // 发送者 UID
	ChannelID    string                 `json:"channel_id"`      // 频道 ID
	ChannelType  ChannelType            `json:"channel_type"`    // 频道类型 (1=个人频道, 2=群组频道)
	Payload      map[string]interface{} `json:"payload"`         // 消息内容对象（通用 payload，用 map 接收）
	Topic        *string                `json:"topic,omitempty"` // 可选，消息 topic
	Timestamp    int64                  `json:"timestamp"`       // 消息时间戳（10位到秒）
}

// 用户消息搜索
func (g *GoWk) SearchUserMessage(ctx context.Context, req SearchUserMessagesRequest) (*SearchUserMessagesResponse, error) {
	var result SearchUserMessagesResponse
	resp, err := g.restyClient.R().
		SetContext(ctx).
		SetBody(req).
		SetResult(&result).
		Post("/plugins/wk.plugin.search/usersearch")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}
