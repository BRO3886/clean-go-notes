package user

import (
	"github.com/BRO3886/clean-go-notes/pkg"
	"github.com/jinzhu/gorm"
)

//repo

type repo struct {
	DB *gorm.DB
}

// func newRepo(db *gorm.DB) Repoistory {

// }

func (r *repo) FindByID(id float64) (*User, error) {
	tx := r.DB.Begin()
	user := &User{}
	if err := tx.Where("id=?", id).Find(id).Error; err != nil {
		tx.Rollback()
		return nil, pkg.ErrNotFound
	}
	tx.Commit()
	return user, nil
}

func (r *repo) Register(user *User) (*User, error) {
	tx := r.DB.Begin()
	if err := tx.Where("email=?", user.Email).Find(user).Error; err == nil {
		tx.Rollback()
		return nil, pkg.ErrAlreadyExists
	} else if err == gorm.ErrRecordNotFound {
		if err := tx.Save(user).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		tx.Commit()
		return user, nil
	} else {
		tx.Rollback()
		return nil, err
	}
}

func (r *repo) FindByEmail(email string) (*User, error) {
	tx := r.DB.Begin()
	user := &User{}
	if err := tx.Where("email=?", email).Find(user).Error; err != nil {
		tx.Rollback()
		return nil, pkg.ErrNotFound
	}
	tx.Commit()
	return user, nil
}