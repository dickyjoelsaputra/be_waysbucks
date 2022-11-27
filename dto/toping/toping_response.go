package topingdto

type TopingResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title" form:"title" validate:"required"`
	Price int    `json:"price" form:"price" validate:"required"`
	Image string `json:"image"`
}

type TopingResponseDelete struct {
	ID int `json:"id"`
}