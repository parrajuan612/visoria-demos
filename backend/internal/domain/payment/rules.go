package payment

type ScholarshipRule struct {
	Scholarship int
	Total       float64
}

var ScholarshipRules = []ScholarshipRule{
	{
		Scholarship: 0,
		Total:       2800,
	},
	{
		Scholarship: 30,
		Total:       1960,
	},
	{
		Scholarship: 50,
		Total:       1550,
	},
	{
		Scholarship: 70,
		Total:       990,
	},
	{
		Scholarship: 100,
		Total:       200,
	},
}
