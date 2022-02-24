package domain

import (
	"context"
	"time"
)

type Photo struct {
	Id         int       `gorm:"primarykey;autoIncrement:true;column:id"`
	Title      string    `gorm:"type:varchar(100);column:title"`
	Caption    string    `gorm:"type:varchar(100);column:caption"`
	PhotoUrl   string    `gorm:"type:varchar(100);column:photo_url"`
	UserId     int       `gorm:"column:user_id"`
	Created_At time.Time `gorm:"column:created_at"`
	Updated_At time.Time `gorm:"column:updated_at"`
	User       User      `gorm:"foreignKey:UserId"`
}

type PhotoGetResponse struct {
	Id         int       `gorm:"primarykey;autoIncrement:true;column:id"`
	Title      string    `gorm:"type:varchar(100);column:title"`
	Caption    string    `gorm:"type:varchar(100);column:caption"`
	PhotoUrl   string    `gorm:"type:varchar(100);column:photo_url"`
	UserId     int       `gorm:"column:user_id"`
	Created_At time.Time `gorm:"column:created_at"`
	Updated_At time.Time `gorm:"column:updated_at"`
	User       UserUpdateRequest
}

type PhotoResponse struct {
	Id         int       `gorm:"primarykey;column:id;autoIncrement:true"`
	Title      string    `json:"title"`
	Caption    string    `json:"caption"`
	PhotoUrl   string    `json:"photo_url"`
	UserId     int       `json:"user_id"`
	Created_At time.Time `json:"created_at"`
}

type PhotoRequest struct {
	Id       int    `json:"id,omitempty"`
	Title    string `json:"title" valid:"Required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" valid:"Required"`
	UserId   int    `json:"user_id"`
}

type PhotoUpdateRequest struct {
	Id       int    `json:"id,omitempty"`
	Title    string `json:"title" valid:"Required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" valid:"Required"`
}

type PhotoUpdateResponse struct {
	Id         int       `gorm:"primarykey;column:id;autoIncrement:true"`
	Title      string    `json:"title"`
	Caption    string    `json:"caption"`
	PhotoUrl   string    `json:"photo_url"`
	UserId     int       `json:"user_id"`
	Updated_At time.Time `json:"updated_at"`
}

func (r Photo) ToPhotoResponse() PhotoResponse {
	return PhotoResponse{
		Id:         r.Id,
		Title:      r.Title,
		Caption:    r.Caption,
		PhotoUrl:   r.PhotoUrl,
		UserId:     r.UserId,
		Created_At: r.Created_At,
	}
}

func (r Photo) ToPhotoGetResponse() PhotoGetResponse {
	return PhotoGetResponse{
		Id:         r.Id,
		Title:      r.Title,
		Caption:    r.Caption,
		PhotoUrl:   r.PhotoUrl,
		UserId:     r.UserId,
		Created_At: r.Created_At,
		Updated_At: r.Updated_At,
		User:       r.User.ToUserUpdateRequest(),
	}
}

func (r PhotoRequest) ToPhoto() Photo {
	return Photo{
		Id:       r.Id,
		Title:    r.Title,
		Caption:  r.Caption,
		PhotoUrl: r.PhotoUrl,
		UserId:   r.UserId,
	}
}

func (r PhotoUpdateRequest) ToPhoto() Photo {
	return Photo{
		Id:       r.Id,
		Title:    r.Title,
		Caption:  r.Caption,
		PhotoUrl: r.PhotoUrl,
	}
}

func (r Photo) ToPhotoUpdateResponse() PhotoUpdateResponse {
	return PhotoUpdateResponse{
		Id:         r.Id,
		Title:      r.Title,
		Caption:    r.Caption,
		PhotoUrl:   r.PhotoUrl,
		UserId:     r.UserId,
		Updated_At: r.Updated_At,
	}
}

func (r Photo) ToPhotoRequest() PhotoRequest {
	return PhotoRequest{
		Id:       r.Id,
		Title:    r.Title,
		Caption:  r.Caption,
		PhotoUrl: r.PhotoUrl,
		UserId:   r.UserId,
	}
}

type PhotoUseCase interface {
	GetPhotos(ctx context.Context, userId int) ([]PhotoGetResponse, error)
	GetPhotoById(ctx context.Context, id int) (*PhotoResponse, error)
	SavePhoto(ctx context.Context, body PhotoRequest) (*PhotoResponse, error)
	UpdatePhoto(ctx context.Context, body PhotoUpdateRequest, id int) (*PhotoUpdateResponse, error)
	DeletePhoto(ctx context.Context, id int) error
}

type PhotoRepository interface {
	Fetch(ctx context.Context, userId int) ([]Photo, error)
	FindByID(ctx context.Context, id int) (*Photo, error)
	Update(ctx context.Context, data Photo) (*Photo, error)
	Store(ctx context.Context, data Photo) (*Photo, error)
	Delete(ctx context.Context, id int) error
}
