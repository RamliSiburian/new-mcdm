package dto

type PerbandinganMopaRequest struct {
	Kode      int     `json:"kode" form:"kode" `
	Deskripsi string  `json:"deskripsi" form:"deskripsi" gorm:"type: varchar(255)" validate:"required"`
	Nilai     float64 `json:"nilai" form:"nilai" gorm:"type: float" `
}

type PerbandinganMopaResponse struct {
	Kode      int     `json:"kode"`
	Deskripsi string  `json:"deskripsi"`
	Nilai     float64 `json:"nilai"`
}

type PerbandinganMopaUpdateRequest struct {
	Deskripsi string  `json:"deskripsi" form:"deskripsi" gorm:"type: varchar(255)" `
	Nilai     float64 `json:"nilai" form:"nilai" gorm:"type: float" `
}
