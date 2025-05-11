package v1

import (
	"github.com/dubininme/go-clean-simple-monolith/internal/application/usecase/command"
	"github.com/dubininme/go-clean-simple-monolith/internal/config"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/repository"
	"github.com/dubininme/go-clean-simple-monolith/internal/domain/service"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echoEngine *echo.Echo
}

func NewServer(
	orderRepo repository.OrderRepository,
	discountService *service.DiscountService,
	orderEventPublisher *service.OrderEventPublisher,
	cfg *config.Config,
	productSearchRepo repository.ProductSearchRepository,
	markOrderPaidUsecase *command.MarkOrderPaidUsecase,
) *Server {
	e := NewRouter(orderRepo, discountService, orderEventPublisher, cfg, productSearchRepo, markOrderPaidUsecase)
	return &Server{echoEngine: e}
}

func (s *Server) Start() error {
	return s.echoEngine.Start(":8080")
}
