package repository

import (
	"fmt"
	"niltasks-auth/pkg/postgres"
)

type Repository struct {
	pg *postgres.Postgres
}

func New(pg *postgres.Postgres) *Repository {
	return &Repository{pg: pg}
}

func (r *Repository) TryAuth() {
	fmt.Println("Hello from auth module!")
}
