package activity

const (
	SqlQueryActivityByID         = "SELECT * FROM activities WHERE id = ? AND deleted_at IS NULL"
	SqlQueryActivitiesByUserID   = "SELECT * FROM activities WHERE id IN (SELECT activity_id FROM activity_user WHERE user_id = ? AND deleted_at IS NULL) AND deleted_at IS NULL"
	SqlDeleteActivitiesByGroupID = "DELETE FROM activities WHERE group_id = ?"
)
