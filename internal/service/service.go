package service

import (
	"github.com/lllllan02/iam/internal/repository"
	"github.com/lllllan02/iam/pkg/log"
)

type Service struct {
	logger *log.Logger
	tm     repository.Transaction
}

func NewService(logger *log.Logger, tm repository.Transaction) *Service {
	return &Service{
		logger: logger,
		tm:     tm,
	}
}
