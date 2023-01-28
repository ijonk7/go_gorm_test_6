package indexes

import (
	"gorm.io/gorm"
)

/*
	Tag Index -GORM menerima banyak pengaturan/settings index, seperti class, type, where, comment, expression, sort, collate, option
	Composite Indexes - Gunakan nama index yang sama untuk dua field akan membuat index Composite
*/
type User12 struct {
	Name         string `gorm:"index"`
	Name2        string `gorm:"index:idx_name,unique"` // Create uniqueIndex
	Name3        string `gorm:"index:,sort:desc,type:btree,length:10,where:name3 != 'jinzhu'"`
	Name4        string `gorm:"uniqueIndex"`                      // Create uniqueIndex
	Name5        string `gorm:"uniqueIndex:idx_name_5,sort:desc"` // Create uniqueIndex
	Age          int64  `gorm:"index:,comment:hello \\, world,where:age > 10"`
	Age2         int64  `gorm:"index:,expression:ABS(age)"`
	Name6        string `gorm:"index:idx_member"`                  // Create Composite Indexes
	Number6      string `gorm:"index:idx_member"`                  // Create Composite Indexes
	OID          int64  `gorm:"index:idx_id;index:idx_oid,unique"` // Create Multiple indexes
	MemberNumber string `gorm:"index:idx_id"`                      // Create Multiple indexes
}

/*
	Indexes - Index Tag / uniqueIndex / Composite Indexes / Multiple indexes
*/
func IndexTag(db *gorm.DB) {
	db.AutoMigrate(User12{})
}
