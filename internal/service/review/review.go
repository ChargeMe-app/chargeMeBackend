package review

import (
	"context"
	reviewDomain "github.com/poorfrombabylon/chargeMeBackend/internal/domain/review"
)

type ReviewService interface {
	CreateReview(context.Context, reviewDomain.Review) error
}

type ReviewStorage interface {
	CreateReview(context.Context, reviewDomain.Review) error
}

type service struct {
	reviewStorage ReviewStorage
}

func NewReviewService(reviewStorage ReviewStorage) ReviewService {
	return &service{
		reviewStorage: reviewStorage,
	}
}

func (s *service) CreateReview(ctx context.Context, review reviewDomain.Review) error {
	return s.CreateReview(ctx, review)
}
