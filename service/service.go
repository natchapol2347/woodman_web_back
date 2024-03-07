package service

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/natchapol2347/woodman_web_back/adaptor/storage"
	"github.com/natchapol2347/woodman_web_back/port/input"
	"github.com/natchapol2347/woodman_web_back/port/output"
)

type Service struct {
	storage storage.IStorage
}

func NewService(storage storage.IStorage) *Service {
	return &Service{
		storage: storage,
	}
}

type IService interface {
	GetPortfolio(ctx echo.Context, req *input.PortfolioReq) (*output.PortfolioRes, error)
}

func (s *Service) GetPortfolio(ctx echo.Context, req *input.PortfolioReq) (*output.PortfolioRes, error) {
	res, err := s.storage.GetPortfolio(ctx, 0)
	if err != nil {
		fmt.Println("heyyy")
		return nil, err
	}

	return res, nil

}
