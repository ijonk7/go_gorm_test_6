package constraints

import "gorm.io/gorm"

type CreditCard13 struct {
	gorm.Model
	Number string
	UserID uint
}

type Company13 struct {
	ID     int
	Name   string
	UserID uint
}

type User13 struct {
	gorm.Model
	CompanyID  int
	Company    Company13
	CreditCard CreditCard13
}

/*
	Constraints - Foreign Key Constraint
*/
func ForeignKeyConstraint(db *gorm.DB) {
	//db.AutoMigrate(CreditCard13{})
	//db.AutoMigrate(Company13{})
	db.AutoMigrate(User13{})
}
