package usecase

import (
	"MyGram/internal/domain"
	"context"
	"time"
)

type socialMediaUseCase struct {
	contextTimeout        time.Duration
	socialMediaRepository domain.SocialMediaRepository
}

func NewSocialMediaUseCase(timeout time.Duration, ur domain.SocialMediaRepository) domain.SocialMediaUseCase {
	return &socialMediaUseCase{
		contextTimeout:        timeout,
		socialMediaRepository: ur,
	}
}

func (r socialMediaUseCase) GetSocialMedias(ctx context.Context, userId int) ([]domain.SocialMediaGetResponse, error) {
	var cList []domain.SocialMediaGetResponse
	c, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	result, err := r.socialMediaRepository.Fetch(c, userId)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		cList = append(cList, v.ToSocialMediaGetResponse())
	}

	return cList, nil
}

func (r socialMediaUseCase) GetSocialMediaById(ctx context.Context, id int) (*domain.SocialMediaResponse, error) {
	c, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	result, err := r.socialMediaRepository.FindByID(c, id)
	if err != nil {
		return nil, err
	}

	product := result.ToSocialMediaResponse()
	return &product, err
}

func (r socialMediaUseCase) SaveSocialMedia(ctx context.Context, body domain.SocialMediaRequest) (*domain.SocialMediaResponse, error) {
	c, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	data := body.ToSocialMedia()
	SocialMedia, err := r.socialMediaRepository.Store(c, data)
	result := SocialMedia.ToSocialMediaResponse()
	if err != nil {
		return &result, err
	}

	return &result, nil
}

func (r socialMediaUseCase) UpdateSocialMedia(ctx context.Context, body domain.SocialMediaRequest, id int) (*domain.SocialMediaUpdateResponse, error) {
	c, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	data := body.ToSocialMedia()
	data.Id = id

	SocialMedia, err := r.socialMediaRepository.Update(c, data)
	result := SocialMedia.ToSocialMediaUpdateResponse()
	if err != nil {
		return &result, err
	}

	return &result, nil
}

func (r socialMediaUseCase) DeleteSocialMedia(ctx context.Context, id int) error {
	c, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	err := r.socialMediaRepository.Delete(c, id)
	if err != nil {
		return err
	}
	return nil
}
