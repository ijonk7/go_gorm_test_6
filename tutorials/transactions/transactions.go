package transactions

import (
	"errors"
	"gorm.io/gorm"
	"gorm_test_6/model"
)

/*
	Transactions - Transaction
*/
func Transaction(db *gorm.DB) {
	db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		if err := tx.Model(&model.User{}).Where("id = ?", 9).Update("name", "jinzhu 999").Error; err != nil {
			// return any error will rollback
			return err
		}

		if err := tx.Where("id = ?", 7).Delete(&model.User{}).Error; err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})
}

/*
	Transactions - Nested Transactions
*/
func NestedTransactions(db *gorm.DB) {
	db.Transaction(func(tx *gorm.DB) error {
		tx.Model(&model.User{}).Where("id = ?", 8).Update("name", "jinzhu 8")

		tx.Transaction(func(tx2 *gorm.DB) error {
			tx.Model(&model.User{}).Where("id = ?", 9).Update("name", "jinzhu 9")
			return errors.New("rollback user2") // Rollback user2
		})

		tx.Transaction(func(tx2 *gorm.DB) error {
			tx.Model(&model.User{}).Where("id = ?", 4).Update("name", "jinzhu 4")
			return nil
		})

		return nil
	})
}

/*
	Transactions - Control the transaction manually
*/
func ControlTransactionManually(db *gorm.DB) {
	// begin a transaction
	tx := db.Begin()

	// do some database operations in the transaction (use 'tx' from this point, not 'db')
	tx.Model(&model.User{}).Where("id = ?", 9).Update("name", "jinzhu 9")
	tx.Model(&model.User{}).Where("id = ?", 4).Update("name", "jinzhu 4")

	// rollback the transaction in case of error
	//tx.Rollback()

	// Or commit the transaction
	tx.Commit()
}

/*
	Transactions - Specific Example
*/
func SpecificExample(db *gorm.DB) error {
	// Note the use of tx as the database handle once you are within a transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Model(&model.User{}).Where("id = ?", 9).Update("name", "jinzhu 9").Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&model.User{}).Where("id = ?", 4).Update("name", "jinzhu 4").Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
