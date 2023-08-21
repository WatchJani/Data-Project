CREATE TABLE team.users (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255),
    LastName VARCHAR(255),
    Password VARCHAR(255),
    Phone VARCHAR(20),
    Mail VARCHAR(255),
    UserName VARCHAR(255),
    BornDate DATE,
    Created TIMESTAMP DEFAULT current_timestamp
);