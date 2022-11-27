package models

type Order struct {
	ID         int `json:"id" gorm:"primary_key:auto_increment"`
	Qty        int `json:"qty" gorm:"type:int"`
	TotalPrice int `json:"price" gorm:"type:int"`

	UserID int          `json:"user_id"`
	User   UserResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	ProductID int     `json:"product_id" form:"product_id" gorm:"foreignKey:ID"`
	Product   Product `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	TopingID []int    `json:"toping_id" gorm:"-"`
	Toping   []Toping `json:"toping" gorm:"many2many:toping_product;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type OrderResponse struct {
	ID         int
	UserID     int          `json:"user_id"`
	User       UserResponse `json:"user"`
	ProductID  int          `json:"product_id"`
	TopingID   []int        `json:"toping_id"`
	Qty        int          `json:"qty"`
	TotalPrice int          `json:"price"`
}

func (OrderResponse) TableName() string {
	return "order"
}