package repository

import (
	"niltasks-auth/internal/models"
	"niltasks-auth/pkg/postgres"
)

type Repository struct {
	pg *postgres.Postgres
}

func New(pg *postgres.Postgres) *Repository {
	return &Repository{pg: pg}
}

func (r *Repository) TryAuth(req *models.AuthRequest) (string, error) {

}
