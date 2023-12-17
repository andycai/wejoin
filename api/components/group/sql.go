package group

const (
	SqlGroupByID = "SELECT * FROM groups WHERE id = ? AND deleted_at IS NULL"
)
