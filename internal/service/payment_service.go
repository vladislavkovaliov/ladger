package service

import (
	"context"

	"github.com/vladislavkovaliov/ledger/internal/domain/payment"
)

type PaymentService struct {
	repo payment.Repository
}

func NewPaymentService(repo payment.Repository) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) Create(ctx context.Context, p *payment.Payment) error {
	return s.repo.Save(ctx, p)
}

func (s *PaymentService) List(ctx context.Context) ([]*payment.Payment, error) {
	return s.repo.List(ctx)
}
