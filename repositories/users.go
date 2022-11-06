package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	GetRole(Role string) ([]models.User, error)
	GetRoleId(ID int, Role string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(user models.User) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	//err := r.db.Find(&users).Error
	err := r.db.Preload("Product").Find(&users).Error // add this code
	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Product").First(&user, ID).Error
	//err := r.db.Preload("User").Preload("Products").First(&user, ID).Error // add this code

	return user, err
}

func (r *repository) GetRole(Role string) ([]models.User, error) {
	var users []models.User

	err := r.db.Where("role = ?", Role).Preload("Product").Find(&users).Error

	return users, err
}

func (r *repository) GetRoleId(ID int, Role string) (models.User, error) {
	var user models.User
	err := r.db.Where("role = ?  && id = ?", Role, ID).First(&user).Error

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.User) (models.User, error) {
	err := r.db.Delete(&user).Error

	return user, err
}
