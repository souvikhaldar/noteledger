package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/souvikhaldar/noteledger/pkg/note"
)

func (db *DB) AddNote(
	userID int,
	noteType string,
	noteContent string,
	bucket string,
) error {
	var datetime = time.Now()
	datetime.Format(time.RFC3339)
	datetime, _ = time.Parse("2006-01-02", "2019-03-13")
	fmt.Println("TOday date: ", datetime)

	_, err := db.mysql.Exec(`
	INSERT INTO notes(
	user_id,
	note_type,
	note_content,
	bucket,
	timestamp) VALUES(?,?,?,?,NOW())
	`, userID, noteType, noteContent, bucket)

	return err
}

func (db *DB) GetNote(
	userID int,
	bucket string,
) ([]note.TextNote, error) {
	query := `
	SELECT 
		note_content,timestamp,bucket
	FROM 
		notes
	WHERE
		user_id=?
	`
	var arg []interface{}
	arg = append(arg, userID)
	if bucket != "" {
		query += `AND bucket=?`
		arg = append(arg, bucket)
	}
	rows, err := db.mysql.Query(
		query,
		arg...,
	)
	if err != nil {
		fmt.Println("Error in query: ", err)
		return nil, err
	}
	var notes []note.TextNote
	defer rows.Close()
	for rows.Next() {
		var ts sql.NullString
		var n note.TextNote
		if err := rows.Scan(&n.Note, &ts, &n.Bucket); err != nil {
			return notes, err
		}
		if ts.Valid {
			n.Timestamp = ts.String
		}
		notes = append(notes, n)
	}
	return notes, nil
}
