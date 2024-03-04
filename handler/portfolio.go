package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/natchapol2347/woodman_web_back/port/input"
	"github.com/natchapol2347/woodman_web_back/service"
)

type Handler struct {
	service service.IService
}

func NewHandler() *Handler {
	return &Handler{}

}
func (h *Handler) GetPortfolio(ctx echo.Context) error {
	req := &input.PortfolioReq{}
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	res, err := h.service.GetPortfolio(ctx, req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, res)

}
