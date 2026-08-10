package services

import "time"

type CategoryService struct{}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

func (s *CategoryService) GetCategory(birthDate time.Time) string {
	year := birthDate.Year()

	switch {
	case year >= 2008 && year <= 2010:
		return "JUVENIL"

	case year >= 2011 && year <= 2012:
		return "CADETE"

	case year >= 2013 && year <= 2014:
		return "INFANTIL"

	case year >= 2015 && year <= 2016:
		return "ALEVÍN"

	case year >= 2017 && year <= 2018:
		return "BENJAMÍN"

	default:
		return "NO DEFINIDA"
	}
}
