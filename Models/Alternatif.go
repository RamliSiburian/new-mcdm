package models

import "time"

type Alternatif struct {
	ID             int       `json:"id" gorm:"primary_key:auto_increment"`
	Kode           string    `json:"kode" gorm:"type: varchar(255)"`
	KodeKriteria   string    `json:"kodeKriteria" gorm:"type: varchar(255)"`
	NamaAlternatif string    `json:"namaAlternatif" gorm:"type: varchar(255)"`
	Nilai          string    `json:"nilai" gorm:"type: varchar(255)"`
	CreatedAt      time.Time `json:"createdAt"`
	Kriteria       Kriteria  `json:"kriteria" gorm:"foreignKey:KodeKriteria;references:Kode"`
}
