package service

import (
	"github.com/lllllan02/iam/internal/data"
	"github.com/lllllan02/iam/pkg/log"
)

type Service struct {
	logger *log.Logger
	tm     data.Transaction
}

func NewService(logger *log.Logger, tm data.Transaction) *Service {
	return &Service{
		logger: logger,
		tm:     tm,
	}
}
