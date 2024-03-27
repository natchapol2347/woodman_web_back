package output

import (
	"time"

	"github.com/google/uuid"
)

type GetProjectRes struct {
	ProjectID      uuid.UUID          `json:"projectID"`
	ProjectName    string             `json:"projectName"`
	Description    string             `json:"description"`
	CompletionDate string             `json:"completionDate"`
	CategoryID     uuid.UUID          `json:"categoryID,omitempty"` //     FOREIGN KEY (CategoryID) REFERENCES Category(CategoryID),
	TagID          uuid.UUID          `json:"TagID,omitempty"`      //     FOREIGN KEY (TagID) REFERENCES Tag(TagID)
	Images         []ProjectImagesRes `json:"images,omitempty"`
}

// for POST operations
type MessageRes struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

type ProjectImagesRes struct {
	ImageID   uuid.UUID `json:"imageID"`
	ProjectID uuid.UUID `json:"projectID"` //     		 FOREIGN KEY (ProjectID) REFERENCES Project(ProjectID)
	ImageUrl  string    `json:"imageUrl"`
}

type ContactFormRes struct {
	SubmissionID uuid.UUID `json:"submissionID"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Message      string    `json:"message"`
	Timestamp    time.Time `json:"timeStamp"`
}

type GetJobRes struct {
	JobID          uuid.UUID `json:"jobID"`
	Title          string    `json:"title"`
	Status         string    `json:"status"`
	Salary         string    `json:"salary"`
	EmploymentType string    `json:"employmentType"`
}

type GetJobResAll struct {
	JobID          uuid.UUID `json:"jobID"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Requirements   string    `json:"requirements"`
	Location       string    `json:"location"`
	DatePosted     string    `json:"datePosted"`
	Status         string    `json:"status"`
	Salary         string    `json:"salary"`
	EmploymentType string    `json:"employmentType"`
}

type ErrorResponse struct {
	StatusCode int    // HTTP status code
	Message    string // Error message
	Details    string // Additional details
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

// NewCustomError creates a new CustomError instance
func NewErrorResponse(statusCode int, message, details string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
		Details:    details,
	}
}
