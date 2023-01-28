package delete

import (
	"gorm.io/gorm"
	"gorm_test_6/model"
)

/*
	Delete a Record
	Saat menghapus record, value yang dihapus harus memiliki primary key atau akan memicu Batch Delete
*/
func DeleteRecord(db *gorm.DB) {
	db.Delete(model.User{ID: 5})
}

/*
	Delete with primary key
	GORM memungkinkan untuk menghapus objek menggunakan primary key dengan kondisi inline, ia berfungsi dengan numbers
*/
func DeleteWithPrimaryKey(db *gorm.DB) {
	//db.Delete(model.User{ID: 5})
	db.Delete(&model.User{}, 2)
}

/*
	Soft Delete
	Jika model Anda menyertakan field gorm.DeletedAt (yang termasuk dalam gorm.Model), itu akan mendapatkan kemampuan soft delete secara otomatis!
	Saat memanggil Delete, catatan TIDAK AKAN dihapus dari database, tetapi GORM akan menetapkan nilai DeletedAt ke waktu saat ini,
	dan data tidak dapat ditemukan lagi dengan metode Query normal.
*/
func SoftDelete(db *gorm.DB) {
	db.Delete(&model.Product{ID: 1})

	//var products []model.Product
	//db.Find(&products)
	//fmt.Println(products)

	// Find soft deleted records
	//var products []model.Product
	//db.Unscoped().Find(&products)
	//fmt.Println(products)

	// Delete permanently
	//db.Unscoped().Delete(&model.Product{ID: 1})
}
