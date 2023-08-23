package repositories

import (
	models "mcdm/Models"

	"gorm.io/gorm"
)

type PerbandinganCriteriaAhpRepository interface {
	FindPerbandinganCriteriaAhp() ([]models.PerbandinganCriteriaAhp, error)
	CreatePerbandinganCriteriaAhp(createPerbandinganCriteriaAhp models.PerbandinganCriteriaAhp) (models.PerbandinganCriteriaAhp, error)
	GetPerbandinganCriteriaAhpByKode(kode string) (models.PerbandinganCriteriaAhp, error)
	UpdatePerbandinganCriteriaAhp(perbandinganCriteriaAhp models.PerbandinganCriteriaAhp) (models.PerbandinganCriteriaAhp, error)
	DeletePerbandinganCriteriaAhp(perbandinganCriteriaAhp models.PerbandinganCriteriaAhp) (models.PerbandinganCriteriaAhp, error)
}

func RepositoryPerbandinganCriteriaAhp(db *gorm.DB) *users {
	return &users{db}
}

func (r *users) FindPerbandinganCriteriaAhp() ([]models.PerbandinganCriteriaAhp, error) {
	var perbandinganCriteriaAhp []models.PerbandinganCriteriaAhp

	err := r.db.Order("CAST(kode AS INT) asc").Find(&perbandinganCriteriaAhp).Error

	return perbandinganCriteriaAhp, err
}

func (r *users) CreatePerbandinganCriteriaAhp(CreatePerbandinganCriteriaAhp models.PerbandinganCriteriaAhp) (models.PerbandinganCriteriaAhp, error) {
	err := r.db.Create(&CreatePerbandinganCriteriaAhp).Error

	return CreatePerbandinganCriteriaAhp, err
}

func (r *users) GetPerbandinganCriteriaAhpByKode(kode string) (models.PerbandinganCriteriaAhp, error) {
	var perbandinganCriteriaAhp models.PerbandinganCriteriaAhp

	err := r.db.Where("kode = ?", kode).First(&perbandinganCriteriaAhp).Error

	return perbandinganCriteriaAhp, err
}

func (r *users) UpdatePerbandinganCriteriaAhp(perbandinganCriteriaAhp models.PerbandinganCriteriaAhp) (models.PerbandinganCriteriaAhp, error) {
	err := r.db.Save(&perbandinganCriteriaAhp).Error

	return perbandinganCriteriaAhp, err
}

func (r *users) DeletePerbandinganCriteriaAhp(perbandinganCriteriaAhp models.PerbandinganCriteriaAhp) (models.PerbandinganCriteriaAhp, error) {
	err := r.db.Delete(&perbandinganCriteriaAhp).Error

	return perbandinganCriteriaAhp, err
}
