package models

type User struct {
	Id        int    `gorm:"AUTO_INCREMENT" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}
