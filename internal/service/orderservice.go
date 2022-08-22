//Package service is for all services used throughout the project. Business Logic will go here
package service

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"order-service/internal/model"
	"time"
)

//OrderService is for all services related to order
type OrderService interface {
	CreateOrder(model.Order) (string, error)
}

//orderServiceImpl Default Implementation for Order Service
type orderServiceImpl struct {
	logger *logrus.Entry
}

//NewOrderService Constructor to return an instance of default implementation of order service
func NewOrderService(logger *logrus.Entry) OrderService {
	return &orderServiceImpl{
		logger: logger,
	}
}

//CreateOrder this method is for creation of order
func (os orderServiceImpl) CreateOrder(order model.Order) (string, error) {
	order.ID = uuid.New().String()
	order.CreatedOn = time.Now()
	os.logger.Infof(`Order Created: %v`, order.String())
	return order.ID, nil
}
