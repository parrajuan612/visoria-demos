package tournament

import "time"

type Category struct {
	Name  string
	Years []int
}

var Categories = []Category{
	{
		Name:  "JUVENIL",
		Years: []int{2008, 2009, 2010},
	},
	{
		Name:  "CADETE",
		Years: []int{2011, 2012},
	},
	{
		Name:  "INFANTIL",
		Years: []int{2013, 2014},
	},
	{
		Name:  "ALEVÍN",
		Years: []int{2015, 2016},
	},
	{
		Name:  "BENJAMÍN",
		Years: []int{2017, 2018},
	},
}

func GetCategory(birthDate time.Time) string {
	year := birthDate.Year()

	for _, category := range Categories {
		for _, categoryYear := range category.Years {
			if year == categoryYear {
				return category.Name
			}
		}
	}

	return "SIN CATEGORÍA"
}
