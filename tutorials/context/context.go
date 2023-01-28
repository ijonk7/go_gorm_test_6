package context

import (
	"context"
	"gorm.io/gorm"
	"gorm_test_6/model"
)

/*
	Continuous session mode
*/
func ContinuousSessionMode(db *gorm.DB) {
	var user model.User

	ctx := context.Background()
	tx := db.WithContext(ctx)
	tx.First(&user, 24)
	tx.Model(&user).Update("age", 40)
}
