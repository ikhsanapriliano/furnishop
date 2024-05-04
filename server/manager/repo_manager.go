package manager

import "furnishop/server/repository"

type RepoManager interface {
	ProductCategoryRepo() repository.ProductCategoryRepository
}

type repoManager struct {
	infra InfraManager
}

func (r *repoManager) ProductCategoryRepo() repository.ProductCategoryRepository {
	return repository.NewProductCategoryRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
