package repositories

// Iimport "gorm.io/gorm"
import "gorm.io/gorm"

// Declare repository struct here ...

type repository struct {
	db *gorm.DB
}
