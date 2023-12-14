package user

const (
	SqlUserByID = "SELECT * FROM users WHERE id = ? AND deleted_at IS NULL"
)
