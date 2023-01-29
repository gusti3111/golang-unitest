package service

import (
	"golang_unitest/hello/entity"
	"golang_unitest/hello/repo"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepo = &repo.CategoryRepoMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repo: categoryRepo}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepo.Mock.On("FindById", "1").Return(nil)
	category, err := categoryService.Get("1")
	assert.NotNil(t, err)
	assert.Nil(t, category)

}

func  TestCategoryService_GetFound(t *testing.T) {
	category := entity.Category{
		Id: "1",
		Name: "Hinata",

	}
	categoryRepo.Mock.On("FindById", "2").Return(category)
	result,err := categoryService.Get("2")
	assert.Nil(t,err)
	assert.NotNil(t,result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name,result.Name)

}
	



