package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// Greeter is a Greeter model.

// OperationRepo is a Greater repo.
type OperationRepo interface {
	AuditReview(context.Context, *AuditParam) error
	AuditAppeal(context.Context, *AuditAppealParam) error
}

// OperationUsecase is a Greeter usecase.
type OperationUsecase struct {
	repo OperationRepo
	log  *log.Helper
}

// NewOperationUsecase NewGreeterUsecase new a Greeter usecase.
func NewOperationUsecase(repo OperationRepo, logger log.Logger) *OperationUsecase {
	return &OperationUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (o *OperationUsecase) OpAuditReview(ctx context.Context, param *AuditParam) error {
	// operation audit user review
	o.log.WithContext(ctx).Infof("[biz] AuditReview(%+v)", param)
	return o.repo.AuditReview(ctx, param)

}

func (o *OperationUsecase) OpAuditAppeal(ctx context.Context, param *AuditAppealParam) error {
	// operation audit user review appeal
	o.log.WithContext(ctx).Infof("[biz] AuditAppeal(%+v)", param)
	return o.repo.AuditAppeal(ctx, param)
}
