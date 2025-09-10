package gowk

import (
	"net/url"
	"strconv"
)

// ListConnectionsResponse 查询连接列表响应
type GetConnzResponse struct {
	Now            string            `json:"now"`             // 当前服务器时间（ISO 8601）
	NumConnections int               `json:"num_connections"` // 当前连接数量
	Total          int               `json:"total"`           // 总连接数量
	Offset         int               `json:"offset"`          // 当前偏移量
	Limit          int               `json:"limit"`           // 当前限制数量
	Connections    []ConnzConnection `json:"connections"`     // 连接详情列表
}

// Connection 连接详情
type ConnzConnection struct {
	CID           int64       `json:"cid"`           // 连接 ID
	UID           string      `json:"uid"`           // 用户 ID
	IP            string      `json:"ip"`            // 客户端 IP 地址
	Port          int         `json:"port"`          // 客户端端口
	Start         string      `json:"start"`         // 连接开始时间
	LastActivity  string      `json:"last_activity"` // 最后活动时间
	Uptime        string      `json:"uptime"`        // 连接持续时间
	Idle          string      `json:"idle"`          // 空闲时间
	PendingBytes  int64       `json:"pending_bytes"` // 待发送字节数
	InMsgs        int64       `json:"in_msgs"`       // 接收消息数
	OutMsgs       int64       `json:"out_msgs"`      // 发送消息数
	InBytes       int64       `json:"in_bytes"`      // 接收字节数
	OutBytes      int64       `json:"out_bytes"`     // 发送字节数
	Subscriptions int         `json:"subscriptions"` // 订阅数量
	DeviceFlag    DeviceFlag  `json:"device_flag"`   // 设备标识
	DeviceLevel   DeviceLevel `json:"device_level"`  // 设备级别
	Version       string      `json:"version"`       // 客户端版本
}

// 获取连接信息
func (g *GoWk) GetConnz(offset, limit, subs *int) (*GetConnzResponse, error) {
	base := "/connz"
	u, err := url.Parse(base)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	if offset != nil {
		q.Set("offset", strconv.Itoa(*offset))
	}
	if limit != nil {
		q.Set("limit", strconv.Itoa(*limit))
	}
	if subs != nil {
		q.Set("subs", strconv.Itoa(*subs))
	}
	var result GetConnzResponse
	resp, err := g.restyClient.R().
		SetResult(&result).
		Get(u.String())
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

// ServerStatusResponse 服务器状态响应
type ServerStatusResponse struct {
	// 服务器信息
	ServerID  string `json:"server_id"`  // 服务器标识符
	Version   string `json:"version"`    // WuKongIM 版本号
	GitCommit string `json:"git_commit"` // Git 提交哈希
	GoVersion string `json:"go_version"` // Go 语言版本

	// 运行时信息
	Start  string `json:"start"`  // 服务器启动时间 (ISO 8601)
	Now    string `json:"now"`    // 当前时间 (ISO 8601)
	Uptime string `json:"uptime"` // 运行时长

	// 连接统计
	Connections      int `json:"connections"`       // 当前连接数
	TotalConnections int `json:"total_connections"` // 历史累计连接数
	SlowConsumers    int `json:"slow_consumers"`    // 慢消费者数量
	Subscriptions    int `json:"subscriptions"`     // 订阅总数

	// 消息统计
	InMsgs   int64 `json:"in_msgs"`   // 接收消息总数
	OutMsgs  int64 `json:"out_msgs"`  // 发送消息总数
	InBytes  int64 `json:"in_bytes"`  // 接收字节总数
	OutBytes int64 `json:"out_bytes"` // 发送字节总数

	// HTTP 请求统计
	HTTPReqStats HTTPReqStats `json:"http_req_stats"`

	// 系统资源
	CPU float64 `json:"cpu"` // CPU 使用率 (%)
	Mem int64   `json:"mem"` // 内存使用量 (字节)

	// 配置信息
	Config Config `json:"config"`
}

// HTTPReqStats HTTP 请求统计信息
type HTTPReqStats struct {
	URIStats []URIStat `json:"uri_stats"` // URI 统计列表
}

// URIStat 单个 URI 统计
type URIStat struct {
	URI     string `json:"uri"`      // 请求 URI
	Count   int    `json:"count"`    // 请求次数
	AvgTime string `json:"avg_time"` // 平均响应时间
}

// Config 系统配置信息
type Config struct {
	MaxConnections          int `json:"max_connections"`            // 最大连接数限制
	MaxSubscriptionsPerConn int `json:"max_subscriptions_per_conn"` // 每连接最大订阅数
	MaxPayload              int `json:"max_payload"`                // 最大消息载荷大小 (字节)
}

// 获取系统变量
func (g *GoWk) GetVarz(sort *string, connLimit *int, nodeId *int) (*ServerStatusResponse, error) {
	base := "/varz"
	u, err := url.Parse(base)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	if sort != nil {
		q.Set("sort", *sort)
	}
	if connLimit != nil {
		q.Set("conn_limit", strconv.Itoa(*connLimit))
	}
	if nodeId != nil {
		q.Set("node_id", strconv.Itoa(*nodeId))
	}
	var result ServerStatusResponse
	resp, err := g.restyClient.R().
		SetResult(&result).
		Get(u.String())
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}
