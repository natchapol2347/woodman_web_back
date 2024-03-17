package input

import (
	"time"

	"github.com/google/uuid"
)

type ProjectImagesReq struct { //not used, but keep it for now just in case!!!
	ImageID  uuid.UUID `json:"imageID,omitempty"` //     		 FOREIGN KEY (ProjectID) REFERENCES Project(ProjectID)
	ImageUrl string    `json:"imageUrl"`
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

type PostJobReq struct {
	Title          string `json:"title"`
	Overview       string `json:"overview"`
	Qualification  string `json:"qualificatoin"`
	Responsibility string `json:"responsibility"`
}

type UpdateProjectReq struct {
	Action         string             `json:"action" validate:"required"`
	ProjectID      uuid.UUID          `json:"uuid.UUID" validate:"required"`
	ProjectName    string             `json:"projectName" validate:"required"`
	Description    string             `json:"description,omitempty"`
	CompletionDate string             `json:"completionDate" validate:"required"`
	CategoryID     *uuid.UUID         `json:"categoryID,omitempty"` //     FOREIGN KEY (CategoryID) REFERENCES Category(CategoryID),
	TagID          *uuid.UUID         `json:"TagID,omitempty"`      //     FOREIGN KEY (TagID) REFERENCES Tag(TagID)
	DeleteImages   []ProjectImagesReq `json:"deleteImages"`
	InsertImages   []ProjectImagesReq `json:"insertImages"`
}
