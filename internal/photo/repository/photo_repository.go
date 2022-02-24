package repository

import (
	"MyGram/internal/domain"
	"context"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type photoRepository struct {
	DB *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) domain.PhotoRepository {
	return &photoRepository{
		DB: db,
	}
}

func (c photoRepository) Fetch(ctx context.Context, UserId int) ([]domain.Photo, error) {
	// var entities []domain.PhotoGetResponse
	// var user domain.User
	var photo []domain.Photo

	err := c.DB.WithContext(ctx).Preload(clause.Associations).Find(&photo, "user_id = ?", UserId).Error
	if err != nil {
		return nil, err
	}

	return photo, nil
}

func (c photoRepository) FindByID(ctx context.Context, id int) (*domain.Photo, error) {
	var entity domain.Photo
	err := c.DB.WithContext(ctx).First(&entity, "id =?", id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (c photoRepository) Update(ctx context.Context, data domain.Photo) (*domain.Photo, error) {
	var entity domain.Photo
	entity.Title = data.Title
	entity.Caption = data.Caption
	entity.PhotoUrl = data.PhotoUrl

	c.DB.WithContext(ctx).First(&data)
	data.Title = entity.Title
	data.Caption = entity.Caption
	data.PhotoUrl = entity.PhotoUrl
	data.Updated_At = time.Now()
	err := c.DB.WithContext(ctx).Updates(&data).Error
	if err != nil {
		return &data, err
	}

	return &data, nil
}

func (c photoRepository) Store(ctx context.Context, data domain.Photo) (*domain.Photo, error) {

	data.Created_At = time.Now()
	data.Updated_At = time.Now()
	err := c.DB.WithContext(ctx).Create(&data).Error
	if err != nil {
		return &data, err
	}
	return &data, nil
}

func (c photoRepository) Delete(ctx context.Context, id int) error {
	err := c.DB.WithContext(ctx).Exec("delete from photos where id =?", id).Error
	if err != nil {
		return err
	}
	return nil
}
