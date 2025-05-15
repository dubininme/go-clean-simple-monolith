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
	ProductId int32
	Quantity  int32
	Price     int64
}

type CreateOrderUsecase struct {
	uow uow.OrderUnitOfWork
}

func NewCreateOrderUsecase(uow uow.OrderUnitOfWork) *CreateOrderUsecase {
	return &CreateOrderUsecase{uow: uow}
}

func (uc *CreateOrderUsecase) Execute(ctx context.Context, cmd CreateOrderCommand) (int32, error) {
	var orderId int32

	err := uc.uow.DoInTx(ctx, func(repos *uow.OrderRepos) error {
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

		id, err := repos.OrderRepository.Save(ctx, order)
		if err != nil {
			return err
		}
		// Use closure to pass orderId from inside the transaction to the outer scope,
		// since DoInTx only allows returning an error, but we also need to return the created order ID.
		orderId = id

		preparedOrderItems := make([]domainEntity.OrderItem, len(cmd.Items))
		for i, item := range cmd.Items {
			preparedOrderItems[i] = domainEntity.OrderItem{
				ProductId: item.ProductId,
				Quantity:  item.Quantity,
				Price:     item.Price,
				OrderId:   id,
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			}
		}

		if err := repos.OrderItemRepository.SaveBatch(ctx, preparedOrderItems); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return orderId, nil
}
