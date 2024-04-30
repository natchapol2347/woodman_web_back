package handler

import (
	"net/http"

	"github.com/google/uuid"
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
	projectID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return err
	}
	res, err := h.service.GetProject(ctx, projectID)
	if err != nil {
		if customErr, ok := err.(*output.ErrorResponse); ok {
			return ctx.JSON(customErr.StatusCode, customErr)
		}
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)

}

func (h *PortfolioHandler) GetManyProjects(ctx echo.Context) error {

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

func (h *PortfolioHandler) DeleteProject(ctx echo.Context) error {
	projectID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return err
	}
	res, err := h.service.DeleteProject(ctx, projectID)

	if err != nil {
		if customErr, ok := err.(*output.ErrorResponse); ok {
			return ctx.JSON(customErr.StatusCode, customErr)
		}
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)

}

func (h *PortfolioHandler) UpdateProject(ctx echo.Context) error {
	projectID := ctx.Param("id")

	req := &input.UpdateProjectReq{}
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	res, err := h.service.UpdateProject(ctx, req, projectID)
	if err != nil {
		if customErr, ok := err.(*output.ErrorResponse); ok {
			return ctx.JSON(customErr.StatusCode, customErr)
		}
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)
}
