package gowk

import "context"

// 健康检查
func (g *GoWk) HealthCheck(ctx context.Context) (*StatusResponse, error) {
	var result StatusResponse
	resp, err := g.restyClient.R().
		SetContext(ctx).
		SetResult(&result).
		Get("/health")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}

type MigrationStatusResponse struct {
	Status   string  `json:"status"`             // 迁移状态 (running, completed, migrated)
	Step     string  `json:"step"`               // 当前迁移步骤 (message, user, channel 等)
	LastErr  *string `json:"last_err,omitempty"` // 最后一次错误信息，可为 null
	TryCount int     `json:"try_count"`          // 尝试次数
}

// 获取迁移结果
func (g *GoWk) GetMigrateReult(ctx context.Context) (*MigrationStatusResponse, error) {
	var result MigrationStatusResponse
	resp, err := g.restyClient.R().
		SetContext(ctx).
		SetResult(&result).
		Get("/migrate/result")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp.Error().(error)
	}
	return &result, nil
}
