package domain

import (
	"context"
	"time"
)

type Comment struct {
	Id         int       `gorm:"primarykey;autoIncrement:true;column:id"`
	UserId     int       `gorm:"column:user_id"`
	PhotoId    int       `gorm:"column:photo_id"`
	Message    string    `gorm:"type:varchar(100);column:message"`
	Created_At time.Time `gorm:"column:created_at"`
	Updated_At time.Time `gorm:"column:updated_at"`
	User       User      `gorm:"foreignKey:UserId"`
	Photo      Photo     `gorm:"foreignKey:PhotoId"`
}

type CommentResponse struct {
	Id         int       `gorm:"primarykey;column:id;autoIncrement:true"`
	UserId     int       `json:"user_id"`
	PhotoId    int       `json:"photo_id"`
	Message    string    `json:"message"`
	Created_At time.Time `json:"created_at"`
}

type CommentUpdateResponse struct {
	Id         int       `gorm:"primarykey;column:id;autoIncrement:true"`
	UserId     int       `json:"user_id"`
	PhotoId    int       `json:"photo_id"`
	Message    string    `json:"message"`
	Updated_At time.Time `json:"updated_at"`
}

type CommentRequest struct {
	Id      int    `json:"id,omitempty"`
	UserId  int    `json:"user_id"`
	PhotoId int    `json:"photo_id"`
	Message string `json:"message" valid:"Required"`
}

type CommentGetResponse struct {
	Id         int       `gorm:"primarykey;autoIncrement:true;column:id"`
	UserId     int       `gorm:"column:user_id"`
	PhotoId    int       `gorm:"column:photo_id"`
	Message    string    `gorm:"type:varchar(100);column:message"`
	Created_At time.Time `gorm:"column:created_at"`
	Updated_At time.Time `gorm:"column:updated_at"`
	User       UserUpdateRequest
	Photo      PhotoRequest
}

func (r Comment) ToCommentResponse() CommentResponse {
	return CommentResponse{
		Id:         r.Id,
		UserId:     r.UserId,
		PhotoId:    r.PhotoId,
		Message:    r.Message,
		Created_At: r.Created_At,
	}
}

func (r Comment) ToCommentUpdateResponse() CommentUpdateResponse {
	return CommentUpdateResponse{
		Id:         r.Id,
		UserId:     r.UserId,
		PhotoId:    r.PhotoId,
		Message:    r.Message,
		Updated_At: r.Updated_At,
	}
}

func (r CommentRequest) ToComment() Comment {
	return Comment{
		Id:      r.Id,
		UserId:  r.UserId,
		PhotoId: r.PhotoId,
		Message: r.Message,
	}
}

func (r Comment) ToCommentGetResponse() CommentGetResponse {
	return CommentGetResponse{
		Id:         r.Id,
		Message:    r.Message,
		PhotoId:    r.PhotoId,
		UserId:     r.UserId,
		Created_At: r.Created_At,
		Updated_At: r.Updated_At,
		User:       r.User.ToUserUpdateRequest(),
		Photo:      r.Photo.ToPhotoRequest(),
	}
}

type CommentUseCase interface {
	GetComments(ctx context.Context, userId int) ([]CommentGetResponse, error)
	GetCommentById(ctx context.Context, id int) (*CommentResponse, error)
	SaveComment(ctx context.Context, body CommentRequest) (*CommentResponse, error)
	UpdateComment(ctx context.Context, body CommentRequest, id int) (*CommentUpdateResponse, error)
	DeleteComment(ctx context.Context, id int) error
}

type CommentRepository interface {
	Fetch(ctx context.Context, userId int) ([]Comment, error)
	FindByID(ctx context.Context, id int) (*Comment, error)
	Update(ctx context.Context, data Comment) (*Comment, error)
	Store(ctx context.Context, data Comment) (*Comment, error)
	Delete(ctx context.Context, id int) error
}
