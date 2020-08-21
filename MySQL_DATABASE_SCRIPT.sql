CREATE DATABASE test_db;

USE test_db;

CREATE TABLE users (
	id INT auto_increment Primary key,
    email VARCHAR(30),
    pass VARCHAR(30),
    age VARCHAR(4),
    gender VARCHAR(10)
);