package authdto

// type AuthRequest struct {
// 	Name     string `gorm:"type: varchar(255)" json:"name"`
// 	Email    string `gorm:"type: varchar(255)" json:"email"`
// 	Password string `gorm:"type: varchar(255)" json:"password"`
// }

type RegisterRequest struct {
	Name     string `gorm:"type: varchar(255)" json:"name" validate:"required"`
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
	Image    string `json:"image" form:"image" gorm:"type: varchar(255)"`
}

type LoginRequest struct {
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}