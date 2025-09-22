package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "review-o/api/review/v1"
	"review-o/internal/biz"
)

type operationRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.OperationRepo {
	return &operationRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// AuditReview 审核评价
func (r *operationRepo) AuditReview(ctx context.Context, param *biz.AuditParam) error {
	r.log.WithContext(ctx).Infof("[data] AuditReview param(%+v)", param)
	// 之前都是写数据库
	// 现在需要通过 RPC 调用 review-service 服务的接口
	_, err := r.data.rc.AuditReview(ctx, &v1.AuditReviewRequest{
		ReviewID:  param.ReviewID,
		Status:    param.Status,
		OpUser:    param.OpUser,
		OpReason:  param.OpReason,
		OpRemarks: &param.OpRemark,
	})
	if err != nil {
		r.log.WithContext(ctx).Errorf("[data] AuditReview call review-service AuditReview error(%v)", err)
		return err
	}
	return nil
}

// AuditAppeal 审核申诉
func (r *operationRepo) AuditAppeal(ctx context.Context, param *biz.AuditAppealParam) error {
	r.log.WithContext(ctx).Infof("[data] AuditAppeal param(%+v)", param)
	// 之前都是写数据库
	// 现在需要通过 RPC 调用 review-service 服务的接口
	_, err := r.data.rc.AuditAppeal(ctx, &v1.AuditAppealRequest{
		AppealID:  param.AppealID,
		ReviewID:  param.ReviewID,
		Status:    param.Status,
		OpReason:  param.OpReason,
		OpRemarks: &param.OpRemark,
		OpUser:    param.OpUser,
	})
	if err != nil {
		r.log.WithContext(ctx).Errorf("[data] AuditAppeal call review-service AuditAppeal error(%v)", err)
		return err
	}
	return nil
}
