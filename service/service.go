package service

import (
	"github.com/labstack/echo"
	"github.com/natchapol2347/woodman_web_back/adaptor/storage"
	"github.com/natchapol2347/woodman_web_back/port/input"
	"github.com/natchapol2347/woodman_web_back/port/output"
)

type Service struct {
	storage *storage.Storage
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		storage: storage,
	}
}

type IService interface {
	GetPortfolio(ctx echo.Context, req *input.PortfolioReq) (*output.PortfolioRes, error)
}

func GetPortfolio(ctx echo.Context, req *input.PortfolioReq) (*output.PortfolioRes, error) {
	return &output.PortfolioRes{}, nil

}
