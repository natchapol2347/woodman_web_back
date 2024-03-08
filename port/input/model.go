package input

import "time"

type GetProjectReq struct {
	ProjectID int `json:"projectID"`
}

type ProjectImagesReq struct { //not used, but keep it for now just in case!!!
	ImageID   int    `json:"imageID"`
	ProjectID int    `json:"projectID"` //     		 FOREIGN KEY (ProjectID) REFERENCES Project(ProjectID)
	ImageUrl  string `json:"imageUrl"`
}

type AllProjectsReq struct {
}

type PostProjectReq struct {
	ProjectID      int                `json:"projectID"`
	ProjectName    string             `json:"projectName"`
	Description    string             `json:"description"`
	CompletionDate string             `json:"completionDate"`
	CategoryID     int                `json:"categoryID"` //     FOREIGN KEY (CategoryID) REFERENCES Category(CategoryID),
	TagID          int                `json:"TagID"`      //     FOREIGN KEY (TagID) REFERENCES Tag(TagID)
	Images         []ProjectImagesReq `json:"images"`
}

type ContactFormReq struct {
	SubmissionID int       `json:"submissionID"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Message      string    `json:"message"`
	Timestamp    time.Time `json:"timeStamp"`
}

type JobReq struct {
	JobID        int    `json:"jobID"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Requirements string `json:"requirements"`
	Location     string `json:"location"`
	DatePosted   string `json:"datePosted"`
	Status       string `json:"status"`
}
