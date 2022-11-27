package usersdto

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `gorm:"type: varchar(255)" json:"name"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	Image string `gorm:"type: varchar(255)" json:"image"`
}

type UserDeleteResponse struct {
	ID int `json:"id"`
}