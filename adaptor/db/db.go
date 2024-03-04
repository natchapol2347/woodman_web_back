package db

import (
	"github.com/labstack/echo"
	"github.com/natchapol2347/woodman_web_back/port/output"
)

type IDB interface {
	GetPortfolio(ctx echo.Context, projectID int) *output.PortfolioRes
}

func GetPortfolio(ctx echo.Context, projectID int) *output.PortfolioRes {
	return &output.PortfolioRes{}

}
