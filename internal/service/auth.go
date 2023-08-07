package service

type Repository interface {
	TryAuth()
}

type Service struct {
	repository Repository
}

func New(r Repository) *Service {
	return &Service{repository: r}
}

func (s *Service) TryAuth() {
	s.repository.TryAuth()
}
