package manager

import "furnishop/server/usecase"

type UseCaseManager interface {
	ProductCategoryUseCase() usecase.ProductCategoryUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) ProductCategoryUseCase() usecase.ProductCategoryUseCase {
	return usecase.NewProductCategoryUseCase(u.repo.ProductCategoryRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
