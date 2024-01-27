package user

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type (
	Service interface {
		Register(name, email, password string) error
		IsValidMail(email string) bool
		GetUserByMail(email string) (User, error)
		EncryptPassword(password string) (string, error)
		ValidPassword(email, password string) error
	}

	service struct {
		repo Repository
	}
)

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) Register(name, email, password string) error {
	passwordEncry, err := s.EncryptPassword(password)
	user := User{
		Name:     name,
		Email:    email,
		Password: passwordEncry,
	}
	err = s.repo.Register(&user)
	if err != nil {
		return err
	}
	return nil
}

func (s service) IsValidMail(email string) bool {
	validMail := regexp.MustCompile("^[_A-Za-z0-9-\\+]+(\\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\\.[A-Za-z0-9]+)*(\\.[A-Za-z]{2,})$")
	return validMail.MatchString(email)
}

func (s service) GetUserByMail(email string) (User, error) {
	return s.repo.GetUserByMail(email)
}

func (s service) EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s service) ValidPassword(email, password string) error {
	user, err := s.repo.GetUserByMail(email)

	if err != nil {
		return err
	}
	passwordByte := []byte(password)
	passwordDB := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(passwordDB, passwordByte)

	if err != nil {
		return err
	}
	return nil
}
