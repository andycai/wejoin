package group

// group
const (
	SqlQueryGroupByID     = "SELECT * FROM groups WHERE id = ? AND deleted_at IS NULL"
	SqlQueryGroupByName   = "SELECT * FROM groups WHERE name = ? AND deleted_at IS NULL"
	SqlQueryGroupByUserID = "SELECT * FROM groups WHERE user_id = ? AND deleted_at IS NULL"
	SqlQueryGroupByPage   = "SELECT * FROM groups WHERE deleted_at IS NULL LIMIT ? OFFSET ?"
	SqlUpdateGroupByID    = "UPDATE groups SET name = ? WHERE id = ?"
	SqlDeleteGroupnByID   = "DELETE FROM group WHERE id = ?"
)

// group application
const (
	SqlQueryGroupApplicationsByGroupID          = "SELECT * FROM group_application WHERE group_id = ? AND deleted_at IS NULL"
	SqlQueryGroupApplicationsByGroupIDAndUserID = "SELECT * FROM group_application WHERE group_id = ? AND user_id = ? AND deleted_at IS NULL"
	SqlDeleteGroupApplicationByGroupIDAndUserID = "DELETE FROM group_application WHERE group_id = ? AND user_id = ?"
	SqlDeleteGroupApplicationByGroupID          = "DELETE FROM group_application WHERE group_id = ?"
)

// group member
const (
	SqlQueryGroupMemberByGroupID                   = "SELECT * FROM group_member WHERE group_id = ? AND deleted_at IS NULL"
	SqlQueryGroupMemberByGroupIDAndPosition        = "SELECT * FROM group_member WHERE group_id = ? AND position > ? AND deleted_at IS NULL"
	SqlQueryGroupMemberByGroupIDAndUserID          = "SELECT * FROM group_member WHERE group_id = ? AND user_id = ? AND deleted_at IS NULL"
	SqlUpdateGroupMemberPositionByGroupIDAndUserID = "UPDATE group_member SET position = ? WHERE group_id = ? AND user_id = ? AND deleted_at IS NULL"
	SqlDeleteGroupMemberByGroupID                  = "DELETE FROM group_member WHERE group_id = ?"
	SqlDeleteGroupMemberByGroupIDAndUserID         = "DELETE FROM group_member WHERE group_id = ? AND user_id = ?"
)
