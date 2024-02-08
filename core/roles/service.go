package roles

import (
	"strconv"
)

type (
	Service interface {
		SaveUserRole(user_id, rol_id string) (UserRoles, error)
		UserExits(id string) error
		GetUserRoles(idUser, idRole string) (bool, error)
		RemoveUserRole(id string) (bool, error)
	}

	service struct {
		repo Repository
	}
)

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) SaveUserRole(user_id, rol_id string) (UserRoles, error) {
	userInd, err := strconv.Atoi(user_id)
	if err != nil {
		return UserRoles{}, err
	}
	roleInd, err := strconv.Atoi(rol_id)
	if err != nil {
		return UserRoles{}, err
	}
	userRol := UserRoles{
		UserID: userInd,
		RoleID: roleInd,
	}

	err = s.repo.SaveUserRole(&userRol)

	if err != nil {
		return userRol, err
	}
	return userRol, nil
}

func (s *service) UserExits(id string) error {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = s.repo.UserExits(idInt)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetUserRoles(idUser, idRole string) (bool, error) {

	idUserInt, err := strconv.Atoi(idUser)
	if err != nil {
		return false, err
	}
	idRoleInt, err := strconv.Atoi(idRole)
	if err != nil {
		return false, err
	}
	usersRole, err := s.repo.GetUserRoles(idUserInt)
	for _, v := range usersRole {
		if v.RoleID == idRoleInt {
			return false, err
		}
	}
	return true, nil
}

func (s *service) RemoveUserRole(id string) (bool, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return false, err
	}
	existRole, err := s.repo.RemoveUserRole(idInt)
	if err != nil {
		return existRole, err
	}
	return existRole, nil
}
