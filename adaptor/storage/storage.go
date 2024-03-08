package storage

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/natchapol2347/woodman_web_back/port/output"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

type IStorage interface {
	GetProject(ctx echo.Context, projectID int) (*output.ProjectRes, error)
	GetAllProjects(ctx echo.Context, limit string, offset string) ([]output.ProjectRes, error)
}

func (s *Storage) GetProject(ctx echo.Context, projectID int) (*output.ProjectRes, error) {
	// Query the database to retrieve the project entry
	project := &output.ProjectRes{}
	queryCtx := ctx.Request().Context()
	err := s.db.QueryRowContext(queryCtx, "SELECT ProjectID, ProjectName, Description, CompletionDate, CategoryID, Tagid FROM project WHERE ProjectID = $1", projectID).Scan(
		&project.ProjectID, &project.ProjectName, &project.Description, &project.CompletionDate, &project.CategoryID, &project.TagID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Return a specific error message if the data is not found
			return nil, output.NewErrorResponse(http.StatusNotFound, fmt.Sprintf("project not found for projectID %d", projectID), "")
		}
		// Return the actual error if it's not a "not found" error
		return nil, err
	}

	// Query the database to retrieve the images associated with the project entry
	rowsImage, err := s.db.QueryContext(queryCtx, "SELECT ImageID, ImageUrl FROM ProjectImage WHERE ProjectID = $1", projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to query images for projectID %d: %s", projectID, err)
	}
	defer rowsImage.Close()

	// Iterate over the rows and populate the Images slice
	for rowsImage.Next() {
		image := &output.ProjectImagesRes{}
		if err := rowsImage.Scan(&image.ImageID, &image.ImageUrl); err != nil {
			return nil, err
		}
		project.Images = append(project.Images, *image)
	}

	return project, nil

}

func (s *Storage) GetAllProjects(ctx echo.Context, limit string, offset string) ([]output.ProjectRes, error) {
	allProjects := []output.ProjectRes{}
	queryCtx := ctx.Request().Context()
	rows, err := s.db.QueryContext(queryCtx, "SELECT ProjectID, ProjectName, Description, CompletionDate, CategoryID, Tagid FROM project LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, output.NewErrorResponse(http.StatusNotFound, fmt.Sprintf("no projects found for limit %s and offset %s", limit, offset), "")
		}
		return nil, err
	}
	for rows.Next() {
		//map item to project struct
		item := &output.ProjectRes{}
		if err := rows.Scan(&item.ProjectID, &item.ProjectName, &item.Description, &item.CompletionDate, &item.CategoryID, &item.TagID); err != nil {
			return nil, err
		}
		//get rows images of that project(item)
		rowsImage, err := s.db.QueryContext(queryCtx, "SELECT ImageID, ImageUrl FROM ProjectImage WHERE ProjectID = $1", item.ProjectID)
		if err != nil {
			return nil, fmt.Errorf("failed to query images for projectID %d: %s", item.ProjectID, err)
		}
		defer rowsImage.Close()

		// Iterate over the rows and populate the Images slice
		for rowsImage.Next() {
			image := &output.ProjectImagesRes{}
			if err := rowsImage.Scan(&image.ImageID, &image.ImageUrl); err != nil {
				return nil, err
			}
			item.Images = append(item.Images, *image)
		}
		//append whole item to allProjects now
		allProjects = append(allProjects, *item)
	}
	return allProjects, nil

}
