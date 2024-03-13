package input

import (
	"time"

	"github.com/google/uuid"
)

type GetProjectReq struct {
	ProjectID uuid.UUID `json:"projectID" validate:"required"`
}

type ProjectImagesReq struct { //not used, but keep it for now just in case!!!
	ImageID   uuid.UUID `json:"imageID"`
	ProjectID uuid.UUID `json:"projectID"` //     		 FOREIGN KEY (ProjectID) REFERENCES Project(ProjectID)
	ImageUrl  string    `json:"imageUrl"`
}

type AllProjectsReq struct {
}

type PostProjectReq struct {
	ProjectName    string             `json:"projectName" validate:"required"`
	Description    string             `json:"description,omitempty"`
	CompletionDate string             `json:"completionDate" validate:"required"`
	CategoryID     *uuid.UUID         `json:"categoryID,omitempty"` //     FOREIGN KEY (CategoryID) REFERENCES Category(CategoryID),
	TagID          *uuid.UUID         `json:"TagID,omitempty"`      //     FOREIGN KEY (TagID) REFERENCES Tag(TagID)
	Images         []ProjectImagesReq `json:"images"`
}

type ContactFormReq struct {
	// SubmissionID uuid.UUID `json:"submissionID"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timeStamp"`
}

type GetJobReq struct {
}

type PostJobReq struct {
	Title          string `json:"title"`
	Overview       string `json:"overview"`
	Qualification  string `json:"qualificatoin"`
	Responsibility string `json:"responsibility"`
}

type UpdateProjectReq struct {
	ProjectID      uuid.UUID          `json:"uuid.UUID" validate:"required"`
	ProjectName    string             `json:"projectName" validate:"required"`
	Description    string             `json:"description,omitempty"`
	CompletionDate string             `json:"completionDate" validate:"required"`
	CategoryID     *uuid.UUID         `json:"categoryID,omitempty"` //     FOREIGN KEY (CategoryID) REFERENCES Category(CategoryID),
	TagID          *uuid.UUID         `json:"TagID,omitempty"`      //     FOREIGN KEY (TagID) REFERENCES Tag(TagID)
	Images         []ProjectImagesReq `json:"images"`
}
