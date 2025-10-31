package schemas

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nome     string    `json:"nome"`
	Email    string    `json:"email" gorm:"unique"`
	PassWord string    `json:"passWord"` // tag alinhada com JSON
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
	UserId   uint   `json:"-"` // será preenchido automaticamente
	User     User   `json:"user,omitempty" gorm:"foreignKey:UserId"`
}

type Cliente struct {
	gorm.Model
	Nome     string `json:"nome"`
	Email    string `json:"email" gorm:"unique"`
	PassWord string `json:"passWord"` // tag alinhada com JSON
	Role     string `json:"role" gorm:"default:user"`
}
