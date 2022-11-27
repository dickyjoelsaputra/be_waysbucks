package models

import (
	"time"
)

type Product struct {
  ID        int       `json:"id"`
  Title      string   `json:"title" gorm:"type: varchar(255)"`
  Price      int      `json:"price" form:"price" gorm:"type: int"`
  Image      string   `json:"image" form:"image" gorm:"type: varchar(255)"`
  UserID      int     `json:"user_id"`
  User      UserResponse  `json:"-"`
  CreatedAt time.Time `json:"-"`
  UpdatedAt time.Time `json:"-"`
}

type ProductResponse struct {
  ID         int                  `json:"id"`
  Title      string              `json:"title"`
  Price      int                  `json:"price"`
  Image      string               `json:"image"`
}

func (ProductResponse) TableName() string {
  return "products"
}