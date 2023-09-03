package repositories

import (
	models "mcdm/Models"

	"gorm.io/gorm"
)

type PerbandinganMopaRepository interface {
	FindPerbandinganMopa() ([]models.PerbandinganMopa, error)
	CreatePerbandinganMopa(createPerbandinganMopa models.PerbandinganMopa) (models.PerbandinganMopa, error)
	GetPerbandinganMopaByKode(kode string) (models.PerbandinganMopa, error)
	UpdatePerbandinganMopa(perbandinganMopa models.PerbandinganMopa) (models.PerbandinganMopa, error)
	DeletePerbandinganMopa(perbandinganMopa models.PerbandinganMopa) (models.PerbandinganMopa, error)
}

func RepositoryPerbandinganMopa(db *gorm.DB) *users {
	return &users{db}
}

func (r *users) FindPerbandinganMopa() ([]models.PerbandinganMopa, error) {
	var perbandinganMopa []models.PerbandinganMopa

	err := r.db.Order("CAST(SUBSTRING(kode FROM 2) AS UNSIGNED) ASC").Find(&perbandinganMopa).Error

	return perbandinganMopa, err
}

func (r *users) CreatePerbandinganMopa(CreatePerbandinganMopa models.PerbandinganMopa) (models.PerbandinganMopa, error) {
	err := r.db.Create(&CreatePerbandinganMopa).Error

	return CreatePerbandinganMopa, err
}

func (r *users) GetPerbandinganMopaByKode(kode string) (models.PerbandinganMopa, error) {
	var perbandinganMopa models.PerbandinganMopa

	err := r.db.Where("kode = ?", kode).First(&perbandinganMopa).Error

	return perbandinganMopa, err
}

func (r *users) UpdatePerbandinganMopa(perbandinganMopa models.PerbandinganMopa) (models.PerbandinganMopa, error) {
	err := r.db.Save(&perbandinganMopa).Error

	return perbandinganMopa, err
}

func (r *users) DeletePerbandinganMopa(perbandinganMopa models.PerbandinganMopa) (models.PerbandinganMopa, error) {
	err := r.db.Delete(&perbandinganMopa).Error

	return perbandinganMopa, err
}
