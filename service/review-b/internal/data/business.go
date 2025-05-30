package data

import (
	"context"

	"review-b/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type reviewReplyRepo struct {
	data *Data
	log  *log.Helper
}

func NewReviewReplyRepo(data *Data, logger log.Logger) biz.ReviewReplyRepo {
	return &reviewReplyRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *reviewReplyRepo) CreateReviewReply(ctx context.Context) error {
	r.log.Infof("CreateReviewReply: %+v", ctx)
	return nil
}
