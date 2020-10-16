package repository

import (
	"data_server/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (userRepo *UserRepository) GetUser(id int) (*model.User, error) {
	var user model.User
	userRepo.db.Raw("SELECT * FROM public.user WHERE id=?", id).Scan(&user)
	return &user, nil
}

func (userRepo *UserRepository) AddUser(user *model.User) error {
	userRepo.db.Create(user)
	return nil
}
