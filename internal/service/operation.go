package service

import (
	"context"
	"review-o/internal/biz"

	pb "review-o/api/operation/v1"
)

type OperationService struct {
	pb.UnimplementedOperationServer
	uc *biz.OperationUsecase
}

func NewOperationService(uc *biz.OperationUsecase) *OperationService {
	return &OperationService{uc: uc}
}

func (s *OperationService) AuditReview(ctx context.Context, req *pb.AuditReviewRequest) (*pb.AuditReviewReply, error) {
	err := s.uc.OpAuditReview(ctx, &biz.AuditParam{
		ReviewID: req.GetReviewID(),
		Status:   req.GetStatus(),
		OpUser:   req.GetOpUser(),
		OpReason: req.GetOpReason(),
		OpRemark: req.GetOpRemarks(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.AuditReviewReply{}, nil
}

func (s *OperationService) AuditAppeal(ctx context.Context, req *pb.AuditAppealRequest) (*pb.AuditAppealReply, error) {
	err := s.uc.OpAuditAppeal(ctx, &biz.AuditAppealParam{
		AppealID: req.GetAppealID(),
		ReviewID: req.GetReviewID(),
		Status:   req.GetStatus(),
		OpUser:   req.GetOpUser(),
		OpReason: req.GetOpReason(),
		OpRemark: req.GetOpRemarks(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.AuditAppealReply{}, nil
}
