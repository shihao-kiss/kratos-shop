package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type ReviewReplyRepo interface {
	CreateReviewReply(ctx context.Context) error
}

type ReviewReplyUsecase struct {
	repo ReviewReplyRepo
	log  *log.Helper
}

func NewReviewReplyUsecase(repo ReviewReplyRepo, logger log.Logger) *ReviewReplyUsecase {
	return &ReviewReplyUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *ReviewReplyUsecase) CreateReviewReply(ctx context.Context) error {
	uc.log.Infof("SaveReviewReplyAndUpdateReview: %+v", ctx)
	err := uc.repo.CreateReviewReply(ctx)
	if err != nil {
		uc.log.Errorf("CreateReviewReply err: %+v", err)
		return err
	}
	return nil
}
