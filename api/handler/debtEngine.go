package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) GetDebt(c echo.Context) (err error) {
	id := c.Param("id")
	loan, _ := h.DAL.GetDebt(id)
	return c.JSON(http.StatusCreated, loan)
}
func (h *Handler) GetDebts(c echo.Context) (err error) {
	return nil
}