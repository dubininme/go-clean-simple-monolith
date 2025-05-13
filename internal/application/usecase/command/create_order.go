package command

import (
	"context"

	"github.com/dubininme/go-clean-simple-monolith/internal/application/uow"
	domainEntity "github.com/dubininme/go-clean-simple-monolith/internal/domain/entity"
)

type CreateOrderUsecase struct {
	uow uow.OrderUnitOfWork
}

func NewCreateOrderUsecase(uow uow.OrderUnitOfWork) *CreateOrderUsecase {
	return &CreateOrderUsecase{uow: uow}
}

func (uc *CreateOrderUsecase) Execute(ctx context.Context, order domainEntity.Order, orderItems []domainEntity.OrderItem) error {
	return uc.uow.DoInTx(ctx, func(repos *uow.OrderRepos) error {
		orderId, err := repos.OrderRepository.Save(ctx, order)
		if err != nil {
			return err
		}

		preparedOrderItems := make([]domainEntity.OrderItem, len(orderItems))
		for i, orderItem := range orderItems {
			preparedOrderItems[i] = orderItem
			preparedOrderItems[i].OrderId = orderId
		}

		if err := repos.OrderItemRepository.SaveBatch(ctx, preparedOrderItems); err != nil {
			return err
		}

		return nil
	})
}
