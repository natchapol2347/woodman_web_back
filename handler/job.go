package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/natchapol2347/woodman_web_back/port/output"
	"github.com/natchapol2347/woodman_web_back/service"
)

type JobHandler struct {
	service service.IService
}

func NewJobHandler(service service.IService) *JobHandler {
	return &JobHandler{
		service: service,
	}

}

func (h *JobHandler) GetJob(ctx echo.Context) error {

	res, err := h.service.GetProject(ctx)
	if err != nil {
		if customErr, ok := err.(*output.ErrorResponse); ok {
			return ctx.JSON(customErr.StatusCode, customErr)
		}
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)

}

func (h *JobHandler) GetManyJobs(ctx echo.Context) error {
	res, err := h.service.GetManyJobs(ctx)
	if err != nil {
		if customErr, ok := err.(*output.ErrorResponse); ok {
			return ctx.JSON(customErr.StatusCode, customErr)
		}
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, res)

}
