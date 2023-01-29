package service

import (
	"errors"
	"golang_unitest/hello/entity"
	"golang_unitest/hello/repo"
)

type CategoryService struct {
	Repo repo.CategoryRepo
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repo.FindById(id)
	if category == nil {
		return category, errors.New("category not found")

	} else {
		return category, nil
	}

}
