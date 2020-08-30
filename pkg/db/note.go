package db

func (db *DB) AddNote(
	userID int,
	noteType string,
	noteContent string,
	bucket string,
) error {
	_, err := db.mysql.Exec(`
	INSERT INTO notes(
	user_id,
	note_type,
	note_content,
	bucket) VALUES(?,?,?,?)
	`, userID, noteType, noteContent, bucket)

	return err
}
