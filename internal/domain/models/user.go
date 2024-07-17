package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Roles    []Role `gorm:"many2many:user_roles;"`
}
