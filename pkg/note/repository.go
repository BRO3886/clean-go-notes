package note

//Repository abstraction layer over db
type Repository interface {
	CreateNote(note *Note) (*Note, error)
	GetAllNotes(userID uint64) (*[]Note, error)
}
