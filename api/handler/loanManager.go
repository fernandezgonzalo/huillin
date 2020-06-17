package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) GetLoan(c echo.Context) (err error) {
	id := c.Param("id")
	loan, _ := h.DAL.GetLoan(id)
	return c.JSON(http.StatusCreated, loan)
}

func (h *Handler) GetLoans(c echo.Context) (err error) {
	return nil
}