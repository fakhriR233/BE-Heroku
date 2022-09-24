package repositories

import (
	"dumbflix_be/models"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
  FindUsers() ([]models.User, error)
  GetUser(ID int) (models.User, error)
  CreateUser(user models.User) (models.User, error)
  UpdateUser(user models.User, ID int) (models.User, error)
  DeleteUser(user models.User, ID int) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
  return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
  var users []models.User
  err := r.db.Raw("SELECT * FROM users").Scan(&users).Error

  return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
  var user models.User
  err := r.db.Raw("SELECT * FROM users WHERE id=?", ID).Scan(&user).Error

  return user, err
}

 func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Exec("INSERT INTO users(full_name,email,password,gender,phone,address,subscribe,created_at,updated_at) VALUES (?,?,?,?,?,?,?,?,?)",user.FullName,user.Email, user.Password, user.Gender, user.Phone,user.Address,user.Subscribe,time.Now(), time.Now()).Error
  
	return user, err
  }

func (r *repository) UpdateUser(user models.User, ID int) (models.User, error) {
	err := r.db.Raw("UPDATE users SET full_name=?, email=?, password=?, gender=?, phone=?, address=? WHERE id=?", user.FullName, user.Email, user.Password,user.Gender, user.Phone,user.Address,ID).Scan(&user).Error
  
	return user, err
  }

 func (r *repository) DeleteUser(user models.User,ID int) (models.User, error) {
	err := r.db.Raw("DELETE FROM users WHERE id=?",ID).Scan(&user).Error
  
	return user, err
  }