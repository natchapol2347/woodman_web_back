package input

import "time"

// CREATE TABLE IF NOT EXISTS Portfolio (
//     ProjectID INT PRIMARY KEY AUTO_INCREMENT,
//     ProjectName VARCHAR(255) NOT NULL,
//     Description TEXT,
//     CategoryID INT,
//     TagID INT,
//     CompletionDate DATE,
//     FOREIGN KEY (CategoryID) REFERENCES Category(CategoryID),
//     FOREIGN KEY (TagID) REFERENCES Tag(TagID)
// );
// CREATE TABLE IF NOT EXISTS PortfolioImages (
//     ImageID INT PRIMARY KEY AUTO_INCREMENT,
//     ProjectID INT NOT NULL,
//     ImageURL VARCHAR(255) NOT NULL,
//     FOREIGN KEY (ProjectID) REFERENCES Portfolio(ProjectID)
// );
// -- ContactForm Table
// CREATE TABLE IF NOT EXISTS ContactForm (
//
//	SubmissionID INT PRIMARY KEY AUTO_INCREMENT,
//	Name VARCHAR(255) NOT NULL,
//	Email VARCHAR(255) NOT NULL,
//	Message TEXT NOT NULL,
//	Timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
//
// );

// CREATE TABLE IF NOT EXISTS Job (
//     JobID INT PRIMARY KEY AUTO_INCREMENT,
//     Title VARCHAR(255) NOT NULL,
//     Description TEXT,
//     Requirements TEXT,
//     Location VARCHAR(255),
//     DatePosted DATE,
//     Status VARCHAR(255) DEFAULT 'Open'
// );

type PortfolioReq struct {
	ProjectID      int                  `json:"projectID"`
	ProjectName    string               `json:"projectName"`
	Description    string               `json:"description"`
	CompletionDate string               `json:"completionDate"`
	CategoryID     int                  `json:"categoryID"` //     FOREIGN KEY (CategoryID) REFERENCES Category(CategoryID),
	TagID          int                  `json:"TagID"`      //     FOREIGN KEY (TagID) REFERENCES Tag(TagID)
	Images         []PortfolioImagesReq `json:"images"`
}

type PortfolioImagesReq struct {
	ImageID   int    `json:"imageID"`
	ProjectID int    `json:"projectID"` //     		 FOREIGN KEY (ProjectID) REFERENCES Portfolio(ProjectID)
	ImageUrl  string `json:"imageUrl"`
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
