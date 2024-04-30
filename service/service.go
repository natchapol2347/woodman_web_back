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
	GetProject(ctx echo.Context, projectID uuid.UUID) (*output.GetProjectRes, error)
	GetManyProjects(ctx echo.Context) ([]output.GetProjectRes, error)
	PostProject(ctx echo.Context, req *input.PostProjectReq) (*output.MessageRes, error)
	DeleteProject(ctx echo.Context, projectID uuid.UUID) (*output.MessageRes, error)
	UpdateProject(ctx echo.Context, req *input.UpdateProjectReq, projectID string) (*output.MessageRes, error)
	GetManyJobs(ctx echo.Context) ([]output.GetManyJobRes, error)
	PostJob(ctx echo.Context, req *input.PostJobReq) (*output.MessageRes, error)
	GetJob(ctx echo.Context, jobID uuid.UUID) (*output.GetJobRes, error)
	DeleteJob(ctx echo.Context, jobID uuid.UUID) (*output.MessageRes, error)
}

func (s *Service) GetProject(ctx echo.Context, projectID uuid.UUID) (*output.GetProjectRes, error) {

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

func (s *Service) DeleteProject(ctx echo.Context, projectID uuid.UUID) (*output.MessageRes, error) {

	res, err := s.storage.DeleteProject(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) UpdateProject(ctx echo.Context, req *input.UpdateProjectReq, projectID string) (*output.MessageRes, error) {
	res, err := s.storage.UpdateProject(ctx, req, projectID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) GetManyJobs(ctx echo.Context) ([]output.GetManyJobRes, error) {
	res, err := s.storage.GetManyJobs(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) PostJob(ctx echo.Context, req *input.PostJobReq) (*output.MessageRes, error) {
	res, err := s.storage.PostJob(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *Service) GetJob(ctx echo.Context, jobID uuid.UUID) (*output.GetJobRes, error) {

	res, err := s.storage.GetJob(ctx, jobID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) DeleteJob(ctx echo.Context, jobID uuid.UUID) (*output.MessageRes, error) {

	res, err := s.storage.DeleteJob(ctx, jobID)
	if err != nil {
		return nil, err
	}
	return res, nil
}
