package activity

const (
	SqlQueryActivityByID = `SELECT * 
							FROM activities 
							WHERE id = ? AND deleted_at IS NULL`
	SqlQueryActivitiesByOrganizerUserID = `SELECT * 
											FROM activities 
											WHERE user_id = ? AND deleted_at IS NULL 
											ORDER BY id DESC`
	SqlQueryActivitiesByUserID = `SELECT * 
								  FROM activities
								  WHERE id IN (SELECT activity_id 
											   FROM activity_users 
											   WHERE user_id = ? AND deleted_at IS NULL) AND deleted_at IS NULL 
											   ORDER BY id DESC`
	SqlQueryActivitiesByUserID2 = `SELECT activities.* 
								   FROM activities, activity_users
								   WHERE activities.user_id = activity_users.user_id AND activity_users.user_id = ? AND activities.deleted_at IS NULL AND activity_users.deleted_at IS NULL`
	SqlQueryActivitiesByGroupID = `SELECT * 
								   FROM activities
								   WHERE group_id = ? AND deleted_at IS NULL
								   ORDER BY id DESC`
	SqlQueryActivitiesByPage = `SELECT * 
								FROM activities 
								WHERE deleted_at IS NULL 
								ORDER BY id DESC 
								LIMIT ? OFFSET ?`
	SqlDeleteActivitiesByGroupID = `DELETE 
									FROM activities 
									WHERE group_id = ?`
)

const (
	SqlQueryActivityUsersByActivityID = `SELECT * 
										 FROM activity_users 
										 WHERE activity_id = ? AND deleted_at IS NULL`
	SqlQueryActivityUsersByUserID = `SELECT * 
									 FROM activity_users 
									 WHERE user_id = ? AND deleted_at IS NULL`
	SqlQueryActivityUserByActivityIDAndUserID = `SELECT * 
												 FROM activity_users 
												 WHERE activity_id = ? AND user_id = ? AND deleted_at IS NULL`
	SqlInsertActivityUser     = `INSERT INTO activity_users (activity_id, user_id, nick, is_friend) VALUES (?, ?, ?, ?)`
	SqlDeleteActivityUserByID = `DELETE 
								 FROM activity_users 
								 WHERE id = ?`
)
