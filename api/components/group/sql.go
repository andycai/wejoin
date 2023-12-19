package group

const (
	SqlGroupByID                           = "SELECT * FROM groups WHERE id = ? AND deleted_at IS NULL"
	SqlGroupByName                         = "SELECT * FROM groups WHERE name = ? AND deleted_at IS NULL"
	SqlGroupByUserID                       = "SELECT * FROM groups WHERE user_id = ? AND deleted_at IS NULL"
	SqlGroupByPage                         = "SELECT * FROM groups WHERE deleted_at IS NULL LIMIT ? OFFSET ?"
	SqlGroupApplicationsByGroupID          = "SELECT * FROM group_application WHERE group_id = ? AND deleted_at IS NULL"
	SqlGroupApplicationsByGroupIDAndUserID = "SELECT * FROM group_application WHERE group_id = ? AND user_id = ? AND deleted_at IS NULL"
)
