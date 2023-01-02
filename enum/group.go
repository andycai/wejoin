package enum

const (
	PositionGroupMember       = iota // 普通成员
	PositionGroupPlaceholder1        // 占位
	PositionGroupPlaceholder2        // 占位
	PositionGroupPlaceholder3        // 占位
	PositionGroupManager             // 群管理员
	PositionGroupOwner               // 群主
)

// DefaultGroupMemmberCount 默认成员数量
const DefaultGroupMemmberCount = 40

// DefaultGroupCount 默认返回给客户端群组的一页数量
const DefaultGroupCount = 20

// DefaultGroupManagerCount 默认管理员的数量
const DefaultGroupManagerCount = 2

// DefaultActivityCount 默认返回给客户端活动的一页数量
const DefaultActivityCount = 20
