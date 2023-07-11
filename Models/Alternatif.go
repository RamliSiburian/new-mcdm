package models

import "time"

type Alternatif struct {
	Kode           string    `json:"kode" gorm:"primary_key"`
	NamaAlternatif string    `json:"namaAlternatif" gorm:"type: varchar(255)"`
	CreatedAt      time.Time `json:"createdAt"`
}
