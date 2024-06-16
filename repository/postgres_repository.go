package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	DB *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{DB: db}
}

func (r *PostgresRepository) FindByID(ctx context.Context, id uuid.UUID) (*Person, error) {
	person := &Person{}
	err := r.DB.QueryRowContext(ctx, "SELECT id, apelido, nome, nascimento, stack FROM people WHERE id = $1", id).Scan(&person.ID, &person.Apelido, &person.Nome, &person.Nascimento, &person.Stack)
	if err != nil {
		return nil, err
	}
	return person, nil
}

func (r *PostgresRepository) Create(ctx context.Context, p *Person) error {
	_, err := r.DB.ExecContext(ctx, "INSERT INTO people (id, apelido, nome, nascimento, stack) VALUES ($1, $2, $3, $4, $5)", p.ID, p.Apelido, p.Nome, p.Nascimento, pq.Array(p.Stack))
	return err
}

func (r *PostgresRepository) Search(ctx context.Context, term string) ([]*Person, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id, apelido, nome, nascimento, stack FROM people WHERE apelido LIKE $1 OR nome LIKE $1 OR $1 = ANY(stack)", fmt.Sprintf("%%%s%%", term))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	persons := []*Person{}
	for rows.Next() {
		person := &Person{}
		if err := rows.Scan(&person.ID, &person.Apelido, &person.Nome, &person.Nascimento, &person.Stack); err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}

	return persons, nil
}
