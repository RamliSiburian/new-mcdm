package dto

type RegisterRequest struct {
	Username string `json:"username" validate:"required" gorm:"type: varchar(255)"`
	Password string `json:"password" validate:"required" gorm:"type: varchar(255)"`
	Role     string `json:"role" gorm:"type: varchar(255)"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required" gorm:"type: varchar(255)"`
	Password string `json:"password" validate:"required" gorm:"type: varchar(255)"`
}

type LoginResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username" gorm:"type: varchar(255)"`
	Role     string `json:"role" gorm:"type: varchar(255)"`
	Token    string `json:"token"`
}
type CheckAuthResponse struct {
	ID       int    `gorm:"type: int" json:"id"`
	Username string `json:"username"`
	Role     string `gorm:"type: varchar(255)" json:"role"`
	// Status   string `gorm:"type: varchar(50)"  json:"status"`
}
