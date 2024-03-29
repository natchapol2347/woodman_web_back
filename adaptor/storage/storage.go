package storage

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
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
	GetProject(ctx echo.Context, projectID uuid.UUID) (*output.GetProjectRes, error)
	GetManyProjects(ctx echo.Context) ([]output.GetProjectRes, error)
	PostProject(ctx echo.Context, req *input.PostProjectReq) (*output.MessageRes, error)
	DeleteProject(ctx echo.Context, projectID uuid.UUID) (*output.MessageRes, error)
	UpdateProject(ctx echo.Context, req *input.UpdateProjectReq) (*output.MessageRes, error)
	GetManyJobs(ctx echo.Context) ([]output.GetJobResAll, error)
}

func (s *Storage) GetProject(ctx echo.Context, projectID uuid.UUID) (*output.GetProjectRes, error) {
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
	rowsImage, err := s.db.QueryContext(queryCtx, "SELECT ImageID, ProjectID, ImageUrl FROM ProjectImage WHERE ProjectID = $1", projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to query images for projectID %d: %s", projectID, err)
	}
	defer rowsImage.Close()

	// Iterate over the rows and populate the Images slice
	for rowsImage.Next() {
		image := &output.ProjectImagesRes{}
		if err := rowsImage.Scan(&image.ImageID, &image.ProjectID, &image.ImageUrl); err != nil {
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
			return nil, err
		}
		//get rows images of that project(item)
		rowsImage, err := s.db.QueryContext(queryCtx, "SELECT ImageID, ProjectID, ImageUrl FROM ProjectImage WHERE ProjectID = $1", item.ProjectID)
		if err != nil {
			return nil, fmt.Errorf("failed to query images for projectID %d: %s", item.ProjectID, err)
		}
		defer rowsImage.Close()

		// Iterate over the rows and populate the Images slice
		for rowsImage.Next() {
			image := &output.ProjectImagesRes{}
			if err := rowsImage.Scan(&image.ImageID, &image.ProjectID, &image.ImageUrl); err != nil {
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
	var projectID string
	queryCtx := ctx.Request().Context()
	err := s.db.QueryRowContext(queryCtx, "INSERT INTO project (ProjectName, Description, CompletionDate, CategoryID, TagID) VALUES($1,$2,$3,$4,$5) RETURNING ProjectID", req.ProjectName, req.Description, req.CompletionDate, req.CategoryID, req.TagID).Scan(&projectID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			// Duplicate key error
			details := fmt.Sprintf("duplicate ID: %s", projectID)
			return nil, output.NewErrorResponse(http.StatusConflict, "Duplicate key error", details)
		}
		return nil, err
	}
	uuidProjID, err := uuid.Parse(projectID)
	if err != nil {
		return nil, err
	}
	for _, v := range req.Images {

		// Upload image to S3 (replace 'your-s3-bucket' with your actual S3 bucket name)
		//  s3URL, err := s.uploadToS3("your-s3-bucket", v.ImageID, v.ImageData)
		//  if err != nil {
		// 	 return nil, err
		//  }
		_, err := s.db.ExecContext(queryCtx, "INSERT INTO projectimage (ProjectID, ImageUrl) VALUES ($1,$2)", uuidProjID, v.ImageUrl) //s3URL
		if err != nil {
			return nil, err
		}
	}

	msg := "Insert project successfully"
	data := fmt.Sprintf("Project ID: %s", projectID)
	response := &output.MessageRes{
		Message: msg,
		Data:    data,
	}
	return response, nil

}

func (s *Storage) UpdateProject(ctx echo.Context, req *input.UpdateProjectReq) (*output.MessageRes, error) {
	queryCtx := ctx.Request().Context()
	var projectID string = ctx.Param("id")
	uuidProjID, err := uuid.Parse(projectID)
	if err != nil {
		return nil, err
	}
	// Construct the UPDATE query
	query := "UPDATE project SET "
	params := []interface{}{}

	var index int = 1
	if req.ProjectName != "" && req.Description != "" && req.CompletionDate != "" && req.CategoryID != uuid.Nil && req.TagID != uuid.Nil {
		if req.ProjectName != "" {
			query += "ProjectName = $" + strconv.Itoa(index) + ", "
			params = append(params, req.ProjectName)
			index++
		}
		if req.Description != "" {
			query += "Description = $" + strconv.Itoa(index) + ", "
			params = append(params, req.Description)
			index++
		}
		if req.CompletionDate != "" {
			query += "CompletionDate = $" + strconv.Itoa(index) + ", "
			params = append(params, req.CompletionDate)
			index++
		}
		if req.CategoryID != uuid.Nil {
			query += "CategoryID = $" + strconv.Itoa(index) + ", "
			params = append(params, req.CategoryID)
			index++
		}
		if req.TagID != uuid.Nil {
			query += "TagID = $" + strconv.Itoa(index) + ", "
			params = append(params, req.TagID)
			index++
		}

		// Remove the trailing comma and space
		query = query[:len(query)-2]

		// Add the WHERE clause
		query += " WHERE ProjectID = $" + strconv.Itoa(index)
		params = append(params, uuidProjID)

		// Execute the UPDATE query
		_, err := s.db.ExecContext(queryCtx, query, params...)
		if err != nil {
			return nil, err
		}

	}

	// Handle image uploads and deletions
	for _, v := range req.InsertImages {

		// Upload image to S3 and insert URL into database
		_, err := s.db.ExecContext(queryCtx, "INSERT INTO projectimage (ProjectID, ImageUrl) VALUES ($1,$2)", uuidProjID, v.ImageUrl) //s3URL
		if err != nil {
			fmt.Println("err!", err)
			return nil, err

		}
	}

	for _, v := range req.DeleteImages {
		_, err := s.db.ExecContext(queryCtx, "DELETE FROM projectimage WHERE ImageID = $1", v.ImageID)
		if err != nil {
			return nil, err
		}
	}

	msg := fmt.Sprintf("Update to project (ID: %s) successfully", projectID)
	data := fmt.Sprintf("inserted %d images and deleted %d images", len(req.InsertImages), len(req.DeleteImages))

	response := &output.MessageRes{
		Message: msg,
		Data:    data,
	}
	return response, nil
}

func (s *Storage) DeleteProject(ctx echo.Context, projectID uuid.UUID) (*output.MessageRes, error) {
	queryCtx := ctx.Request().Context()
	_, err := s.db.ExecContext(queryCtx, "DELETE FROM project WHERE ProjectID = $1", projectID)
	if err != nil {
		// Handle error
		return nil, err
	}

	msg := "Delete project successfully"
	data := fmt.Sprintf("Project ID: %s", projectID.String())

	response := &output.MessageRes{
		Message: msg,
		Data:    data,
	}
	return response, nil
}

// make separate struct, many jobs and get one job
func (s *Storage) GetManyJobs(ctx echo.Context) ([]output.GetJobResAll, error) {
	queryCtx := ctx.Request().Context()
	allJobs := []output.GetJobResAll{}
	rows, err := s.db.QueryContext(queryCtx, "SELECT JobID, Title, Description, Requirements, Location, Dateposted, Status, Salary, EmploymentType FROM job")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, output.NewErrorResponse(http.StatusNotFound, "no jobs found", "")
		}
		return nil, err
	}
	for rows.Next() {
		//map item to project struct
		item := &output.GetJobResAll{}
		if err := rows.Scan(&item.JobID, &item.Title, &item.Description, &item.Requirements, &item.Location, &item.DatePosted, &item.Status, &item.Salary, &item.EmploymentType); err != nil {
			return nil, err
		}

		allJobs = append(allJobs, *item)
	}
	return allJobs, nil
}
