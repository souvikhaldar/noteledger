CREATE TABLE users(
	user_id INT NOT NULL AUTO_INCREMENT,
	name VARCHAR(40) NOT NULL,
	password VARCHAR(30) NOT NULL,
	email VARCHAR(50),
	PRIMARY KEY(user_id)
);