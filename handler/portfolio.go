package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/natchapol2347/woodman_web_back/port/input"
	"github.com/natchapol2347/woodman_web_back/port/output"
	"github.com/natchapol2347/woodman_web_back/service"
)

type Handler struct {
	service service.IService
}

func NewHandler(service service.IService) *Handler {
	return &Handler{
		service: service,
	}

}
func (h *Handler) GetProject(ctx echo.Context) error {
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

func (h *Handler) GetAllProjects(ctx echo.Context) error {
	// req := &input.AllProjectsReq{}
	// if err := ctx.Bind(&req); err != nil {
	// 	return err
	// }

	res, err := h.service.GetAllProjects(ctx)
	if err != nil {
		if customErr, ok := err.(*output.ErrorResponse); ok {
			return ctx.JSON(customErr.StatusCode, customErr)
		}
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)

}

func (h *Handler) PostProject(ctx echo.Context) error {
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
