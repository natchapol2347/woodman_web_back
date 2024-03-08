package storage

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

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
	GetManyProjects(ctx echo.Context) ([]output.GetProjectRes, error)
	PostProject(ctx echo.Context, req *input.PostProjectReq) (*output.MessageRes, error)
}

func (s *Storage) GetProject(ctx echo.Context, projectID int) (*output.GetProjectRes, error) {
	// Query the database to retrieve the project entry
	project := &output.GetProjectRes{}
	queryCtx := ctx.Request().Context()

	err := s.db.QueryRowContext(queryCtx, "SELECT ProjectID, ProjectName, Description, CompletionDate, CategoryID, TagID FROM project WHERE ProjectID = $1", projectID).Scan(
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

func (s *Storage) GetManyProjects(ctx echo.Context) ([]output.GetProjectRes, error) {

	limit, err := strconv.Atoi(ctx.QueryParam("limit"))
	fmt.Printf("limit %d \n", limit)
	if err != nil {

		limit = 10 // Default limit
	}

	offset, err := strconv.Atoi(ctx.QueryParam("offset"))
	fmt.Printf("offset %d \n", offset)

	if err != nil {

		offset = 0 // Default offset
	}

	tag := ctx.QueryParam("tagID")
	category := ctx.QueryParam("categoryID")

	allProjects := []output.GetProjectRes{}
	queryCtx := ctx.Request().Context()
	baseQuery := "SELECT ProjectID, ProjectName, Description, CompletionDate, CategoryID, TagID FROM project"
	queryParams := []interface{}{}
	whereClause := ""

	// concatenate filter of tag and category permutation
	if tag != "" {
		whereClause += "WHERE TagID = $1"
		queryParams = append(queryParams, tag)
	} else {
		fmt.Printf("tag? %s \n", tag)
	}
	if category != "" {
		if whereClause != "" {
			whereClause += " AND "
		} else {
			whereClause += "WHERE "
		}
		whereClause += "CategoryID = $" + strconv.Itoa(len(queryParams)+1)
		queryParams = append(queryParams, category)
	} else {
		fmt.Printf("category? %s \n", category)
	}
	// Add space if WHERE exist
	if whereClause != "" {
		baseQuery += " " + whereClause
	}

	// Add limit and offset
	baseQuery += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

	rows, err := s.db.QueryContext(queryCtx, baseQuery, queryParams...)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, output.NewErrorResponse(http.StatusNotFound, fmt.Sprintf("no projects found for limit %d and offset %d", limit, offset), "")
		}
		return nil, err
	}
	for rows.Next() {
		//map item to project struct
		item := &output.GetProjectRes{}
		if err := rows.Scan(&item.ProjectID, &item.ProjectName, &item.Description, &item.CompletionDate, &item.CategoryID, &item.TagID); err != nil {
			fmt.Println("taek!")
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
	execRes, err := s.db.ExecContext(queryCtx, "INSERT INTO project (ProjectID, ProjectName, Description, CompletionDate, CategoryID, TagID) VALUES($1,$2,$3,$4,$5,$6)", projectID, req.ProjectName, req.Description, req.CompletionDate, req.CategoryID, req.TagID)
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

func (s *Storage) UpdateProject(ctx echo.Context, req *input.PostProjectReq) (*output.MessageRes, error) {
	queryCtx := ctx.Request().Context()

	// Construct the UPDATE query
	query := "UPDATE project SET "
	params := []interface{}{}

	if req.ProjectName != "" {
		query += "ProjectName = $1, "
		params = append(params, req.ProjectName)
	}
	if req.Description != "" {
		query += "Description = $2, "
		params = append(params, req.Description)
	}
	if req.CompletionDate != "" {
		query += "CompletionDate = $3, "
		params = append(params, req.CompletionDate)
	}
	if req.CategoryID != 0 {
		query += "CategoryID = $4, "
		params = append(params, req.CategoryID)
	}
	if req.TagID != 0 {
		query += "TagID = $5, "
		params = append(params, req.TagID)
	}

	// Remove the trailing comma and space
	query = query[:len(query)-2]

	// Add the WHERE clause
	query += " WHERE ProjectID = $6"
	params = append(params, req.ProjectID)

	// Execute the UPDATE query
	_, err := s.db.ExecContext(queryCtx, query, params...)
	if err != nil {
		return nil, err
	}

	// Handle image uploads and deletions
	// for _, v := range req.Images {
	// 	if v.Action == "upload" {
	// 		// Upload image to S3 and insert URL into database
	// 		// Your implementation here
	// 	} else if v.Action == "delete" {
	// 		// Delete image from S3 and database
	// 		// Your implementation here
	// 	}
	// }

	msg := fmt.Sprintf("Update to project (ID: %d) successfully", req.ProjectID)

	response := &output.MessageRes{
		Message: msg,
	}
	return response, nil
}
