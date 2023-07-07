package models

import "time"

type Kriteria struct {
	Kode         string    `json:"kode" gorm:"primary_key"`
	NamaKriteria string    `json:"namaKriteria" gorm:"type: varchar(255)"`
	Bobot        float64   `json:"bobot" `
	Kategory     string    `json:"kategory" gorm:"type: varchar(255)"`
	CreatedAt    time.Time `json:"createdAt"`
}
