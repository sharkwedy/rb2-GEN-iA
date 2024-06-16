package service

import (
	"context"
	"github.com/yourusername/yourproject/repository"
)

type PersonService struct {
	repo repository.PersonRepository
}

func NewPersonService(repo repository.PersonRepository) *PersonService {
	return &PersonService{
		repo: repo,
	}
}

func (s *PersonService) CreatePerson(ctx context.Context, apelido string, nome string, nascimento string, stack []string) (string, error) {
	return s.repo.Create(ctx, apelido, nome, nascimento, stack)
}

func (s *PersonService) GetPersonByID(ctx context.Context, id string) (*repository.Person, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *PersonService) SearchPeople(ctx context.Context, term string) ([]*repository.Person, error) {
	return s.repo.Search(ctx, term)
}
