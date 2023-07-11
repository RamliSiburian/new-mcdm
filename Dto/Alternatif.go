package dto

type AlternatifRequest struct {
	// Kode           string `json:"kode" form:"kode" gorm:"type: varchar(255)" validate:"required"`
	NamaAlternatif string `json:"namaAlternatif" form:"namaAlternatif" gorm:"type: varchar(255)" validate:"required"`
}

type AlternatifResponse struct {
	Kode           string `json:"kode"`
	NamaAlternatif string `json:"namaAlternatif"`
}
