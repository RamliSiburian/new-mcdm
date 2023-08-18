package dto

type PerbandinganAhpRequest struct {
	Kode      int     `json:"kode" form:"kode" gorm:"type: varchar(255)" validate:"required"`
	Deskripsi string  `json:"deskripsi" form:"deskripsi" gorm:"type: varchar(255)" validate:"required"`
	Nilai     float64 `json:"nilai" form:"nilai" gorm:"type: float" validate:"required"`
}

type PerbandinganAhpResponse struct {
	Kode      int     `json:"kode"`
	Deskripsi string  `json:"deskripsi"`
	Nilai     float64 `json:"nilai"`
}

type PerbandinganAhpUpdateRequest struct {
	Deskripsi string  `json:"deskripsi" form:"deskripsi" gorm:"type: varchar(255)" `
	Nilai     float64 `json:"nilai" form:"nilai" gorm:"type: float" `
}
