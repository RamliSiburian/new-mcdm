package repositories

import (
	"fmt"
	models "mcdm/Models"
	"sort"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type users struct {
	db *gorm.DB
}

type KriteriaRepository interface {
	FindKriteria() ([]models.Kriteria, error)
	CreateKriteria(createKriteria models.Kriteria) (models.Kriteria, error)
	GetKriteriaByKode(kode string) (models.Kriteria, error)
	GetLastKode() ([]int, error)
	UpdateKriteria(kriteria models.Kriteria) (models.Kriteria, error)
	DeleteKriteria(kriteria models.Kriteria) (models.Kriteria, error)
}

func RepositoryKriteria(db *gorm.DB) *users {
	return &users{db}
}

func (r *users) FindKriteria() ([]models.Kriteria, error) {
	var kriteria []models.Kriteria

	err := r.db.Order("CAST(kode AS INT) asc").Find(&kriteria).Error

	return kriteria, err
}

func (r *users) CreateKriteria(createKriteria models.Kriteria) (models.Kriteria, error) {
	err := r.db.Create(&createKriteria).Error

	return createKriteria, err
}

func (r *users) GetKriteriaByKode(kode string) (models.Kriteria, error) {
	var kriteria models.Kriteria

	err := r.db.Where("kode = ?", kode).First(&kriteria).Error

	return kriteria, err
}

func (r *users) GetLastKode() ([]int, error) {
	var lastKode []string
	err := r.db.Model(&models.Kriteria{}).Select("kode").
		Order("kode asc").
		Scan(&lastKode).Error
	if err != nil {
		return nil, err
	}

	var numbers []int

	for _, item := range lastKode {
		numberStr := strings.TrimPrefix(item, "C")
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			fmt.Println(err)
			continue
		}
		numbers = append(numbers, number)
	}

	if len(numbers) == 0 {
		return nil, err
	}

	sort.Ints(numbers)

	var missingNumbers []int
	lastNumber := numbers[len(numbers)-1]

	for i := 1; i < lastNumber; i++ {
		if !contains(numbers, i) {
			missingNumbers = append(missingNumbers, i)
			return missingNumbers, err
		}
	}
	missingNumbers = append(missingNumbers, lastNumber+1)

	return missingNumbers, err
}

func contains(data []int, num int) bool {
	for _, val := range data {
		if val == num {
			return true
		}
	}
	return false
}

func (r *users) UpdateKriteria(kriteria models.Kriteria) (models.Kriteria, error) {
	err := r.db.Save(&kriteria).Error

	return kriteria, err
}

func (r *users) DeleteKriteria(kriteria models.Kriteria) (models.Kriteria, error) {
	err := r.db.Delete(&kriteria).Error

	return kriteria, err
}
