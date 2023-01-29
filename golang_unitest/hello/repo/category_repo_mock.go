package repo

import (
	"golang_unitest/hello/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepoMock struct {
	Mock mock.Mock
}

func (repo *CategoryRepoMock) FindById(id string) *entity.Category {
	argu := repo.Mock.Called(id)
	if argu.Get(0) == nil {
		return nil
	}
	category := argu.Get(0).(entity.Category)
	return &category

}
