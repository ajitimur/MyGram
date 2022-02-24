package usecase

import (
	"MyGram/internal/domain"
	"context"
	"time"
)

type userUseCase struct {
	contextTimeout time.Duration
	userRepository domain.UserRepository
}

func NewUserUseCase(timeout time.Duration, ur domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		contextTimeout: timeout,
		userRepository: ur,
	}
}

func (r userUseCase) GetUsers(ctx context.Context) ([]domain.UserResponse, error) {
	var cList []domain.UserResponse
	c, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	result, err := r.userRepository.Fetch(c)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		cList = append(cList, v.ToUserResponse())
	}

	return cList, nil
}

func (r userUseCase) GetUserById(ctx context.Context, id int) (*domain.UserResponse, error) {
	c, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	result, err := r.userRepository.FindByID(c, id)
	if err != nil {
		return nil, err
	}

	product := result.ToUserResponse()
	return &product, err
}

func (r userUseCase) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	c, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	result, err := r.userRepository.FindByEmail(c, email)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (r userUseCase) SaveUser(ctx context.Context, body domain.UserRequest) (*domain.UserResponse, error) {
	c, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	data := body.ToUser()
	user, err := r.userRepository.Store(c, data)
	result := user.ToUserResponse()
	if err != nil {
		return &result, err
	}

	return &result, nil
}

func (r userUseCase) UpdateUser(ctx context.Context, body domain.UserUpdateRequest, id int) (*domain.UserUpdateResponse, error) {
	c, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	data := body.ToUser()
	data.Id = id

	user, err := r.userRepository.Update(c, data)
	result := user.ToUserUpdateResponse()
	if err != nil {
		return &result, err
	}

	return &result, nil
}

func (r userUseCase) DeleteUser(ctx context.Context, id int) error {
	c, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	err := r.userRepository.Delete(c, id)
	if err != nil {
		return err
	}
	return nil
}
