package command

import (
	"context"

	"github.com/dubininme/go-clean-simple-monolith/internal/application/uow"
	domainEntity "github.com/dubininme/go-clean-simple-monolith/internal/domain/entity"
	domainEvent "github.com/dubininme/go-clean-simple-monolith/internal/domain/event"
)

type MarkOrderCancelledUsecase struct {
	uow uow.OrderUnitOfWork
}

func NewMarkOrderCancelledUsecase(uow uow.OrderUnitOfWork) *MarkOrderCancelledUsecase {
	return &MarkOrderCancelledUsecase{uow: uow}
}

func (uc *MarkOrderCancelledUsecase) Execute(ctx context.Context, event domainEvent.OrderCancelled) error {
	return uc.uow.DoInTx(ctx, func(repos *uow.OrderRepos) error {
		order, err := repos.OrderRepository.FindById(ctx, event.OrderId)
		if err != nil {
			return err
		}

		order.Status = domainEntity.OrderStatusCancelled
		order.CancelledAt = &event.CancelledAt
		if _, err := repos.OrderRepository.Save(ctx, *order); err != nil {
			return err
		}

		return nil
	})
}
