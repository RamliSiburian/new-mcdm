package repositories

import (
	"fmt"
	models "mcdm/Models"
	"sort"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type AlternatifRepository interface {
	FindAlternatif() ([]models.Alternatif, error)
	GetLastKodeAlternatif() ([]int, error)
	CreateAlternatif(createAlternatif models.Alternatif) (models.Alternatif, error)
	GetAlternatifByKode(kode string) (models.Alternatif, error)
	UpdateAlternatif(alternatif models.Alternatif) (models.Alternatif, error)
	DeleteAlternatif(kode string) (models.Alternatif, error)
}

func RepositoryAlternatif(db *gorm.DB) *users {
	return &users{db}
}

func (r *users) FindAlternatif() ([]models.Alternatif, error) {
	var alternatif []models.Alternatif

	err := r.db.Order("CAST(kode AS INT) asc").Find(&alternatif).Error

	return alternatif, err
}

func (r *users) CreateAlternatif(createAlternatif models.Alternatif) (models.Alternatif, error) {
	err := r.db.Create(&createAlternatif).Error

	return createAlternatif, err
}

func (r *users) GetAlternatifByKode(kode string) (models.Alternatif, error) {
	var alternatif models.Alternatif

	err := r.db.Where("kode = ?", kode).First(&alternatif).Error

	return alternatif, err
}

func (r *users) UpdateAlternatif(alternatif models.Alternatif) (models.Alternatif, error) {
	err := r.db.Save(&alternatif).Error

	return alternatif, err
}

// func (r *users) DeleteAlternatif(alternatif models.Alternatif) (models.Alternatif, error) {
// 	err := r.db.Where("kode = ?", &alternatif.Kode).Delete(&alternatif).Error

// 	return alternatif, err
// }

func (r *users) DeleteAlternatif(kode string) (models.Alternatif, error) {
	var alternatif models.Alternatif

	err := r.db.Where("kode = ?", kode).Delete(&alternatif).Error

	return alternatif, err
}

func (r *users) GetLastKodeAlternatif() ([]int, error) {
	var lastKode []string
	err := r.db.Model(&models.Alternatif{}).Select("kode").
		Order("kode asc").
		Scan(&lastKode).Error
	if err != nil {
		return nil, err
	}

	if len(lastKode) == 0 {
		var missingNumbers []int

		missingNumbers = append(missingNumbers, 1)

		return missingNumbers, err
	} else {
		var numbers []int

		for _, item := range lastKode {
			numberStr := strings.TrimPrefix(item, "A")
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
			if !containsKode(numbers, i) {
				missingNumbers = append(missingNumbers, i)
				return missingNumbers, err
			}
		}
		missingNumbers = append(missingNumbers, lastNumber+1)

		return missingNumbers, err
	}
}

func containsKode(data []int, num int) bool {
	for _, val := range data {
		if val == num {
			return true
		}
	}
	return false
}
