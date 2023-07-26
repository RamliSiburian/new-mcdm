package dto

type AlternatifRequest struct {
	Kode           string `json:"kode" form:"kode" gorm:"type: varchar(255)" validate:"required"`
	KodeKriteria   string `json:"kodeKriteria" form:"kodeKriteria" gorm:"type: varchar(255)" validate:"required"`
	NamaAlternatif string `json:"namaAlternatif" form:"namaAlternatif" gorm:"type: varchar(255)" validate:"required"`
	Nilai          string `json:"nilai" form:"nilai" gorm:"type: varchar(255)" validate:"required"`
}

type AlternatifResponse struct {
	Kode           string `json:"kode"`
	KodeKriteria   string `json:"kodeKriteria"`
	NamaAlternatif string `json:"namaAlternatif"`
	Nilai          string `json:"nilai"`
}
