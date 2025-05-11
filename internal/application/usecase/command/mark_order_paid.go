package command

import (
	"context"
	"encoding/json"
	"time"

	appRepository "github.com/dubininme/go-clean-simple-monolith/internal/application/repository"
	"github.com/dubininme/go-clean-simple-monolith/internal/application/uow"
	domainEntity "github.com/dubininme/go-clean-simple-monolith/internal/domain/entity"
	domainEvent "github.com/dubininme/go-clean-simple-monolith/internal/domain/event"
	"github.com/google/uuid"
)

type MarkOrderPaidUsecase struct {
	uow uow.OrderUnitOfWork
}

func (uc *MarkOrderPaidUsecase) Execute(ctx context.Context, id string) error {
	return uc.uow.DoInTx(ctx, func(repos *uow.OrderRepos) error {
		order, err := repos.OrderRepository.FindById(ctx, id)
		if err != nil {
			return err
		}

		order.Status = domainEntity.OrderStatusPaid
		if err := repos.OrderRepository.Save(ctx, order); err != nil {
			return err
		}

		orderPaidEvent := domainEvent.OrderPaid{
			OrderId: order.Id,
			Amount:  order.Amount,
			PaidAt:  time.Now(),
		}

		payload, err := json.Marshal(orderPaidEvent)
		if err != nil {
			return err
		}

		event := appRepository.OutboxMessage{
			Id:        uuid.New().String(),
			EventType: appRepository.OutboxMessageTypeOrderPaid,
			Payload:   payload,
		}
		if err := repos.OutboxRepository.Add(event); err != nil {
			return err
		}
		return nil
	})
}
