package service

import (
	"context"

	pb "review-b/api/business/v1"
	"review-b/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type BusinessService struct {
	pb.UnimplementedBusinessServer

	uc  *biz.ReviewReplyUsecase
	log *log.Helper
}

func NewBusinessService(uc *biz.ReviewReplyUsecase, logger log.Logger) *BusinessService {
	return &BusinessService{uc: uc, log: log.NewHelper(logger)}
}

func (s *BusinessService) ReplyReview(ctx context.Context, req *pb.ReplyReviewRequest) (*pb.ReplyReviewReply, error) {
	return &pb.ReplyReviewReply{}, nil
}
