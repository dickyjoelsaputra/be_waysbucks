package repositories

import (
	"be_waysbucks/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
  CreateOrder(order models.Order) (models.Order, error)
  GetToping(ID int) (models.Toping, error)
  GetOrder(ID int) (models.Order, error)
  GetTopingOrder(ID []int)([]models.Toping, error)
  GetProductOrder(ID int) (models.Product, error)
  DeleteOrder(order models.Order, ID int)(models.Order, error)
  GetOrderByUserID(ID int) ([]models.Order, error)
  GetUserOrder(ID int) (models.User, error)
}

func RepositoryOrder(db *gorm.DB) *repository {
  return &repository{db}
}

func (r *repository) CreateOrder(order models.Order) (models.Order, error) {
  err := r.db.Create(&order).Error

  return order, err
}



func (r *repository) GetUserOrder(ID int) (models.User, error) {
  var user models.User
  // err := r.db.Raw("SELECT * FROM users WHERE id=?", ID).Scan(&user).Error
  err := r.db.First(&user, ID).Error
  // err := r.db.Preload("Profile").First(&user, ID).Error
  // err := r.db.Preload("Profile").Preload("Products").First(&user, ID).Error

  return user, err
}

func (r *repository) GetOrder(ID int) (models.Order, error) {
  var order models.Order
  // not yet using category relation, cause this step doesnt Belong to Many
  err := r.db.Preload("Product").Preload("User").Preload("Toping").First(&order, ID).Error
  // err := r.db.Preload("User").Preload("Category").First(&product, ID).Error

  return order, err
}

func(r *repository)GetTopingOrder(ID []int)([]models.Toping, error){
  // fmt.Println(ID)
  var toping []models.Toping
  // not yet using category relation, cause this step doesnt Belong to Many
  err := r.db.Preload("User").Find(&toping, ID).Error
  // err := r.db.Preload("User").Preload("Category").First(&toping, ID).Error

  return toping, err
}

func (r *repository) GetProductOrder(ID int) (models.Product, error) {
  var product models.Product
  // not yet using category relation, cause this step doesnt Belong to Many
  err := r.db.Preload("User").First(&product, ID).Error
  // err := r.db.Preload("User").Preload("Category").First(&product, ID).Error

  return product, err
}

// DELETE ORDER
 func (r *repository) DeleteOrder(order models.Order, ID int)(models.Order, error){
  err := r.db.Delete(&order).Error
  return order, err
}


// MY TRANSACTION

// func (r *repository) MyTransactions(ID int) (models.Order, error) {
//   var order models.Order
//   // not yet using category relation, cause this step doesnt Belong to Many
//   err := r.db.Preload("User").First(&order, ID).Error
//   // err := r.db.Preload("User").Preload("Category").First(&product, ID).Error

//   return order, err
// }

func (r *repository) GetOrderByUserID(ID int) ([]models.Order, error) {
	var order []models.Order
	err := r.db.Preload("Product").Preload("Toping").Preload("User").Where("user_id = ?", ID).Find(&order).Error
	return order, err
}
