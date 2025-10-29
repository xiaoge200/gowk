package gowk

type Bool int

const (
	Intranet_External = 0 // 外网地址
	Intranet_Intent   = 1 // 内网地址
)

const (
	NoPersist_Y = 0 // 持久化
	NoPersist_N = 1 // 不持久化

	RedHot_Hide = 0 // 不显示红点
	RedHot_Show = 1 // 显示红点

	SyncOnce_Read  = 0 // 读扩散
	SyncOnce_Write = 1 // 写扩散（cmd消息）
)

const (
	Ban_N = 0 // 允许发言
	Ban_Y = 1 // 全员禁言

	Disband_N = 0 // 不解散频道
	Disband_Y = 1 // 解散频道

	SendBan_N = 0 // 不禁止发消息
	SendBan_Y = 1 // 禁止发消息

	AllowStranger_N = 0 // 不允许陌生人发消息
	AllowStranger_Y = 1 // 允许陌生人发消息
)

const (
	Reset_N = 0 // 不重置
	Reset_Y = 1 // 重置

	TempSubscriber_N = 0 // 永久订阅者
	TempSubscriber_Y = 1 // 临时订阅者
)

const (
	OnlyUnread_N = 0 // 所有
	OnlyUnread_Y = 1 // 只有未读
)

const (
	Force_N = 0 // 不强制
	Force_Y = 1 // 强制
)

const (
	IncludeSubs_N = 0 // 不包含订阅者信息
	IncludeSubs_Y = 1 // 包含订阅者信息
)

type DeviceFlag uint8 // DeviceFlag 设备类型
const (
	DeviceFlag_APP    DeviceFlag = iota // APP
	DeviceFlag_WEB               = 1    // WEB
	DeviceFlag_PC                = 2    // PC
	DeviceFlag_SYSTEM            = 99   // SYSTEM 系统
)

type DeviceLevel uint8 // 设备等级
const (
	DeviceLevel_Slave  DeviceLevel = 0 // 从设备
	DeviceLevel_Master DeviceLevel = 1 // 主设备
)

type ChannelType uint8 // 频道类型
const (
	ChannelType_Person          ChannelType = 1  // 个人频道
	ChannelType_Group           ChannelType = 2  // 群组频道
	ChannelType_CustomerService ChannelType = 3  // 客服频道
	ChannelType_Community       ChannelType = 4  // 社区频道
	ChannelType_CommunityTopic  ChannelType = 5  // 社区话题频道
	ChannelType_Info            ChannelType = 6  // 资讯频道（有临时订阅者的概念，查看资讯的时候加入临时订阅，退出资讯的时候退出临时订阅）
	ChannelType_Data            ChannelType = 7  // 数据频道
	ChannelType_Temp            ChannelType = 8  // 临时频道
	ChannelType_Live            ChannelType = 9  // 直播频道(直播频道不会保存最近会话数据)
	ChannelType_Visitors        ChannelType = 10 // 访客频道 (频道id即为访客id，此频道只支持一个访客订阅者，多个客服订阅者，ChannelTypeCustomerService频道已过时，使用ChannelTypeVisitors代替)
	ChannelType_Agent           ChannelType = 11 // 单聊Agent频道（AI Agent频道，频道ID内部结构为UID@AgentID的结构，类似单聊频道，此频道会针对于AI Agent场景做优化）
	ChannelType_AgentGroup      ChannelType = 12 // 群聊Agent频道（AI Agent群聊频道，类似群聊频道，此频道会针对于多Agent协同场景做优化）

)

type PullMode int // 拉取模式
const (
	PullMode_Down PullMode = iota // 向下拉取
	PullMode_Up                   // 向上拉取
)

const (
	OnlineStatus_Offline = 0 // 离线
	OnlineStatus_Online  = 1 // 在线
)

const (
	VarzSort_InMsgs   = "in_msgs"   // 按接收消息数排序
	VarzSort_OutMsgs  = "out_msgs"  // 按发送消息数排序
	VarzSort_InBytes  = "in_bytes"  //  按接收字节数排序
	VarzSort_OutBytes = "out_bytes" //  按发送字节数排序
)

const (
	MigrateStatus_Running   = "running"   // 迁移正在进行中
	MigrateStatus_Completed = "completed" //  迁移已完成
	MigrateStatus_Migrated  = "migrated"  //  迁移已完成（历史状态）
)

const (
	WebHookEvent_UserOnlineStatus = "user.onlinestatus" // 用户在线状态通知
	MigrateStatus_MsgOffline      = "msg.offline"       // 离线消息通知
	MigrateStatus_MsgNotify       = "msg.notify"        // 所有消息通知
)
