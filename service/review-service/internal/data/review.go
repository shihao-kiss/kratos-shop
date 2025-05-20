package data

import (
	"context"
	"review-service/internal/biz"
	"review-service/internal/data/model"

	"github.com/go-kratos/kratos/v2/log"
)

type reviewRepo struct {
	data *Data
	log  *log.Helper
}

func NewReviewRepo(data *Data, logger log.Logger) biz.ReviewRepo {
	return &reviewRepo{
		data: data,
		log: log.NewHelper(logger),
	}
}

func (r *reviewRepo) SaveReview(ctx context.Context, review *model.ReviewInfo) (*model.ReviewInfo, error) {
	r.log.Infof("SaveReview req %+v", review)
	err := r.data.query.ReviewInfo.WithContext(ctx).Save(review)
	if err != nil {
		r.log.Errorf("SaveReview err %+v", err)
		return nil, err
	}
	return review, err
}