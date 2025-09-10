package gowk

type SyncConversationRequest struct {
	UID                 string        `json:"uid"`                             // 必传：用户 ID
	Version             *int64        `json:"version,omitempty"`               // 可选：版本时间戳（增量同步）
	LastMsgSeqs         *string       `json:"last_msg_seqs,omitempty"`         // 可选：最后消息序列号
	MsgCount            *int          `json:"msg_count,omitempty"`             // 可选：每个会话返回的最近消息数量
	OnlyUnread          *int          `json:"only_unread,omitempty"`           // 可选：是否只返回未读 (0=全部,1=未读)
	ExcludeChannelTypes []ChannelType `json:"exclude_channel_types,omitempty"` // 可选：要排除的频道类型
}

// Conversation 会话信息
type Conversation struct {
	ChannelID   string                `json:"channel_id"`   // 频道 ID
	ChannelType ChannelType           `json:"channel_type"` // 频道类型 (1=个人,2=群组)
	Unread      int                   `json:"unread"`       // 未读消息数量
	Timestamp   int64                 `json:"timestamp"`    // 最后消息时间戳
	LastMsgSeq  int64                 `json:"last_msg_seq"` // 最后消息序列号
	Version     int64                 `json:"version"`      // 会话版本号（纳秒级时间戳）
	Messages    []ConversationMessage `json:"messages"`     // 最新消息列表
}

// Message 会话中的消息对象
type ConversationMessage struct {
	MessageID   int64  `json:"message_id"`    // 消息 ID
	MessageSeq  int64  `json:"message_seq"`   // 消息序列号
	ClientMsgNo string `json:"client_msg_no"` // 客户端消息编号
	FromUID     string `json:"from_uid"`      // 发送者用户 ID
	Timestamp   int64  `json:"timestamp"`     // 消息时间戳
	Payload     string `json:"payload"`       // Base64 编码的消息内容
}

// 同步用户会话
func (g *GoWk) SyncConversation(req SyncConversationRequest) ([]Conversation, error) {
	var result []Conversation
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/conversation/sync")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return result, nil
}

type ConversationInfo struct {
	UID         string      `json:"uid"`          // 必传，用户 ID
	ChannelID   string      `json:"channel_id"`   // 必传，频道 ID
	ChannelType ChannelType `json:"channel_type"` // 必传，频道类型 (1=个人,2=群组)
}

type ClearConversationUnreadRequest struct {
	ConversationInfo
	MessageSeq *int64 `json:"message_seq,omitempty"` // 可选，清除到哪条消息
}

// 清除未读消息
func (g *GoWk) ClearConversationUnread(req ClearConversationUnreadRequest) (*StatusResponse, error) {
	var result StatusResponse
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/conversations/clearUnread")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

type SetConversationUnreadRequest struct {
	ConversationInfo
	Unread int `json:"unread"` // 未读消息数量
}

// 设置会话未读数
func (g *GoWk) SetConversationUnread(req SetConversationUnreadRequest) (*StatusResponse, error) {
	var result StatusResponse
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/conversations/setUnread")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

// 删除会话
func (g *GoWk) DeleteConversation(req ConversationInfo) (*StatusResponse, error) {
	var result StatusResponse
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/conversations/delete")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}
