package user

import "gorm.io/gorm"

type (
	Repository interface {
		Register(user *User) error
		GetUserByMail(email string) (User, error)
	}

	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		db: db,
	}
}

func (repo *repo) Register(user *User) error {
	if err := repo.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetUserByMail(mail string) (User, error) {
	user := User{}
	err := repo.db.Where("email = ?", mail).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
