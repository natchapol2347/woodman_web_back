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
}

func (s *Storage) GetProject(ctx echo.Context, projectID int) (*output.ProjectRes, error) {
	// Query the database to retrieve the project entry
	project := &output.ProjectRes{}
	queryCtx := ctx.Request().Context()
	err := s.db.QueryRowContext(queryCtx, "SELECT ProjectID, ProjectName, Description, CompletionDate, CategoryID, tagid FROM project WHERE ProjectID = $1", projectID).Scan(
		&project.ProjectID, &project.ProjectName, &project.Description, &project.CompletionDate, &project.CategoryID, &project.TagID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Return a specific error message if the data is not found
			fmt.Println("here??")
			return nil, output.NewErrorResponse(http.StatusNotFound, fmt.Sprintf("project not found for projectID %d", projectID), "")
		}
		// Return the actual error if it's not a "not found" error
		return nil, err
	}

	// Query the database to retrieve the images associated with the project entry
	rows, err := s.db.QueryContext(queryCtx, "SELECT ImageID, ImageUrl FROM ProjectImage WHERE ProjectID = $1", projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to query images for projectID %d: %s", projectID, err)
	}
	defer rows.Close()

	// Iterate over the rows and populate the Images slice
	for rows.Next() {
		image := &output.ProjectImagesRes{}
		if err := rows.Scan(&image.ImageID, &image.ImageUrl); err != nil {
			return nil, err
		}
		project.Images = append(project.Images, *image)
	}

	return project, nil

}
