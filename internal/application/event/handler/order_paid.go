package handler

import "github.com/dubininme/go-clean-simple-monolith/internal/application/usecase/command"

type OrderPaidHandler struct {
	orderPaidUsecase command.MarkOrderPaidUsecase
}
