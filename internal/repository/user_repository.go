package repository

import (
	"github.com/tonquangtu/data_server/internal/model"
	"gorm.io/gorm"
	"time"
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
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}
	user.UpdatedAt = user.CreatedAt
	userRepo.db.Create(user)
	return nil
}
