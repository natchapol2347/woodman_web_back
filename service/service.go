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
	GetProject(ctx echo.Context, req *input.ProjectReq) (*output.ProjectRes, error)
}

func (s *Service) GetProject(ctx echo.Context, req *input.ProjectReq) (*output.ProjectRes, error) {
	portFolioID := req.ProjectID
	fmt.Printf("what is the ID here? %d \n", portFolioID)
	res, err := s.storage.GetProject(ctx, portFolioID)
	if err != nil {
		return nil, err
	}

	return res, nil

}
