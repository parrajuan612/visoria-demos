package payment

import (
	"fmt"

	domainPayment "github.com/juanparra/visoria-demo/internal/domain/payment"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetTotal(scholarship int) (float64, error) {
	// Se usa domainPayment.ScholarshipRules para acceder a la variable exportada
	for _, rule := range domainPayment.ScholarshipRules {
		if rule.Scholarship == scholarship {
			return rule.Total, nil
		}
	}

	return 0, fmt.Errorf("beca no configurada: %d%%", scholarship)
}
