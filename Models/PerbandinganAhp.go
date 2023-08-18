package models

import "time"

type PerbandinganAhp struct {
	Kode      int       `json:"kode" gorm:"primary_key"`
	Deskripsi string    `json:"deskripsi" gorm:"type: varchar(255)"`
	Nilai     float64   `json:"nilai" `
	CreatedAt time.Time `json:"createdAt"`
}
