package schemas

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nome     string    `json:"nome"`
	Email    string    `json:"email" gorm:"unique"`
	PassWord string    `json:"password"` // <--- aqui estava o problema
	Role     string    `json:"role" gorm:"default:user"`
	Openings []Opening `json:"openings,omitempty" gorm:"foreignKey:UserId"`
}

type Opening struct {
	gorm.Model
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   bool   `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
	UserId   uint   `json:"-"`
	User     User   `json:"user,omitempty" gorm:"foreignKey:UserId"`
}
