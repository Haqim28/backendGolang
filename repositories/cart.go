package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	CreateCart(order models.Cart) (models.Cart, error)
	GetCart(ID int) (models.Cart, error)
	UpdateCart(Cart models.Cart, ID int) (models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetCart(ID int) (models.Cart, error) {
	var cart models.Cart

	err := r.db.Preload("User").Preload("Order.Product.User").Preload("Order.Cart.User").First(&cart, ID).Error

	return cart, err
}

func (r *repository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error

	return cart, err
}

func (r *repository) UpdateCart(cart models.Cart, ID int) (models.Cart, error) {
	err := r.db.Save(&cart).Error

	return cart, err
}
