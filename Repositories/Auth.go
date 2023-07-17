package repositories

import (
	models "mcdm/Models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	Login(username string) (models.User, error)
	GetUsers(ID int) (models.User, error)
	CheckEmail(email string) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *users {
	return &users{db}
}

func (r *users) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *users) CheckEmail(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}
func (r *users) Login(username string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "username=?", username).Error
	return user, err
}
func (r *users) GetUsers(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}
