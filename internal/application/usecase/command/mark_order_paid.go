package command

import (
	"context"
	"encoding/json"
	"time"

	appRepository "github.com/dubininme/go-clean-simple-monolith/internal/application/repository"
	"github.com/dubininme/go-clean-simple-monolith/internal/application/uow"
	domainEntity "github.com/dubininme/go-clean-simple-monolith/internal/domain/entity"
	domainEvent "github.com/dubininme/go-clean-simple-monolith/internal/domain/event"
)

type MarkOrderPaidUsecase struct {
	uow uow.OrderUnitOfWork
}

func NewMarkOrderPaidUsecase(uow uow.OrderUnitOfWork) *MarkOrderPaidUsecase {
	return &MarkOrderPaidUsecase{uow: uow}
}

func (uc *MarkOrderPaidUsecase) Execute(ctx context.Context, id int32) error {
	return uc.uow.DoInTx(ctx, func(repos *uow.OrderRepos) error {
		order, err := repos.OrderRepository.FindById(ctx, id)
		if err != nil {
			return err
		}

		order.Status = domainEntity.OrderStatusPaid
		if _, err := repos.OrderRepository.Save(ctx, *order); err != nil {
			return err
		}

		orderPaidEvent := domainEvent.OrderPaid{
			OrderId: order.Id,
			Amount:  order.Amount,
			PaidAt:  time.Now().Unix(),
		}

		payload, err := json.Marshal(orderPaidEvent)
		if err != nil {
			return err
		}

		event := appRepository.OutboxMessage{
			EventType: appRepository.OutboxMessageTypeOrderPaid,
			Payload:   payload,
		}
		if err := repos.OutboxRepository.Add(ctx, event); err != nil {
			return err
		}
		return nil
	})
}
