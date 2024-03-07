package storage

import (
	"database/sql"
	"fmt"

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
	GetPortfolio(ctx echo.Context, projectID int) (*output.PortfolioRes, error)
}

func (s *Storage) GetPortfolio(ctx echo.Context, projectID int) (*output.PortfolioRes, error) {
	// Query the database to retrieve the portfolio entry
	portfolio := &output.PortfolioRes{}
	queryCtx := ctx.Request().Context()
	err := s.db.QueryRowContext(queryCtx, "SELECT ProjectID, ProjectName, Description, CompletionDate, CategoryID, tagid FROM Portfolio WHERE ProjectID = $1", projectID).Scan(
		&portfolio.ProjectID, &portfolio.ProjectName, &portfolio.Description, &portfolio.CompletionDate, &portfolio.CategoryID, &portfolio.TagID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Return a specific error message if the data is not found
			return nil, fmt.Errorf("portfolio not found for projectID %d", projectID)
		}
		// Return the actual error if it's not a "not found" error
		return nil, err
	}

	// Query the database to retrieve the images associated with the portfolio entry
	rows, err := s.db.QueryContext(queryCtx, "SELECT ImageID, ImageUrl FROM PortfolioImages WHERE ProjectID = $1", projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to query images for projectID %d: %s", projectID, err)
	}
	defer rows.Close()

	// Iterate over the rows and populate the Images slice
	for rows.Next() {
		image := &output.PortfolioImagesRes{}
		if err := rows.Scan(&image.ImageID, &image.ImageUrl); err != nil {
			return nil, err
		}
		portfolio.Images = append(portfolio.Images, *image)
	}

	return portfolio, nil

}
