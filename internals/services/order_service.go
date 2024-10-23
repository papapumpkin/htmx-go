package services

import (
	"htmx-go/internals/models"
	"htmx-go/internals/repository"
)

type OrderService struct {
	OrderRepo repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) *OrderService {
	return &OrderService{OrderRepo: orderRepo}
}

func (s *OrderService) GetAllOrders() ([]models.Order, error) {
	return s.OrderRepo.GetAllOrders()
}

func (s *OrderService) GetOrderByID(id uint) (*models.Order, error) {
	return s.OrderRepo.GetOrderByID(id)
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	return s.OrderRepo.CreateOrder(order)
}

func (s *OrderService) UpdateOrder(order *models.Order) error {
	existingOrder, err := s.GetOrderByID(order.ID)
	if err != nil {
		return err
	}

	existingOrder.PatientFirstName = order.PatientFirstName
	existingOrder.PatientLastName = order.PatientLastName
	existingOrder.ProviderFirstName = order.ProviderFirstName
	existingOrder.ProviderFirstName = order.ProviderLastName

	return s.OrderRepo.UpdateOrder(existingOrder)
}

func (s *OrderService) DeleteOrder(id uint) error {
	_, err := s.GetOrderByID(id)
	if err != nil {
		return err
	}

	return s.OrderRepo.DeleteOrder(id)
}
