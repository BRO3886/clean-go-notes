package user

import (
	"errors"
	"fmt"
	"strings"

	"github.com/BRO3886/clean-go-notes/pkg"
	"golang.org/x/crypto/bcrypt"
)

//Service layer
type Service interface {
	Register(user *User) (*User, error)
	Login(email, password string) (*User, error)
	GetUserByID(id uint64) (*User, error)
	GetRepo() Repository
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

func (user *User) validate() (bool, error) {
	if !strings.Contains(user.Email, "@") {
		return false, pkg.ErrEmail
	}

	if len(user.Password) < 6 || len(user.Password) > 60 {
		return false, pkg.ErrPassword
	}

	return true, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *service) Register(user *User) (*User, error) {
	validation, err := user.validate()
	if !validation {
		return nil, err
	}

	emailExists := s.repo.DoesEmailExist(user.Email)
	if !emailExists {
		return nil, errors.New("Email already Exists")
	}

	user.Password, err = hashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	return s.repo.Register(user)
}

func (s *service) Login(email, password string) (*User, error) {
	user := &User{}
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	passHash, err := hashPassword(password)
	if err != nil {
		fmt.Println("error hashing password on login")
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(passHash), []byte(password)); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) GetUserByID(id uint64) (*User, error) {
	return s.repo.FindByID(id)
}
