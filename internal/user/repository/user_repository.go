package repository

import (
	"MyGram/internal/domain"
	"context"
	"time"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (c userRepository) Fetch(ctx context.Context) ([]domain.User, error) {
	var entities []domain.User

	err := c.DB.Find(ctx).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (c userRepository) FindByID(ctx context.Context, id int) (*domain.User, error) {
	var entity domain.User
	err := c.DB.WithContext(ctx).First(&entity, "id =?", id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (c userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	var entity domain.User
	err := c.DB.WithContext(ctx).First(&entity, "email =?", email).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (c userRepository) Update(ctx context.Context, data domain.User) (*domain.User, error) {
	var entity domain.User
	entity.Username = data.Username
	entity.Email = data.Email
	c.DB.WithContext(ctx).First(&data)
	data.Email = entity.Email
	data.Username = entity.Username
	data.Updated_At = time.Now()

	err := c.DB.WithContext(ctx).Updates(&data).Error
	if err != nil {
		return &data, err
	}

	return &data, nil
}

func (c userRepository) Store(ctx context.Context, data domain.User) (*domain.User, error) {

	data.Created_At = time.Now()
	data.Updated_At = time.Now()
	err := c.DB.WithContext(ctx).Create(&data).Error
	if err != nil {
		return &data, err
	}
	return &data, nil
}

func (c userRepository) Delete(ctx context.Context, id int) error {
	err := c.DB.WithContext(ctx).Exec("delete from users where id =?", id).Error
	if err != nil {
		return err
	}
	return nil
}
