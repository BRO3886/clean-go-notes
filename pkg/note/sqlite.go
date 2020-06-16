package note

import (
	"github.com/BRO3886/clean-go-notes/pkg"
	"github.com/jinzhu/gorm"
)

type repo struct {
	DB *gorm.DB
}

//NewSqliteRepo creation
func NewSqliteRepo(db *gorm.DB) Repository {
	return &repo{DB: db}
}

func (r *repo) CreateNote(note *Note) (*Note, error) {
	tx := r.DB.Begin()
	if err := tx.Create(note).Error; err != nil {
		tx.Rollback()
		return nil, pkg.ErrDatabase
	}
	tx.Commit()
	return note, nil
}

func (r *repo) GetAllNotes(userID uint64) (*[]Note, error) {
	var notes []Note
	tx := r.DB.Begin()
	if err := tx.Where("user_id=?", userID).Find(notes).Error; err != nil {
		tx.Rollback()
		return nil, pkg.ErrDatabase
	}
	tx.Commit()
	return &notes, nil
}

func (r *repo) DeleteNote(id uint64) (bool, error) {
	//TODO: complete func
	return true, nil
}
