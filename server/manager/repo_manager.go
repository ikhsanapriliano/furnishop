package manager

type RepoManager interface {
}

type repoManager struct {
	infra InfraManager
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
