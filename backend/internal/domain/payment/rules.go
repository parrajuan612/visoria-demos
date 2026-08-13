package payment

import "github.com/juanparra/visoria-demo/internal/domain"

type ScholarshipRule struct {
	Scholarship int
	Total       int
	Payment1    int
	Payment2    int
	Payment3    int
}

var ScholarshipRules = []ScholarshipRule{
	{
		Scholarship: 0,
		Total:       2800,
		Payment1:    930,
		Payment2:    930,
		Payment3:    940,
	},
	{
		Scholarship: 30,
		Total:       1960,
		Payment1:    650,
		Payment2:    650,
		Payment3:    660,
	},
	{
		Scholarship: 50,
		Total:       1550,
		Payment1:    500,
		Payment2:    500,
		Payment3:    550,
	},
	{
		Scholarship: 70,
		Total:       990,
		Payment1:    495,
		Payment2:    495,
		Payment3:    0,
	},
	{
		Scholarship: 100,
		Total:       200,
		Payment1:    200,
		Payment2:    0,
		Payment3:    0,
	},
}

func GetPlan(scholarship int) (domain.PaymentPlan, bool) {
	for _, rule := range ScholarshipRules {
		if rule.Scholarship == scholarship {
			return domain.PaymentPlan{
				Total:    rule.Total,
				Payment1: rule.Payment1,
				Payment2: rule.Payment2,
				Payment3: rule.Payment3,
			}, true
		}
	}

	return domain.PaymentPlan{}, false
}
