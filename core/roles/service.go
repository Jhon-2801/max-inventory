package roles

import (
	"strconv"
)

type (
	Service interface {
		SaveUserRole(user_id, rol_id string) (UserRoles, error)
		UserByID(id string) error
	}

	service struct {
		repo Repository
	}
)

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) SaveUserRole(user_id, rol_id string) (UserRoles, error) {
	userString, err := strconv.Atoi(user_id)
	if err != nil {
		return UserRoles{}, err
	}
	roleString, err := strconv.Atoi(rol_id)
	if err != nil {
		return UserRoles{}, err
	}
	userRol := UserRoles{
		UserID: userString,
		RoleID: roleString,
	}

	err = s.repo.SaveUserRole(&userRol)

	if err != nil {
		return userRol, err
	}
	return userRol, nil
}

func (s *service) UserByID(id string) error {
	idString, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = s.repo.UserById(idString)

	if err != nil {
		return err
	}

	return nil
}
