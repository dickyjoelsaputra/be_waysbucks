package repositories

import (
	"be_waysbucks/models"

	"gorm.io/gorm"
)

type TopingRepository interface {
  FindTopings() ([]models.Toping, error)
  GetToping(ID int) (models.Toping, error)
  CreateToping(toping models.Toping) (models.Toping, error)
  UpdateToping(toping models.Toping, ID int) (models.Toping, error)
  DeleteToping(toping models.Toping, ID int)(models.Toping, error)
}

func RepositoryToping(db *gorm.DB) *repository {
  return &repository{db}
}

func (r *repository) FindTopings() ([]models.Toping, error) {
  var topings []models.Toping
    err := r.db.Find(&topings).Error
  // err := r.db.Preload("User").Find(&topings).Error
//   err := r.db.Preload("User").Preload("Category").Find(&topings).Error

  return topings, err
}

func (r *repository) GetToping(ID int) (models.Toping, error) {
  var toping models.Toping
  // not yet using category relation, cause this step doesnt Belong to Many
  err := r.db.Preload("User").First(&toping, ID).Error
  // err := r.db.Preload("User").Preload("Category").First(&toping, ID).Error

  return toping, err
}


func (r *repository) CreateToping(toping models.Toping) (models.Toping, error) {
  err := r.db.Create(&toping).Error

  return toping, err
}


func (r *repository) UpdateToping(toping models.Toping, ID int) (models.Toping, error) {
  // err := r.db.Raw("UPDATE users SET name=?, email=?, password=? WHERE id=?", user.Name, user.Email, user.Password,ID).Scan(&user).Error
  err := r.db.Save(&toping).Error

  return toping, err
}

func(r *repository)DeleteToping(toping models.Toping, ID int)(models.Toping, error){
  err := r.db.Delete(&toping).Error
  return toping, err
}