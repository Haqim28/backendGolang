package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order models.Order) (models.Order, error)
	GetOrder(ID int) (models.Order, error)
	GetOrderByCart(ID int) ([]models.Order, error)
	UpdateOrder(Order models.Order, ID int) (models.Order, error)
	DeleteOrder(ID int, order models.Order) (models.Order, error)
}

func RepositoryOrder(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetOrder(ID int) (models.Order, error) {
	var order models.Order

	err := r.db.Preload("Product").Preload("Cart").First(&order, ID).Error

	return order, err
}

func (r *repository) CreateOrder(order models.Order) (models.Order, error) {
	err := r.db.Create(&order).Error

	return order, err
}

func (r *repository) GetOrderByCart(ID int) ([]models.Order, error) {
	var order []models.Order
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.Where("cart_id=?", ID).Preload("Product.User").Preload("Cart.User").Find(&order).Error

	return order, err
}

func (r *repository) UpdateOrder(order models.Order, ID int) (models.Order, error) {
	err := r.db.Save(&order).Error

	return order, err
}

func (r *repository) DeleteOrder(ID int, order models.Order) (models.Order, error) {
	err := r.db.Delete(&order).Error

	return order, err
}
