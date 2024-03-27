package service

import (
	"github.com/google/uuid"
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
	GetProject(ctx echo.Context) (*output.GetProjectRes, error)
	GetManyProjects(ctx echo.Context) ([]output.GetProjectRes, error)
	PostProject(ctx echo.Context, req *input.PostProjectReq) (*output.MessageRes, error)
	DeleteProject(ctx echo.Context) (*output.MessageRes, error)
	UpdateProject(ctx echo.Context, req *input.UpdateProjectReq) (*output.MessageRes, error)
	GetManyJobs(ctx echo.Context) ([]output.GetJobResAll, error)
}

func (s *Service) GetProject(ctx echo.Context) (*output.GetProjectRes, error) {
	projectID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return nil, err
	}
	res, err := s.storage.GetProject(ctx, projectID)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (s *Service) GetManyProjects(ctx echo.Context) ([]output.GetProjectRes, error) {

	res, err := s.storage.GetManyProjects(ctx)
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

func (s *Service) DeleteProject(ctx echo.Context) (*output.MessageRes, error) {
	projectID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return nil, err
	}
	res, err := s.storage.DeleteProject(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) UpdateProject(ctx echo.Context, req *input.UpdateProjectReq) (*output.MessageRes, error) {
	res, err := s.storage.UpdateProject(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) GetManyJobs(ctx echo.Context) ([]output.GetJobResAll, error) {
	res, err := s.storage.GetManyJobs(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
