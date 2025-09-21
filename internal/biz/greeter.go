package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

// Greeter is a Greeter model.

// OperationRepo is a Greater repo.
type OperationRepo interface {
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
