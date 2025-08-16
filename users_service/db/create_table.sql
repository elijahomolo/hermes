CREATE TABLE Users (
    ID int PRIMARY KEY,
    LastName varchar(255),
    FirstName varchar(255),
    DateOfBirth date,
    Country varchar(255),
    Language varchar(255),
    Email varchar(255),
    Password varchar(255),
    CreatedAt datetime,
);