package data

import (
	"github.com/go-kratos/kratos/v2/log"
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
