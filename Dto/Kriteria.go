package dto

type KriteriaRequest struct {
	// Kode         string  `json:"kode" form:"kode" gorm:"type: varchar(255)" validate:"required"`
	Bobot        float64 `json:"bobot" form:"bobot" gorm:"type: float" validate:"required"`
	NamaKriteria string  `json:"namaKriteria" form:"namaKriteria" gorm:"type: varchar(255)" validate:"required"`
	Kategori     string  `json:"kategori" form:"kategori" gorm:"type: varchar(255)" validate:"required"`
}

type KriteriaResponse struct {
	Kode         string  `json:"kode"`
	Bobot        float64 `json:"bobot"`
	NamaKriteria string  `json:"namaKriteria"`
	Kategori     string  `json:"kategori"`
}

type KriteriaUpdateRequest struct {
	Bobot        float64 `json:"bobot" form:"bobot" gorm:"type: float" `
	NamaKriteria string  `json:"namaKriteria" form:"namaKriteria" gorm:"type: varchar(255)" `
	Kategori     string  `json:"kategori" form:"kategori" gorm:"type: varchar(255)" `
}
