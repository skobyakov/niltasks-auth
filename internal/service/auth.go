package service

import "niltasks-auth/internal/models"

type Repository interface {
	TryAuth(*models.AuthRequest) (string, error)
}

type Service struct {
	repository Repository
}

func New(r Repository) *Service {
	return &Service{repository: r}
}

func (s *Service) TryAuth(req *models.AuthRequest) (string, error) {
	session, err := s.repository.TryAuth(req)
	return session, err
}
