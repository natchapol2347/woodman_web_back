package service

import (
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
	GetProject(ctx echo.Context, req *input.GetProjectReq) (*output.GetProjectRes, error)
	GetAllProjects(ctx echo.Context) ([]output.GetProjectRes, error)
	PostProject(ctx echo.Context, req *input.PostProjectReq) (*output.MessageRes, error)
}

func (s *Service) GetProject(ctx echo.Context, req *input.GetProjectReq) (*output.GetProjectRes, error) {
	var portFolioID int = req.ProjectID
	res, err := s.storage.GetProject(ctx, portFolioID)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (s *Service) GetAllProjects(ctx echo.Context) ([]output.GetProjectRes, error) {
	limit := ctx.QueryParam("limit")
	offset := ctx.QueryParam("offset")
	res, err := s.storage.GetAllProjects(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (s *Service) PostProject(ctx echo.Context, req *input.PostProjectReq) (*output.MessageRes, error) {
	res, err := s.storage.PostProject(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil

}
