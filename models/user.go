package models

import (
	"gorm.io/gorm"
)

type User struct {
  gorm.Model
  ID        int                     `json:"id"`
  Name      string                  `json:"name" gorm:"type: varchar(255)"`
  Email     string                  `json:"email" gorm:"type: varchar(255)"`
  Password  string                  `json:"-" gorm:"type: varchar(255)"`
  Role      string                  `json:"role" gorm:"type: varchar(255)"`
  Image      string                 `json:"image" form:"image" gorm:"type: varchar(255)"`
}

type UserResponse struct {
  ID   int    `json:"id"`
  Name string `json:"name"`
  Email string  `json:"email"`
}

func (UserResponse) TableName() string {
  return "users"
}