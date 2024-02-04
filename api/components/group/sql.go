package group

// group
const (
	SqlQueryGroupByID = `SELECT * 
						 FROM groups 
						 WHERE id = ? AND deleted_at IS NULL`
	SqlQueryGroupByName = `SELECT * 
						   FROM groups 
						   WHERE name = ? AND deleted_at IS NULL`
	SqlQueryGroupByUserID = `SELECT * 
							 FROM groups 
							 WHERE id IN (SELECT group_id 
										  FROM group_members 
										  WHERE user_id = ? AND deleted_at IS NULL) AND deleted_at IS NULL 
										  ORDER BY id DESC`
	SqlQueryGroupByPage = `SELECT * 
						   FROM groups 
						   WHERE deleted_at IS NULL 
						   ORDER BY id DESC 
						   LIMIT ? OFFSET ?`
	SqlUpdateGroupNameByID = `UPDATE groups 
							  SET name = ? 
							  WHERE id = ?`
	SqlUpdateGroupLogoByID = `UPDATE groups 
							  SET logo = ? 
							  WHERE id = ?`
	SqlUpdateGroupNoticeByID = `UPDATE groups 
								SET notice = ? 
								WHERE id = ?`
	SqlUpdateGroupAddrByID = `UPDATE groups 
							  SET addr = ? 
							  WHERE id = ?`
	SqlDeleteGroupnByID = `DELETE 
						   FROM group 
						   WHERE id = ?`
)

// group application
const (
	SqlQueryGroupApplicationsByGroupID = `SELECT * 
										  FROM group_application 
										  WHERE group_id = ? AND deleted_at IS NULL`
	SqlQueryGroupApplicationsByUserID = `SELECT * 
										FROM group_application 
										WHERE user_id = ? AND deleted_at IS NULL`
	SqlQueryGroupApplicationByGroupIDAndUserID = `SELECT * 
												  FROM group_application 
												  WHERE group_id = ? AND user_id = ? AND deleted_at IS NULL`
	SqlDeleteGroupApplicationByGroupIDAndUserID = `DELETE 
												   FROM group_application 
												   WHERE group_id = ? AND user_id = ?`
	SqlDeleteGroupApplicationByGroupID = `DELETE 
										  FROM group_application 
										  WHERE group_id = ?`
)

// group member
const (
	SqlQueryGroupMemberByGroupID = `SELECT * 
									FROM group_members 
									WHERE group_id = ? AND deleted_at IS NULL`
	SqlQueryGroupMemberByGroupIDAndPosition = `SELECT * 
											   FROM group_members 
											   WHERE group_id = ? AND position > ? AND deleted_at IS NULL`
	SqlQueryGroupMemberByGroupIDAndUserID = `SELECT * 
											 FROM group_members 
											 WHERE group_id = ? AND user_id = ? AND deleted_at IS NULL`
	SqlUpdateGroupMemberPositionByGroupIDAndUserID = `UPDATE group_members 
													  SET position = ? 
													  WHERE group_id = ? AND user_id = ? AND deleted_at IS NULL`
	SqlDeleteGroupMemberByGroupID = `DELETE 
									 FROM group_members 
									 WHERE group_id = ?`
	SqlDeleteGroupMemberByGroupIDAndUserID = `DELETE 
											  FROM group_members 
											  WHERE group_id = ? AND user_id = ?`
)
