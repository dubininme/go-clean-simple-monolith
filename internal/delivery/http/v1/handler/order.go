package v1

import (
	"net/http"

	"github.com/dubininme/go-clean-simple-monolith/internal/application/usecase/command"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/repository"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/service"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderRepo            repository.OrderRepository
	discountService      *service.DiscountService
	orderEventPublisher  *service.OrderEventPublisher
	markOrderPaidUsecase *command.MarkOrderPaidUsecase
}

func NewOrderHandler(
	orderRepo repository.OrderRepository,
	discountService *service.DiscountService,
	orderEventPublisher *service.OrderEventPublisher,
	markOrderPaidUsecase *command.MarkOrderPaidUsecase,
) *OrderHandler {
	return &OrderHandler{
		orderRepo:            orderRepo,
		discountService:      discountService,
		orderEventPublisher:  orderEventPublisher,
		markOrderPaidUsecase: markOrderPaidUsecase,
	}
}

func (h *OrderHandler) GetOrder(c echo.Context) error {
	id := c.Param("id")
	order, err := h.orderRepo.FindById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "order not found"})
	}
	return c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) CreateOrder(c echo.Context) error {
	req := CreateOrderRequest{}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	order := &domain.Order{}

	CreateOrderUsecase := &command.CreateOrderUsecase{
		OrderRepo:           h.orderRepo,
		DiscountService:     h.discountService,
		OrderEventPublisher: h.orderEventPublisher,
	}

	if err := c.Bind(order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	return c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) MarkOrderPaid(c echo.Context) error {
	id := c.Param("id")
	if err := h.markOrderPaidUsecase.Execute(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not mark order as paid"})
	}
	return c.NoContent(http.StatusOK)
}
