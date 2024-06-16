package repository

import (
	"context"
	"github.com/google/uuid"
)

type Person struct {
	ID          uuid.UUID `db:"id"`
	Nickname    string    `db:"apelido"`
	Name        string    `db:"nome"`
	Birthdate   string    `db:"nascimento"`
	TechStack   []string  `db:"stack"`
}

type PersonRepository interface {
	Create(ctx context.Context, p Person) (uuid.UUID, error)
	FindByID(ctx context.Context, id uuid.UUID) (Person, error)
	Search(ctx context.Context, query string) ([]Person, error)
}