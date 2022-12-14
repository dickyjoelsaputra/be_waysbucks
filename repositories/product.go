package repositories

import (
	"be_waysbucks/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
  FindProducts() ([]models.Product, error)
  GetProduct(ID int) (models.Product, error)
  CreateProduct(product models.Product) (models.Product, error)
  UpdateProduct(product models.Product, ID int) (models.Product, error)
  DeleteProduct(product models.Product, ID int)(models.Product, error)
}

func RepositoryProduct(db *gorm.DB) *repository {
  return &repository{db}
}

func (r *repository) FindProducts() ([]models.Product, error) {
  var products []models.Product
    err := r.db.Find(&products).Error
  // err := r.db.Preload("User").Find(&products).Error
//   err := r.db.Preload("User").Preload("Category").Find(&products).Error

  return products, err
}

func (r *repository) GetProduct(ID int) (models.Product, error) {
  var product models.Product
  // not yet using category relation, cause this step doesnt Belong to Many
  err := r.db.Preload("User").First(&product, ID).Error
  // err := r.db.Preload("User").Preload("Category").First(&product, ID).Error

  return product, err
}


func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
  err := r.db.Create(&product).Error

  return product, err
}


func (r *repository) UpdateProduct(product models.Product, ID int) (models.Product, error) {
  // err := r.db.Raw("UPDATE users SET name=?, email=?, password=? WHERE id=?", user.Name, user.Email, user.Password,ID).Scan(&user).Error
  err := r.db.Save(&product).Error

  return product, err
}

func(r *repository)DeleteProduct(product models.Product, ID int)(models.Product, error){
  err := r.db.Delete(&product).Error
  return product, err
}