package productdto

type ProductResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title" form:"title" validate:"required"`
	Price int    `json:"price" form:"price" validate:"required"`
	Image string `json:"image"`
}

type ProductResponseDelete struct {
	ID int `json:"id"`
}