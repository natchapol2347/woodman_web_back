package storage

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/natchapol2347/woodman_web_back/port/input"
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
	GetProject(ctx echo.Context, projectID int) (*output.GetProjectRes, error)
	GetAllProjects(ctx echo.Context, limit string, offset string) ([]output.GetProjectRes, error)
	PostProject(ctx echo.Context, req *input.PostProjectReq) (*output.MessageRes, error)
}

func (s *Storage) GetProject(ctx echo.Context, projectID int) (*output.GetProjectRes, error) {
	// Query the database to retrieve the project entry
	project := &output.GetProjectRes{}
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

func (s *Storage) GetAllProjects(ctx echo.Context, limit string, offset string) ([]output.GetProjectRes, error) {
	allProjects := []output.GetProjectRes{}
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
		item := &output.GetProjectRes{}
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

func (s *Storage) PostProject(ctx echo.Context, req *input.PostProjectReq) (*output.MessageRes, error) {
	var projectID int = req.ProjectID
	queryCtx := ctx.Request().Context()
	execRes, err := s.db.ExecContext(queryCtx, "INSERT INTO project (ProjectID, ProjectName, Description, CompletionDate, CategoryID, Tagid) VALUES($1,$2,$3,$4,$5,$6)", projectID, req.ProjectName, req.Description, req.CompletionDate, req.CategoryID, req.TagID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			// Duplicate key error
			fmt.Println("am i not here?")
			details := fmt.Sprintf("duplicate ID: %d", projectID)
			return nil, output.NewErrorResponse(http.StatusConflict, "Duplicate key error", details)
		}
		return nil, err
	}
	rowsAffected, err := execRes.RowsAffected()
	if err != nil {
		return nil, err
	}
	// Construct a message based on the rows affected
	data := fmt.Sprintf("Inserted %d rows", rowsAffected)

	for _, v := range req.Images {
		//image projectID should always reference from project
		if v.ProjectID != projectID {
			details := fmt.Sprintf("projectID of project: %d, projectID from image %d", v.ProjectID, projectID)
			return nil, output.NewErrorResponse(http.StatusBadRequest, "ProjectID doesn't match", details)
		}

		// Upload image to S3 (replace 'your-s3-bucket' with your actual S3 bucket name)
		//  s3URL, err := s.uploadToS3("your-s3-bucket", v.ImageID, v.ImageData)
		//  if err != nil {
		// 	 return nil, err
		//  }
		_, err := s.db.ExecContext(queryCtx, "INSERT INTO projectimage (ImageID, ProjectID, ImageUrl) VALUES ($1,$2,$3)", v.ImageID, v.ProjectID, v.ImageUrl) //s3URL
		if err != nil {
			return nil, err
		}
	}

	msg := fmt.Sprintf("Insert to project (ID: %d ) successfully", projectID)

	response := &output.MessageRes{
		Message: msg,
		Data:    data,
	}
	return response, nil

}

// func (s *Storage) uploadToS3(bucketName string, imageID int, imageData []byte) (string, error) {
// 	// Upload imageData to S3 bucket 'bucketName' and return the URL
// 	// Example using AWS SDK for Go (replace 'your-region' with your actual AWS region)
// 	uploader := s3manager.NewUploaderWithClient(s3.New(s.session.NewSession(&aws.Config{
// 		Region: aws.String("your-region"),
// 	})))
// 	result, err := uploader.Upload(&s3manager.UploadInput{
// 		Bucket: aws.String(bucketName),
// 		Key:    aws.String(fmt.Sprintf("images/%d.jpg", imageID)),
// 		Body:   bytes.NewReader(imageData),
// 	})
// 	if err != nil {
// 		return "", err
// 	}
// 	return result.Location, nil
// }
