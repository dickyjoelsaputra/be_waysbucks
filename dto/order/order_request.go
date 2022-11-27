package orderdto

type OrderRequest struct {
	UserID    int   `json:"user_id" gorm:"type: int"`
	ProductId int   `json:"product_id" form:"product_id"`
	TopingId  []int `json:"toping_id" form:"toping_id"`
	Quantity  int   `json:"quantity" gorm:"type: int"`
	Price     int   `json:"price" gorm:"type: int"`
}