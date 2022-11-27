package database

import (
	"be_waysbucks/models"
	"be_waysbucks/pkg/mysql"
	"fmt"
)

// Automatic Migration if Running App
func RunMigration() {
  err := mysql.DB.AutoMigrate(
		&models.User{},
		// &models.Profile{},
		&models.Product{},
		&models.Toping{},
		&models.Order{},
  )

  if err != nil {
    fmt.Println(err)
    panic("Migration Failed")
  }

  fmt.Println("Migration Success")
}