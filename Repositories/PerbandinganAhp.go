package repositories

import (
	models "mcdm/Models"

	"gorm.io/gorm"
)

type PerbandinganAhpRepository interface {
	FindPerbandinganAhp() ([]models.PerbandinganAhp, error)
	CreatePerbandinganAhp(createPerbandinganAhp models.PerbandinganAhp) (models.PerbandinganAhp, error)
	GetPerbandinganAhpByKode(kode string) (models.PerbandinganAhp, error)
	UpdatePerbandinganAhp(perbandinganAhp models.PerbandinganAhp) (models.PerbandinganAhp, error)
	DeletePerbandinganAhp(perbandinganAhp models.PerbandinganAhp) (models.PerbandinganAhp, error)
}

func RepositoryPerbandinganAhp(db *gorm.DB) *users {
	return &users{db}
}

func (r *users) FindPerbandinganAhp() ([]models.PerbandinganAhp, error) {
	var perbandinganAhp []models.PerbandinganAhp

	err := r.db.Order("CAST(kode AS INT) asc").Find(&perbandinganAhp).Error

	return perbandinganAhp, err
}

func (r *users) CreatePerbandinganAhp(CreatePerbandinganAhp models.PerbandinganAhp) (models.PerbandinganAhp, error) {
	err := r.db.Create(&CreatePerbandinganAhp).Error

	return CreatePerbandinganAhp, err
}

func (r *users) GetPerbandinganAhpByKode(kode string) (models.PerbandinganAhp, error) {
	var perbandinganAhp models.PerbandinganAhp

	err := r.db.Where("kode = ?", kode).First(&perbandinganAhp).Error

	return perbandinganAhp, err
}

func (r *users) UpdatePerbandinganAhp(perbandinganAhp models.PerbandinganAhp) (models.PerbandinganAhp, error) {
	err := r.db.Save(&perbandinganAhp).Error

	return perbandinganAhp, err
}

func (r *users) DeletePerbandinganAhp(perbandinganAhp models.PerbandinganAhp) (models.PerbandinganAhp, error) {
	err := r.db.Delete(&perbandinganAhp).Error

	return perbandinganAhp, err
}
