package domain

import (
	"context"
	"time"
)

type SocialMedia struct {
	Id             int       `gorm:"primarykey;autoIncrement:true;column:id"`
	Name           string    `gorm:"type:varchar(20);column:name"`
	SocialMediaUrl string    `gorm:"type:varchar(100);column:social_media_url"`
	UserId         int       `gorm:"column:user_id"`
	Created_At     time.Time `gorm:"column:created_at"`
	Updated_At     time.Time `gorm:"column:updated_at"`
	User           User      `gorm:"foreignKey:UserId"`
}

type SocialMediaResponse struct {
	Id             int       `gorm:"primarykey;column:id;autoIncrement:true"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         int       `json:"user_id"`
	Created_At     time.Time `gorm:"column:created_at"`
}

type SocialMediaUpdateResponse struct {
	Id             int       `gorm:"primarykey;column:id;autoIncrement:true"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         int       `json:"user_id"`
	Updated_At     time.Time `gorm:"column:updated_at"`
}

type SocialMediaGetResponse struct {
	Id             int       `gorm:"primarykey;autoIncrement:true;column:id"`
	Name           string    `gorm:"type:varchar(20);column:name"`
	SocialMediaUrl string    `gorm:"type:varchar(100);column:social_media_url"`
	UserId         int       `gorm:"column:user_id"`
	Created_At     time.Time `gorm:"column:created_at"`
	Updated_At     time.Time `gorm:"column:updated_at"`
	User           UserUpdateRequest
}

type SocialMediaRequest struct {
	Id             int    `json:"id,omitempty"`
	Name           string `json:"name" valid:"Required"`
	SocialMediaUrl string `json:"social_media_url" valid:"Required"`
	UserId         int    `json:"user_id"`
}

func (r SocialMedia) ToSocialMediaResponse() SocialMediaResponse {
	return SocialMediaResponse{
		Id:             r.Id,
		Name:           r.Name,
		SocialMediaUrl: r.SocialMediaUrl,
		UserId:         r.UserId,
		Created_At:     r.Created_At,
	}
}

func (r SocialMedia) ToSocialMediaUpdateResponse() SocialMediaUpdateResponse {
	return SocialMediaUpdateResponse{
		Id:             r.Id,
		Name:           r.Name,
		SocialMediaUrl: r.SocialMediaUrl,
		UserId:         r.UserId,
		Updated_At:     r.Updated_At,
	}
}

func (r SocialMediaRequest) ToSocialMedia() SocialMedia {
	return SocialMedia{
		Id:             r.Id,
		Name:           r.Name,
		SocialMediaUrl: r.SocialMediaUrl,
		UserId:         r.UserId,
	}
}

func (r SocialMedia) ToSocialMediaGetResponse() SocialMediaGetResponse {
	return SocialMediaGetResponse{
		Id:             r.Id,
		Name:           r.Name,
		SocialMediaUrl: r.SocialMediaUrl,
		UserId:         r.UserId,
		Created_At:     r.Created_At,
		Updated_At:     r.Updated_At,
		User:           r.User.ToUserUpdateRequest(),
	}
}

type SocialMediaUseCase interface {
	GetSocialMedias(ctx context.Context, userId int) ([]SocialMediaGetResponse, error)
	GetSocialMediaById(ctx context.Context, id int) (*SocialMediaResponse, error)
	SaveSocialMedia(ctx context.Context, body SocialMediaRequest) (*SocialMediaResponse, error)
	UpdateSocialMedia(ctx context.Context, body SocialMediaRequest, id int) (*SocialMediaUpdateResponse, error)
	DeleteSocialMedia(ctx context.Context, id int) error
}

type SocialMediaRepository interface {
	Fetch(ctx context.Context, userId int) ([]SocialMedia, error)
	FindByID(ctx context.Context, id int) (*SocialMedia, error)
	Update(ctx context.Context, data SocialMedia) (*SocialMedia, error)
	Store(ctx context.Context, data SocialMedia) (*SocialMedia, error)
	Delete(ctx context.Context, id int) error
}
