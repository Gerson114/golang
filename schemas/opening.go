package schemas

import "gorm.io/gorm"

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string // <- Corrigido aqui
	Remote   bool
	Link     string
	Salary   int64
	UserId   uint   `json:"-"`
	User     Create `gorm:"foreignKey:UserId"`
}

type Create struct {
	gorm.Model
	Nome     string
	Email    string `gorm:"unique"`
	PassWord string
	Openings []Opening `gorm:"foreignKey:UserId"` // Um usuário tem muitos openings
	Role     string    `gorm:"default:user"`      // user ou admin
}
