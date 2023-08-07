package repository

import "fmt"

type Repository struct{}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) TryAuth() {
	fmt.Println("Hello from auth module!")
}
