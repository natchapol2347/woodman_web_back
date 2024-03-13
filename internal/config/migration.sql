CREATE TABLE IF NOT EXISTS Project (
    ProjectID UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    ProjectName VARCHAR(255) NOT NULL,
    Description TEXT,
    CategoryID INT,
    TagID INT,
    CompletionDate DATE,
    CONSTRAINT fk_category FOREIGN KEY (CategoryID) REFERENCES Category(CategoryID),
    CONSTRAINT fk_tag FOREIGN KEY (TagID) REFERENCES Tag(TagID)
);

CREATE TABLE IF NOT EXISTS ProjectImage (
    ImageID  UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    ProjectID INT NOT NULL,
    ImageURL VARCHAR(255) NOT NULL,
    CONSTRAINT fk_portfolio FOREIGN KEY (ProjectID) REFERENCES Project(ProjectID)
);

CREATE TABLE IF NOT EXISTS ContactForm (
    SubmissionID  UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Email VARCHAR(255) NOT NULL,
    Message TEXT NOT NULL,
    Timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS Job (
    JobID  UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    Title VARCHAR(255) NOT NULL,
    Description TEXT,
    Requirements TEXT,
    Location VARCHAR(255),
    DatePosted DATE,
    Status VARCHAR(255) DEFAULT 'Open'
);

CREATE TABLE IF NOT EXISTS Category (
    CategoryID  UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    CategoryName VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS Tag (
    TagID  UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    TagName VARCHAR(255) NOT NULL
);
