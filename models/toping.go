package models

type Toping struct {
	ID     int          `json:"id"`
	Title  string       `json:"title" gorm:"type: varchar(255)"`
	Price  int          `json:"price" form:"price" gorm:"type: int"`
	Image  string       `json:"image" gorm:"type: varchar(255)"`
	UserID int          `json:"user_id"`
	User   UserResponse `json:"-"`
}

type TopingResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
	Image string `json:"image"`
}

func (TopingResponse) TableName() string {
	return "topings"
}