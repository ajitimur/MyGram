package repository

import (
	"MyGram/internal/domain"
	"context"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type socialMediaRepository struct {
	DB *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) domain.SocialMediaRepository {
	return &socialMediaRepository{
		DB: db,
	}
}

func (c socialMediaRepository) Fetch(ctx context.Context, UserId int) ([]domain.SocialMedia, error) {
	// var entities []domain.SocialMediaGetResponse
	// var user domain.User
	var SocialMedia []domain.SocialMedia

	err := c.DB.WithContext(ctx).Preload(clause.Associations).Find(&SocialMedia, "user_id = ?", UserId).Error
	if err != nil {
		return nil, err
	}

	return SocialMedia, nil
}

func (c socialMediaRepository) FindByID(ctx context.Context, id int) (*domain.SocialMedia, error) {
	var entity domain.SocialMedia
	err := c.DB.WithContext(ctx).First(&entity, "id =?", id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (c socialMediaRepository) Update(ctx context.Context, data domain.SocialMedia) (*domain.SocialMedia, error) {
	var entity domain.SocialMedia

	entity.Name = data.Name
	entity.SocialMediaUrl = data.SocialMediaUrl

	c.DB.WithContext(ctx).First(&data)

	data.Name = entity.Name
	data.SocialMediaUrl = entity.SocialMediaUrl
	data.Updated_At = time.Now()

	err := c.DB.WithContext(ctx).Updates(&data).Error
	if err != nil {
		return &data, err
	}

	return &data, nil
}

func (c socialMediaRepository) Store(ctx context.Context, data domain.SocialMedia) (*domain.SocialMedia, error) {

	data.Created_At = time.Now()
	data.Updated_At = time.Now()
	err := c.DB.WithContext(ctx).Create(&data).Error
	if err != nil {
		return &data, err
	}
	return &data, nil
}

func (c socialMediaRepository) Delete(ctx context.Context, id int) error {
	err := c.DB.WithContext(ctx).Exec("delete from social_media where id =?", id).Error
	if err != nil {
		return err
	}
	return nil
}
