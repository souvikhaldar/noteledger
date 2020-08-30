CREATE TABLE notes(
	note_id INT NOT NULL AUTO_INCREMENT,
	user_id INT NOT NULL,
	note_type VARCHAR(10) NOT NULL,
	note_content VARCHAR(1000),
	bucket VARCHAR(10),
	PRIMARY KEY(note_id),
	FOREIGN KEY(user_id) REFERENCES users(user_id)
);


