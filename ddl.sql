CREATE DATABASE score_db;
USE score_db;
CREATE TABLE users(
    id int PRIMARY KEY AUTO_INCREMENT, 
    name VARCHAR(100) NOT NULL, 
    score int DEFAULT NULL
);
INSERT INTO users(
    name, 
    score
) 
VALUES 
(
    'Player1', 
    100
), 
(
    'Player2', 
    1000
),
(
    'Player3',
    2300
);
