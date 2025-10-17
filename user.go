package gowk

type SetUserTokenRequest struct {
	UID         string       `json:"uid"`                    // 必传，用户唯一 ID
	Token       string       `json:"token"`                  // 必传，校验的 token
	DeviceFlag  DeviceFlag   `json:"device_flag"`            // 必传，设备标识 (0=app,1=web,2=desktop)
	DeviceLevel *DeviceLevel `json:"device_level,omitempty"` // 可选，设备等级 (0=从设备,1=主设备)
}

// 设置用户Token
func (g *GoWk) SetUserToken(req SetUserTokenRequest) (*StatusResponse, error) {
	var result StatusResponse
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/user/token")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

type UpdateUserTokenRequest SetUserTokenRequest

// 更新用户Token
func (g *GoWk) UpdateUserToken(req UpdateUserTokenRequest) (*StatusResponse, error) {
	return g.SetUserToken(SetUserTokenRequest(req))
}

type QuitUserDeviceRequest struct {
	UID        string     `json:"uid"`         // 必传，用户唯一 ID
	DeviceFlag DeviceFlag `json:"device_flag"` // 必传，设备标识 (0=app,1=web,2=desktop)
}

// 强制设备退出
func (g *GoWk) QuitUserDevice(req QuitUserDeviceRequest) (*StatusResponse, error) {
	var result StatusResponse
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/user/device_quit")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

type OnlineStatus struct {
	UID        string     `json:"uid"`         // 用户 ID
	Online     int        `json:"online"`      // 在线状态 (0=离线, 1=在线)
	DeviceFlag DeviceFlag `json:"device_flag"` // 设备标识 (0=app, 1=web, 2=desktop)
}

// 获取设备在线状态
func (g *GoWk) GetOnlineStatus(uid ...string) ([]OnlineStatus, error) {
	var result []OnlineStatus
	resp, err := g.restyClient.R().
		SetBody(uid).
		SetResult(&result).
		Post("/user/onlinestatus")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return result, nil
}

// 获取系统用户ID
func (g *GoWk) GetSystemUserID() ([]string, error) {
	var result []string
	resp, err := g.restyClient.R().
		SetResult(&result).
		Get("/user/systemuids")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return result, nil
}

type AddOrRemoveSystemUsersRequest struct {
	UIDs []string `json:"uids"` // 用户 ID 数组
}

// 添加系统用户ID
func (g *GoWk) AddSystemUser(req AddOrRemoveSystemUsersRequest) (*StatusResponse, error) {
	var result StatusResponse
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/user/systemuids_add")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

// 移除系统用户ID
func (g *GoWk) RemoveSystemUser(req AddOrRemoveSystemUsersRequest) (*StatusResponse, error) {
	var result StatusResponse
	resp, err := g.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post("/user/systemuids_remove")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}
