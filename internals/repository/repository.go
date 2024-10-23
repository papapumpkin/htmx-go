package repository

import "htmx-go/internals/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id uint) error
}

type OrderRepository interface {
	CreateOrder(order *models.Order) error
	GetOrderByID(id uint) (*models.Order, error)
	GetAllOrders() ([]models.Order, error)
	UpdateOrder(order *models.Order) error
	DeleteOrder(id uint) error
}
