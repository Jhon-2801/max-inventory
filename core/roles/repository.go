package roles

import (
	"github.com/Jhon-2801/max-inventory/core/user"
	"gorm.io/gorm"
)

type (
	Repository interface {
		SaveUserRole(roles *UserRoles) error
		UserById(id int) error
		RemoveUserRole(roles *UserRoles) error
	}

	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (repo *repo) SaveUserRole(roles *UserRoles) error {
	if err := repo.db.Create(roles).Error; err != nil {
		return err
	}
	return nil
}
func (repo *repo) UserById(id int) error {
	user := user.User{
		Id: id,
	}
	err := repo.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) RemoveUserRole(roles *UserRoles) error {
	return nil
}
