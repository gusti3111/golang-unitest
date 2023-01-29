package repo

import "golang_unitest/hello/entity"

type CategoryRepo interface {
	FindById(id string) *entity.Category
}
