package match

const (
	SqlMatchByID = "SELECT * FROM matches WHERE id = ? AND deleted_at IS NULL"
)
