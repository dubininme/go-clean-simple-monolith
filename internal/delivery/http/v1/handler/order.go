package v1

import (
	"net/http"
	"strconv"

	"github.com/dubininme/go-clean-simple-monolith/internal/application/usecase/command"
	"github.com/dubininme/go-clean-simple-monolith/internal/application/usecase/query"
	"github.com/dubininme/go-clean-simple-monolith/pkg/gen/oapi"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	findOrderUsecase     query.FindOrderUsecase
	createOrderUsecase   command.CreateOrderUsecase
	markOrderPaidUsecase command.MarkOrderPaidUsecase
}

func NewOrderHandler(
	findOrderUsecase query.FindOrderUsecase,
	createOrderUsecase command.CreateOrderUsecase,
	markOrderPaidUsecase command.MarkOrderPaidUsecase,
) *OrderHandler {
	return &OrderHandler{
		findOrderUsecase:     findOrderUsecase,
		createOrderUsecase:   createOrderUsecase,
		markOrderPaidUsecase: markOrderPaidUsecase,
	}
}

func (h *OrderHandler) CreateOrder(c echo.Context) error {
	var req oapi.CreateOrderRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	orderItems := make([]command.CreateOrderItemCommand, len(req.Items))
	for i, item := range req.Items {
		orderItems[i] = command.CreateOrderItemCommand{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
	}
	cmd := command.CreateOrderCommand{
		Currency: req.Currency,
		Items:    orderItems,
	}

	orderId, err := h.createOrderUsecase.Execute(c.Request().Context(), cmd)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not create order"})
	}

	return c.JSON(http.StatusOK, oapi.CreateOrderResponse{Id: &orderId})
}

func (h *OrderHandler) GetOrder(c echo.Context) error {
	paramId := c.Param("id")

	orderId, err := strconv.Atoi(paramId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid order id"})
	}

	order, err := h.findOrderUsecase.Execute(c.Request().Context(), int32(orderId))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "order not found"})
	}

	resItems := make([]oapi.OrderItem, len(order.Items))
	for i, item := range order.Items {
		resItems[i] = oapi.OrderItem{
			Id:        item.Id,
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
	}

	res := oapi.Order{
		Id:          order.Id,
		Status:      oapi.OrderStatus(order.Status),
		Amount:      order.Amount,
		Currency:    order.Currency,
		CreatedAt:   order.CreatedAt,
		CancelledAt: order.CancelledAt,
		Items:       resItems,
	}

	return c.JSON(http.StatusOK, res)
}

func (h *OrderHandler) MarkOrderPaid(c echo.Context) error {
	paramId := c.Param("id")

	orderId, err := strconv.Atoi(paramId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid order id"})
	}

	if err := h.markOrderPaidUsecase.Execute(c.Request().Context(), int32(orderId)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not mark order as paid"})
	}
	return c.NoContent(http.StatusOK)
}
