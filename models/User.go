package models

type User struct {
	Id       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

func Login(username, password string) (*User, bool) {
	var user User
	db.Select("id").Where(User{Username: username, Pwd: password}).First(&user)
	if user.Id > 0 {
		return &user, true
	}
	return &user, false
}
