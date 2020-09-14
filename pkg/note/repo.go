package note

type Repository interface {
	AddNote(userID int, noteType string, noteContent string, bucket string) error
	GetNote(userID int, bucket string) ([]TextNote, error)
}

//func AddUserNote(
//	repo Repository,
//	userID int,
//	noteType string,
//	noteContent string,
//	bucket string,
//) error {
//	log.Println("Adding note from domain")
//	err := repo.AddNote(userID, noteType, noteContent, bucket)
//	return err
//}
