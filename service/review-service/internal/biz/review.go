package biz

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"review-service/internal/data/model"
)

// GreeterRepo is a Greater repo.
type ReviewRepo interface {
	SaveReview(context.Context, *model.ReviewInfo) (*model.ReviewInfo, error)
	GetReviewByID(context.Context, int64) ([]*model.ReviewInfo, error)
}

// ReviewUsecase is a Review usecase.
type ReviewUsecase struct {
	repo ReviewRepo
	log  *log.Helper
}

// NewReviewUsecase new a Review usecase.
func NewReviewUsecase(repo ReviewRepo, logger log.Logger) *ReviewUsecase {
	return &ReviewUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

// CreateReview creates a Review, and returns the new Review.
func (uc *ReviewUsecase) CreateReview(ctx context.Context, r *model.ReviewInfo) (*model.ReviewInfo, error) {
	uc.log.WithContext(ctx).Infof("CreateReview: %+v", r)
	// 1. 数据校验
	reviews, err := uc.repo.GetReviewByID(ctx, r.OrderID)
	if err != nil {
		return nil, err
	}
	if len(reviews) > 0 {
		return nil, errors.New("订单已评价")
	}
	// 2. 生成 review ID
	r.ReviewID = time.Now().UnixNano()
	// 3. 查询订单和商品快照信息
	// 4. 拼装数据入库
	return uc.repo.SaveReview(ctx, r)
}
