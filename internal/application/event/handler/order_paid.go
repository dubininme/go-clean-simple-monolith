package handler

import (
	"context"

	"github.com/dubininme/go-clean-simple-monolith/internal/application/usecase/command"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/event"
)

type OrderPaidHandler struct {
	orderPaidUsecase command.MarkOrderPaidUsecase
}

func NewOrderPaidHandler(orderPaidUsecase command.MarkOrderPaidUsecase) *OrderPaidHandler {
	return &OrderPaidHandler{orderPaidUsecase: orderPaidUsecase}
}

func (h *OrderPaidHandler) Handle(ctx context.Context, event *event.OrderPaid) error {
	return h.orderPaidUsecase.Execute(ctx, event.OrderId)
}
