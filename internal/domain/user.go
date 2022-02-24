package domain

import (
	"context"
	"time"
)

type User struct {
	Id         int       `gorm:"primarykey;autoIncrement:true;"`
	Username   string    `gorm:"type:varchar(20);column:username;unique"`
	Email      string    `gorm:"type:varchar(100);column:email;unique"`
	Password   string    `gorm:"type:varchar(100);column:password"`
	Age        int       `gorm:"column:age"`
	Created_At time.Time `gorm:"column:created_at"`
	Updated_At time.Time `gorm:"column:updated_at"`
}

type UserResponse struct {
	Id       int    `gorm:"primarykey;column:id;autoIncrement:true"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type UserUpdateResponse struct {
	Id         int       `gorm:"primarykey;column:id;autoIncrement:true"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Age        int       `json:"age"`
	Updated_At time.Time `gorm:"column:updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" valid:"Required; Email"`
	Password string `json:"password" valid:"Required; MinSize(6)"`
}

type UserRequest struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username" valid:"Required"`
	Email    string `json:"email" valid:"Required; Email"`
	Password string `json:"password" valid:"Required; MinSize(6)"`
	Age      int    `json:"age" valid:"Required; Min(9)"`
}

type UserUpdateRequest struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username" valid:"Required"`
	Email    string `json:"email" valid:"Required; Email"`
}

func (r User) ToUserResponse() UserResponse {
	return UserResponse{
		Id:       r.Id,
		Username: r.Username,
		Email:    r.Email,
		Age:      r.Age,
	}
}

func (r User) ToUserUpdateResponse() UserUpdateResponse {
	return UserUpdateResponse{
		Id:         r.Id,
		Username:   r.Username,
		Email:      r.Email,
		Age:        r.Age,
		Updated_At: r.Updated_At,
	}
}

func (r User) ToUserUpdateRequest() UserUpdateRequest {
	return UserUpdateRequest{
		Id:       r.Id,
		Username: r.Username,
		Email:    r.Email,
	}
}

func (r UserRequest) ToUser() User {
	return User{
		Id:       r.Id,
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
		Age:      r.Age,
	}
}

func (r UserUpdateRequest) ToUser() User {
	return User{
		Id:       r.Id,
		Username: r.Username,
		Email:    r.Email,
	}
}

func (r UserRequest) ToUserResponse() UserResponse {
	return UserResponse{
		Id:       r.Id,
		Username: r.Username,
		Email:    r.Email,
		Age:      r.Age,
	}
}

type UserUseCase interface {
	GetUsers(ctx context.Context) ([]UserResponse, error)
	GetUserById(ctx context.Context, id int) (*UserResponse, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	SaveUser(ctx context.Context, body UserRequest) (*UserResponse, error)
	UpdateUser(ctx context.Context, body UserUpdateRequest, id int) (*UserUpdateResponse, error)
	DeleteUser(ctx context.Context, id int) error
}

type UserRepository interface {
	Fetch(ctx context.Context) ([]User, error)
	FindByID(ctx context.Context, id int) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, data User) (*User, error)
	Store(ctx context.Context, data User) (*User, error)
	Delete(ctx context.Context, id int) error
}
