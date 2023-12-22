package enum

const (
	PositionGroupMember       = iota // normal member
	PositionGroupOwner               // owner
	PositionGroupManager             // manager
	PositionGroupPlaceholder1        // holder
	PositionGroupPlaceholder2        // holder
	PositionGroupPlaceholder3        // holder
)

// DefaultGroupMemmberCount 默认成员数量
const DefaultGroupMemmberCount = 40

// DefaultGroupCount 默认返回给客户端群组的一页数量
const DefaultGroupCount = 20

// DefaultGroupManagerCount 默认管理员的数量
const DefaultGroupManagerCount = 2

// DefaultActivityCount 默认返回给客户端活动的一页数量
const DefaultActivityCount = 20
