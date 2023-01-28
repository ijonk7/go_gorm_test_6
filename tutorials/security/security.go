package security

import (
	"gorm.io/gorm"
	"time"
)

type User13 struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

/*
Security - Query Condition
*/
func QueryCondition(db *gorm.DB) {
	//db.AutoMigrate(User13{})

	var user User13
	userInput := "jinzhu;drop table user13;"

	// Query Condition - safe, will be escaped
	db.Where("name = ?", userInput).First(&user)
	//Output: SELECT * FROM "user13" WHERE name = 'jinzhu;drop table user13;' ORDER BY "user13"."id" LIMIT 1

	// Query Condition - SQL injection
	//db.Where(fmt.Sprintf("name = %v", userInput)).First(&user)
	// Output: SELECT * FROM "user13" WHERE name = jinzhu;drop table user13; ORDER BY "user13"."id" LIMIT 1

	// Inline Condition - will be escaped
	//db.First(&user, "name = ?", userInput)
	// Output: SELECT * FROM "user13" WHERE name = 'jinzhu;drop table user13;' ORDER BY "user13"."id" LIMIT 1

	// Inline Condition - SQL injection
	//db.First(&user, fmt.Sprintf("name = %v", userInput))
	// Output: SELECT * FROM "user13" WHERE name = jinzhu;drop table user13; ORDER BY "user13"."id" LIMIT 1

	//userInput2 := "name; drop table user13;"

	// SQL injection Methods
	//db.Select(userInput2).First(&user)
	// Output: SELECT name; drop table user13; FROM "user13" ORDER BY "user13"."id" LIMIT 1
}
