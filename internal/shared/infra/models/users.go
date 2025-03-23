package shared_models

type User struct {
	ID        int    `json: "id"`
	UserName  string `json: "user_name"`
	FullName  string `json: "full_name"`
	Email     string `json: "email"`
	Address   string `json: "address"`
	Password  string `json: "password"`
	UserType  string `json: "user_type"`
	UserLevel int    `json: "user_level"`
	IsOnline  int    `json: "is_online"`
	IsLocked  int    `json: "is_locked"`
	LastSeen  string `json: "last_seen"`
	CreatedAt string `json: "created_at"`
	UpdatedAt string `json: "updated_at"`
}

func (obj *User) GetID() int {
	return obj.ID
}
func (obj *User) GetUsername() string {
	return obj.UserName
}
func (obj *User) GetFullname() string {
	return obj.FullName
}
func (obj *User) GetEmail() string {
	return obj.Email
}
func (obj *User) GetAddress() string {
	return obj.Address
}
func (obj *User) GetPassword() string {
	return obj.Password
}
func (obj *User) GetUserType() string {
	return obj.UserType
}
func (obj *User) GetUserLevel() int {
	return obj.UserLevel
}
func (obj *User) GetIsOnline() int {
	return obj.IsOnline
}
func (obj *User) GetIsLocked() int {
	return obj.IsLocked
}
func (obj *User) GetLastSeen() string {
	return obj.LastSeen
}
func (obj *User) GetCreatedAt() string {
	return obj.CreatedAt
}
func (obj *User) GetUpdatedAt() string {
	return obj.UpdatedAt
}
