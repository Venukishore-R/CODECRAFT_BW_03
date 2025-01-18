package models

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Role     string `json:"role"`
}
