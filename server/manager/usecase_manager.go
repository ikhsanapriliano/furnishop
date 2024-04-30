package manager

type UseCaseManager interface {
}

type useCaseManager struct {
	repo RepoManager
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
