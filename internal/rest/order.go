package rest

import (
	"fmt"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"net/http"
	"order-service/internal/middleware/httplogger"
	"order-service/internal/model"
	"order-service/internal/service"
)

//CreateOrderEntry - To be invoked by http handler
func CreateOrderEntry(rw http.ResponseWriter, r *http.Request) {
	var (
		restHandler   = new(DefaultHandler)
		restValidator = new(DefaultValidator)
	)
	requestContext := r.Context()
	logger := httplogger.GetLogger(requestContext)
	orderService := service.NewOrderService(logger)
	orderController := &OrderController{
		logger:       logger,
		orderService: orderService,
	}
	Order.createOrderInternal(orderController, rw, r, restHandler, restValidator)
}

//Order - Interface used by REST to do CRUD operations on Order
type Order interface {
	createOrderInternal(rw http.ResponseWriter, r *http.Request, handler Handler, validator Validator)
}

//OrderController - Implementation of the Order Interface
type OrderController struct {
	logger       *logrus.Entry
	orderService service.OrderService
}

//createOrderInternal creates an order
// If there is no error it will create an order by its bankId, amount and information.
func (oc *OrderController) createOrderInternal(rw http.ResponseWriter, r *http.Request, restHandler Handler, validator Validator) {
	requestContext := r.Context()
	logger := httplogger.GetLogger(requestContext)
	var order model.Order
	err := restHandler.ParseRequest(r.Body, &order, logger)
	if err != nil {
		_ = render.Render(rw, r, ErrInvalidRequest(err))
		return
	}
	err = validator.ValidateStruct(order, logger)
	if err != nil {
		_ = render.Render(rw, r, ErrInvalidRequest(err))
		return
	}
	orderID, err := oc.orderService.CreateOrder(order)
	if err != nil {
		_ = render.Render(rw, r, ErrInternalServer(err))
		return
	}
	message := fmt.Sprintf("Order has been created successfully. OrderID : %v", orderID)
	response := Response{
		Code:    http.StatusCreated,
		Message: message,
	}
	restHandler.RespondCreated(rw, response, logger)
}
