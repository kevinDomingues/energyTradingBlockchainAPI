CREATE TABLE Admin (
    adminID SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE Users (
    userID SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    birthdayDate DATE NOT NULL
);

CREATE TABLE Category (
    categoryID SERIAL PRIMARY KEY,
    categoryName VARCHAR(255) NOT NULL,
    categoryDescription TEXT
);

CREATE TABLE Event (
    eventID SERIAL PRIMARY KEY,
    ownerID INT NOT NULL REFERENCES Users (userID),
    categoryID INT NOT NULL REFERENCES Category (categoryID),
    eventName VARCHAR(255) NOT NULL,
    eventDescription TEXT,
    eventDate DATE NOT NULL,
    eventAgeRestriction INT
);

CREATE TABLE AttendeeStatus (
    attendeeStatusID SERIAL PRIMARY KEY,
    attendeeStatusName VARCHAR(255) NOT NULL
);

CREATE TABLE UserAttendees (
    userID INT NOT NULL REFERENCES Users (userID),
    eventID INT NOT NULL REFERENCES Event (eventID),
    attendeeStatusID INT NOT NULL REFERENCES AttendeeStatus (attendeeStatusID),
    dateChanged DATE NOT NULL,
    PRIMARY KEY (userID, eventID)
);

CREATE TABLE Feedback (
    feedbackID SERIAL PRIMARY KEY,
    eventID INT NOT NULL REFERENCES Event (eventID),
    userID INT NOT NULL REFERENCES Users (userID),
    feedback TEXT NOT NULL,
    feedbackDate DATE NOT NULL
);

CREATE TABLE Roles (
    roleID SERIAL PRIMARY KEY,
    roleName VARCHAR(255) NOT NULL
);

CREATE TABLE UserRoles (
    userID INT NOT NULL REFERENCES Users (userID),
    roleID INT NOT NULL REFERENCES Roles (roleID),
    PRIMARY KEY (userID, roleID)
);

CREATE TABLE Favorites (
    userID INT NOT NULL REFERENCES Users (userID),
    eventID INT NOT NULL REFERENCES Event (eventID),
    active BOOLEAN NOT NULL,
    PRIMARY KEY (userID, eventID)
);
