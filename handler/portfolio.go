package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/natchapol2347/woodman_web_back/port/input"
	"github.com/natchapol2347/woodman_web_back/port/output"
	"github.com/natchapol2347/woodman_web_back/service"
)

type PortfolioHandler struct {
	service service.IService
}

func NewPortfolioHandler(service service.IService) *PortfolioHandler {
	return &PortfolioHandler{
		service: service,
	}

}
func (h *PortfolioHandler) GetProject(ctx echo.Context) error {
	req := &input.GetProjectReq{}
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	res, err := h.service.GetProject(ctx, req)
	if err != nil {
		if customErr, ok := err.(*output.ErrorResponse); ok {
			return ctx.JSON(customErr.StatusCode, customErr)
		}
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)

}

func (h *PortfolioHandler) GetManyProjects(ctx echo.Context) error {
	// req := &input.AllProjectsReq{}
	// if err := ctx.Bind(&req); err != nil {
	// 	return err
	// }

	res, err := h.service.GetManyProjects(ctx)
	if err != nil {
		if customErr, ok := err.(*output.ErrorResponse); ok {
			return ctx.JSON(customErr.StatusCode, customErr)
		}
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)

}

func (h *PortfolioHandler) PostProject(ctx echo.Context) error {
	req := &input.PostProjectReq{}
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	res, err := h.service.PostProject(ctx, req)
	if err != nil {
		if customErr, ok := err.(*output.ErrorResponse); ok {
			return ctx.JSON(customErr.StatusCode, customErr)
		}
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)

}
