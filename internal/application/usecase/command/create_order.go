package command

import (
	"context"
	"time"

	"github.com/dubininme/go-clean-simple-monolith/internal/application/uow"
	domainEntity "github.com/dubininme/go-clean-simple-monolith/internal/domain/entity"
)

type CreateOrderCommand struct {
	Currency string
	Items    []CreateOrderItemCommand
}
type CreateOrderItemCommand struct {
	ProductID int32
	Quantity  int32
	Price     int64
}

type CreateOrderUsecase struct {
	uow uow.OrderUnitOfWork
}

func NewCreateOrderUsecase(uow uow.OrderUnitOfWork) *CreateOrderUsecase {
	return &CreateOrderUsecase{uow: uow}
}

func (uc *CreateOrderUsecase) Execute(ctx context.Context, cmd CreateOrderCommand) error {
	return uc.uow.DoInTx(ctx, func(repos *uow.OrderRepos) error {
		var amount int64
		for _, item := range cmd.Items {
			amount += int64(item.Price) * int64(item.Quantity)
		}

		order := domainEntity.Order{
			Status:    domainEntity.OrderStatusPending,
			Currency:  cmd.Currency,
			Amount:    amount,
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		}

		orderId, err := repos.OrderRepository.Save(ctx, order)
		if err != nil {
			return err
		}

		preparedOrderItems := make([]domainEntity.OrderItem, len(cmd.Items))
		for i, item := range cmd.Items {
			preparedOrderItems[i] = domainEntity.OrderItem{
				ProductId: item.ProductID,
				Quantity:  item.Quantity,
				Price:     item.Price,
				OrderId:   orderId,
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			}
		}

		if err := repos.OrderItemRepository.SaveBatch(ctx, preparedOrderItems); err != nil {
			return err
		}

		return nil
	})
}
