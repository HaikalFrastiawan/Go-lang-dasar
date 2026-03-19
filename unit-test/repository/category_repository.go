package repository

import "go-unit-test/entity"

type CateogryRepository interface {
	FindById(Id string) *entity.Category
}