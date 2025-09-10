package gowk

type Channel struct {
	ChannelID     string      `json:"channel_id"`               // 必传，频道 ID，必须唯一
	ChannelType   ChannelType `json:"channel_type"`             // 必传，频道类型 (1=个人频道, 2=群组频道)
	Ban           *int        `json:"ban,omitempty"`            // 可选，是否禁言 (0=允许发言,1=全员禁言)
	Disband       *int        `json:"disband,omitempty"`        // 可选，是否解散频道 (1=解散)
	SendBan       *int        `json:"send_ban,omitempty"`       // 可选，是否禁止发送消息 (0=不禁止,1=禁止)
	AllowStranger *int        `json:"allow_stranger,omitempty"` // 可选，是否允许陌生人发送消息 (仅个人频道支持)
	Subscribers   *[]string   `json:"subscribers,omitempty"`    // 可选，订阅者用户 ID 列表
}

// 创建频道
func (g *GoWk) CreateChannel(req Channel) (*StatusResponse, error) {
	var result StatusResponse
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/channel")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

// 更新频道信息
func (g *GoWk) UpdateChannelInfo(req Channel) (*StatusResponse, error) {
	var result StatusResponse
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/channel/info")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

// 删除频道
func (g *GoWk) DeleteChannel(req Channel) (*StatusResponse, error) {
	var result StatusResponse
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/channel/delete")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

type ChannelSubscriber struct {
	ChannelID      string      `json:"channel_id"`                // 必传，频道 ID
	ChannelType    ChannelType `json:"channel_type"`              // 必传，频道类型 (1=个人频道, 2=群组频道)
	Subscribers    []string    `json:"subscribers"`               // 必传，要添加的订阅者用户 ID 列表
	Reset          *int        `json:"reset,omitempty"`           // 可选，是否重置现有订阅者 (0=不重置,1=重置)
	TempSubscriber *int        `json:"temp_subscriber,omitempty"` // 可选，是否为临时订阅者 (0=永久,1=临时)
}

// 添加频道订阅者
func (g *GoWk) AddChannelSubscriber(req ChannelSubscriber) (*StatusResponse, error) {
	var result StatusResponse
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/channel/subscriber_add")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

// 移除频道订阅者
func (g *GoWk) RemoveChannelSubscriber(req ChannelSubscriber) (*StatusResponse, error) {
	var result StatusResponse
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/channel/subscriber_remove")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

type ChannelBlackOrWhiteList struct {
	ChannelID   string      `json:"channel_id"`   // 必传，频道 ID
	ChannelType ChannelType `json:"channel_type"` // 必传，频道类型 (1=个人频道, 2=群组频道)
	UIDs        []string    `json:"uids"`         // 要加入黑名单的用户 ID 列表
}

// 添加频道黑名单
func (g *GoWk) AddChannelBlackList(req ChannelBlackOrWhiteList) (*StatusResponse, error) {
	var result StatusResponse
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/channel/blacklist_add")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

// 设置频道黑名单
func (g *GoWk) SetChannelBlackList(req ChannelBlackOrWhiteList) (*StatusResponse, error) {
	var result StatusResponse
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/channel/blacklist_set")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

// 移除频道黑名单
func (g *GoWk) RemoveChannelBlackList(req ChannelBlackOrWhiteList) (*StatusResponse, error) {
	var result StatusResponse
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/channel/blacklist_remove")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

// 添加频道白名单
func (g *GoWk) AddChannelWhiteList(req ChannelBlackOrWhiteList) (*StatusResponse, error) {
	var result StatusResponse
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/channel/whitelist_add")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

// 设置频道白名单
func (g *GoWk) SetChannelWhiteList(req ChannelBlackOrWhiteList) (*StatusResponse, error) {
	var result StatusResponse
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/channel/whitelist_set")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

// 移除频道白名单
func (g *GoWk) RemoveChannelWhiteList(req ChannelBlackOrWhiteList) (*StatusResponse, error) {
	var result StatusResponse
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/channel/whitelist_remove")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

type GetChannelWhiteListRequest struct {
	ChannelID   string      `json:"channel_id"`   // 必传，频道 ID
	ChannelType ChannelType `json:"channel_type"` // 必传，频道类型 (1=个人频道, 2=群组频道)
}

// 获取频道白名单
func (g *GoWk) GetChannelWhiteList(req GetChannelWhiteListRequest) ([]string, error) {
	var result []string
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Get("/channel/whitelist")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return result, nil
}
