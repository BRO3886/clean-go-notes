package note

//Service layer for another abstraction over higher db func
type Service interface {
	CreateNote(note *Note) (*Note, error)
	FetchAllNotes(userID uint64) (*[]Note, error)
}

type service struct {
	repo Repository
}

//NewService func
func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetRepo() Repository {
	return s.repo
}

func (s service) CreateNote(note *Note) (*Note, error) {
	return s.repo.CreateNote(note)
}

func (s service) FetchAllNotes(userID uint64) (*[]Note, error) {
	return s.repo.GetAllNotes(userID)
}
