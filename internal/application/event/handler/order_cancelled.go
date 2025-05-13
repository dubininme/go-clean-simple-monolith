package handler

import (
	"context"

	"github.com/dubininme/go-clean-simple-monolith/internal/application/usecase/command"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/event"
)

type OrderCancelledHandler struct {
	orderCancelledUsecase command.MarkOrderCancelledUsecase
}

func NewOrderCancelledHandler(orderCancelledUsecase command.MarkOrderCancelledUsecase) *OrderCancelledHandler {
	return &OrderCancelledHandler{orderCancelledUsecase: orderCancelledUsecase}
}

func (h *OrderCancelledHandler) Handle(ctx context.Context, event event.OrderCancelled) error {
	return h.orderCancelledUsecase.Execute(ctx, event)
}
