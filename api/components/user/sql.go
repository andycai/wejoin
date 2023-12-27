package user

const (
	SqlQueryUserByID = `SELECT * 
						FROM users 
						WHERE id = ? AND deleted_at IS NULL`
)
