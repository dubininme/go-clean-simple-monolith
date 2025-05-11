package v1

import (
	"net/http"

	"github.com/dubininme/go-clean-simple-monolith/internal/domain/repository"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	searchRepo repository.ProductSearchRepository
}

func NewProductHandler(searchRepo repository.ProductSearchRepository) *ProductHandler {
	return &ProductHandler{searchRepo: searchRepo}
}

func (h *ProductHandler) Search(c echo.Context) error {
	q := c.QueryParam("q")
	results, err := h.searchRepo.Search(c.Request().Context(), q)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "search failed"})
	}
	return c.JSON(http.StatusOK, results)
}
