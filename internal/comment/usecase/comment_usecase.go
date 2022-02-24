package usecase

import (
	"MyGram/internal/domain"
	"context"
	"time"
)

type commentUseCase struct {
	contextTimeout    time.Duration
	commentRepository domain.CommentRepository
}

func NewCommentUseCase(timeout time.Duration, ur domain.CommentRepository) domain.CommentUseCase {
	return &commentUseCase{
		contextTimeout:    timeout,
		commentRepository: ur,
	}
}

func (r commentUseCase) GetComments(ctx context.Context, userId int) ([]domain.CommentGetResponse, error) {
	var cList []domain.CommentGetResponse
	c, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	result, err := r.commentRepository.Fetch(c, userId)
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		cList = append(cList, v.ToCommentGetResponse())
	}

	return cList, nil
}

func (r commentUseCase) GetCommentById(ctx context.Context, id int) (*domain.CommentResponse, error) {
	c, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()
	result, err := r.commentRepository.FindByID(c, id)
	if err != nil {
		return nil, err
	}

	product := result.ToCommentResponse()
	return &product, err
}

func (r commentUseCase) SaveComment(ctx context.Context, body domain.CommentRequest) (*domain.CommentResponse, error) {
	c, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	data := body.ToComment()
	comment, err := r.commentRepository.Store(c, data)
	result := comment.ToCommentResponse()
	if err != nil {
		return &result, err
	}

	return &result, nil
}

func (r commentUseCase) UpdateComment(ctx context.Context, body domain.CommentRequest, id int) (*domain.CommentUpdateResponse, error) {
	c, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	data := body.ToComment()
	data.Id = id

	comment, err := r.commentRepository.Update(c, data)
	result := comment.ToCommentUpdateResponse()
	if err != nil {
		return &result, err
	}

	return &result, nil
}

func (r commentUseCase) DeleteComment(ctx context.Context, id int) error {
	c, cancel := context.WithTimeout(ctx, r.contextTimeout)
	defer cancel()

	err := r.commentRepository.Delete(c, id)
	if err != nil {
		return err
	}
	return nil
}
