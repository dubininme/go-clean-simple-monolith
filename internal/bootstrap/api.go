package bootstrap

import (
	"log"

	"github.com/dubininme/go-clean-simple-monolith/internal/application/usecase/command"
	"github.com/dubininme/go-clean-simple-monolith/internal/application/usecase/query"
	"github.com/dubininme/go-clean-simple-monolith/internal/config"
	v1 "github.com/dubininme/go-clean-simple-monolith/internal/delivery/http/v1"
	v1handler "github.com/dubininme/go-clean-simple-monolith/internal/delivery/http/v1/handler"
	"github.com/dubininme/go-clean-simple-monolith/internal/infrastructure/mysql"
)

func Run(cfg *config.Config) {
	db, err := mysql.NewDB(cfg.DB)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	orderRepo := mysql.NewMysqlOrderRepository(db)
	uow := mysql.NewMysqlOrderUow(db)

	findOrderUsecase := query.NewFindOrderUsecase(orderRepo)
	createOrderUsecase := command.NewCreateOrderUsecase(uow)
	markOrderPaidUsecase := command.NewMarkOrderPaidUsecase(uow)

	orderHandler := v1handler.NewOrderHandler(findOrderUsecase, createOrderUsecase, markOrderPaidUsecase)

	router := v1.NewRouter(orderHandler)
	server := v1.NewServer(router)
	if err := server.Start(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
