package usecase

import (
	"MyGram/internal/domain"
	"context"
	"time"
)

type photoUseCase struct {
	contextTimeout  time.Duration
	photoRepository domain.PhotoRepository
}

func NewPhotoUseCase(timeout time.Duration, ur domain.PhotoRepository) domain.PhotoUseCase {
	return &photoUseCase{
		contextTimeout:  timeout,
		photoRepository: ur,
	}
}

func (r photoUseCase) GetPhotos(ctx context.Context, userId int) ([]domain.PhotoGetResponse, error) {
	var cList []domain.PhotoGetResponse
	c, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	result, err := r.photoRepository.Fetch(c, userId)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		cList = append(cList, v.ToPhotoGetResponse())
	}

	return cList, nil
}

func (r photoUseCase) GetPhotoById(ctx context.Context, id int) (*domain.PhotoResponse, error) {
	c, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	result, err := r.photoRepository.FindByID(c, id)
	if err != nil {
		return nil, err
	}

	product := result.ToPhotoResponse()
	return &product, err
}

func (r photoUseCase) SavePhoto(ctx context.Context, body domain.PhotoRequest) (*domain.PhotoResponse, error) {
	c, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	data := body.ToPhoto()
	Photo, err := r.photoRepository.Store(c, data)
	result := Photo.ToPhotoResponse()
	if err != nil {
		return &result, err
	}

	return &result, nil
}

func (r photoUseCase) UpdatePhoto(ctx context.Context, body domain.PhotoUpdateRequest, id int) (*domain.PhotoUpdateResponse, error) {
	c, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	data := body.ToPhoto()
	data.Id = id

	Photo, err := r.photoRepository.Update(c, data)
	result := Photo.ToPhotoUpdateResponse()
	if err != nil {
		return &result, err
	}

	return &result, nil
}

func (r photoUseCase) DeletePhoto(ctx context.Context, id int) error {
	c, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	err := r.photoRepository.Delete(c, id)
	if err != nil {
		return err
	}
	return nil
}
