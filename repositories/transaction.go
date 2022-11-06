package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransaction(ID int) ([]models.Transaction, error)
	FindTransactionBySeller(ID int) ([]models.Transaction, error)

	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
	DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction

	err := r.db.Preload("Users").Preload("Cart.User").Preload("Cart.Order.Product.User").First(&transaction, ID).Error

	return transaction, err
}

func (r *repository) FindTransaction(ID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Where("buyer_id=?", ID).Preload("Buyer").Preload("Cart.User").Preload("Cart.Order.Product.User").Find(&transaction).Error

	return transaction, err
}
func (r *repository) FindTransactionBySeller(ID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Where("seller_id=?", ID).Preload("Buyer").Preload("Cart.User").Preload("Cart.Order.Product.User").Find(&transaction).Error

	return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error

	return transaction, err
}

func (r *repository) UpdateTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Model(&transaction).Where("id=?", ID).Updates(&transaction).Error

	return transaction, err
}

func (r *repository) DeleteTransaction(transaction models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Delete(&transaction).Error

	return transaction, err
}
