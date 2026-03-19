package service

import (
	"errors"
	"go-unit-test/entity"
	"go-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CateogryRepository
}

func (service CategoryService) Get(id string) (*entity.Category,error){
	category := service.Repository.FindById(id)
	if  category == nil {
		return nil, errors.New("CAtegory Not FOund")
	}else {
		return category, nil
	}
}