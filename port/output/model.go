package output

import (
	"time"
)

type GetProjectRes struct {
	ProjectID      int                `json:"projectID"`
	ProjectName    string             `json:"projectName"`
	Description    string             `json:"description"`
	CompletionDate string             `json:"completionDate"`
	CategoryID     int                `json:"categoryID"` //     FOREIGN KEY (CategoryID) REFERENCES Category(CategoryID),
	TagID          int                `json:"TagID"`      //     FOREIGN KEY (TagID) REFERENCES Tag(TagID)
	Images         []ProjectImagesRes `json:"images"`
}

type MessageRes struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

type ProjectImagesRes struct {
	ImageID   int    `json:"imageID"`
	ProjectID int    `json:"projectID"` //     		 FOREIGN KEY (ProjectID) REFERENCES Project(ProjectID)
	ImageUrl  string `json:"imageUrl"`
}

type ContactFormRes struct {
	SubmissionID int       `json:"submissionID"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Message      string    `json:"message"`
	Timestamp    time.Time `json:"timeStamp"`
}

type JobRes struct {
	JobID        int    `json:"jobID"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Requirements string `json:"requirements"`
	Location     string `json:"location"`
	DatePosted   string `json:"datePosted"`
	Status       string `json:"status"`
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
