package payment

import (
	"fmt"

	domain "github.com/juanparra/visoria-demo/internal/domain"
	domainPayment "github.com/juanparra/visoria-demo/internal/domain/payment"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetTotal(scholarship int) (int, error) {
	for _, rule := range domainPayment.ScholarshipRules {
		if rule.Scholarship == scholarship {
			return rule.Total, nil
		}
	}

	return 0, fmt.Errorf("beca no configurada: %d%%", scholarship)
}

func (s *Service) GetPlan(scholarship int) (domain.PaymentPlan, error) {
	for _, rule := range domainPayment.ScholarshipRules {
		if rule.Scholarship == scholarship {
			return domain.PaymentPlan{
				Total:    rule.Total,
				Payment1: rule.Payment1,
				Payment2: rule.Payment2,
				Payment3: rule.Payment3,
			}, nil
		}
	}

	return domain.PaymentPlan{}, fmt.Errorf(
		"beca no configurada: %d%%",
		scholarship,
	)
}
