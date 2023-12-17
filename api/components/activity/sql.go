package activity

const (
	SqlActivityByID       = "SELECT * FROM activities WHERE id = ? AND deleted_at IS NULL"
	SqlActivitiesByUserID = "SELECT * FROM activities WHERE user_id = ? AND deleted_at IS NULL"
)
