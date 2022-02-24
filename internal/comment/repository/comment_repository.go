package repository

import (
	"MyGram/internal/domain"
	"context"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type commentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) domain.CommentRepository {
	return &commentRepository{
		DB: db,
	}
}

func (c commentRepository) Fetch(ctx context.Context, UserId int) ([]domain.Comment, error) {
	// var entities []domain.CommentGetResponse
	// var user domain.User
	var Comment []domain.Comment

	err := c.DB.WithContext(ctx).Preload(clause.Associations).Find(&Comment, "user_id = ?", UserId).Error
	if err != nil {
		return nil, err
	}

	return Comment, nil
}

func (c commentRepository) FindByID(ctx context.Context, id int) (*domain.Comment, error) {
	var entity domain.Comment
	err := c.DB.WithContext(ctx).First(&entity, "id =?", id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (c commentRepository) Update(ctx context.Context, data domain.Comment) (*domain.Comment, error) {
	var entity domain.Comment
	entity.Message = data.Message

	c.DB.WithContext(ctx).First(&data)
	data.Message = entity.Message
	data.Updated_At = time.Now()

	err := c.DB.WithContext(ctx).Updates(&data).Error
	if err != nil {
		return &data, err
	}

	return &data, nil
}

func (c commentRepository) Store(ctx context.Context, data domain.Comment) (*domain.Comment, error) {

	data.Created_At = time.Now()
	data.Updated_At = time.Now()
	err := c.DB.WithContext(ctx).Create(&data).Error
	if err != nil {
		return &data, err
	}
	return &data, nil
}

func (c commentRepository) Delete(ctx context.Context, id int) error {
	err := c.DB.WithContext(ctx).Exec("delete from comments where id =?", id).Error
	if err != nil {
		return err
	}
	return nil
}
