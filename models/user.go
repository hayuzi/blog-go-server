package models

type User struct {
	Id       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
	UserType int    `json:"userType"`
	Email    string `json:"email"`
}

const (
	UserTypeAdmin  = 1
	UserTypeNormal = 2
)

func CheckAuth(username, password string) (*User, bool) {
	var user User
	db.Select("id").Where(User{Username: username, Pwd: password}).First(&user)
	if user.Id > 0 {
		return &user, true
	}
	return &user, false
}

func GetUsers(offset int, pageSize int, maps interface{}) (users []User) {
	db.Where(maps).Order("id DESC").Offset(offset).Limit(pageSize).Find(&users)
	return
}

func GetUserTotal(maps interface{}) (count int) {
	db.Model(&User{}).Where(maps).Count(&count)
	return
}
