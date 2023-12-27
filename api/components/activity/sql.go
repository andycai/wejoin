package activity

const (
	SqlQueryActivityByID                = "SELECT * FROM activities WHERE id = ? AND deleted_at IS NULL"
	SqlQueryActivitiesByOrganizerUserID = "SELECT * FROM activities WHERE id IN (SELECT activity_id FROM activity_user WHERE user_id = ? AND deleted_at IS NULL) AND deleted_at IS NULL ORDER BY id DESC"
	SqlQueryActivitiesByUserID          = "SELECT * FROM activities WHERE id IN (SELECT activity_id FROM activity_user WHERE user_id = ? AND deleted_at IS NULL) AND deleted_at IS NULL ORDER BY id DESC"
	SqlQueryActivitiesByGroupID         = "SELECT * FROM activities WHERE group_id = ? AND deleted_at IS NULL ORDER BY id DESC"
	SqlQueryActivitiesByPage            = "SELECT * FROM activities WHERE deleted_at IS NULL ORDER BY id DESC LIMIT ? OFFSET ?"
	SqlDeleteActivitiesByGroupID        = "DELETE FROM activities WHERE group_id = ?"
)

const (
	SqlQueryActivityUsersByActivityID         = "SELECT * FROM activity_user WHERE activity_id = ? AND deleted_at IS NULL"
	SqlQueryActivityUsersByUserID             = "SELECT * FROM activity_user WHERE user_id = ? AND deleted_at IS NULL"
	SqlQueryActivityUserByActivityIDAndUserID = "SELECT * FROM activity_user WHERE activity_id = ? AND user_id = ? AND deleted_at IS NULL"
	SqlDeleteActivityUserByID                 = "DELETE FROM activity_user WHERE id = ?"
)
