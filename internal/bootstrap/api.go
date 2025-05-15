package bootstrap

import (
	"log"

	"github.com/dubininme/go-clean-simple-monolith/internal/application/usecase/command"
	"github.com/dubininme/go-clean-simple-monolith/internal/config"
	v1 "github.com/dubininme/go-clean-simple-monolith/internal/delivery/http/v1"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/service"
	"github.com/dubininme/go-clean-simple-monolith/internal/infrastructure/elasticsearch"
	"github.com/dubininme/go-clean-simple-monolith/internal/infrastructure/mysql"
)

func Run(cfg *config.Config) {
	db, err := mysql.NewDB(cfg.DB)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	orderRepo := mysql.NewMysqlOrderRepository(db)

	uow := mysql.NewMysqlOrderUow(db)
	markOrderPaidUsecase := command.NewMarkOrderPaidUsecase(uow)

	orderHandler := v1.NewOrderHandler(orderRepo, discountService, orderEventPublisher, markOrderPaidUsecase)
	productHandler := v1.NewProductHandler(productSearchRepo)
	authHandler := v1.NewAuthHandler(cfg.JWTSecret)

	esClient, err := elasticsearch.NewClient(cfg.ElasticAddress)
	if err != nil {
		log.Fatalf("failed to connect to elasticsearch: %v", err)
	}
	productSearchRepo := elasticsearch.NewProductSearchRepository(esClient)

	discountService := &service.DiscountService{}
	orderEventPublisher := &service.OrderEventPublisher{}

	server := v1.NewServer(orderRepo, discountService, orderEventPublisher, cfg, productSearchRepo, markOrderPaidUsecase)
	if err := server.Start(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
