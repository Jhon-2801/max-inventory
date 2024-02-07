package roles

type UserRoles struct {
	ID     int
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
}
