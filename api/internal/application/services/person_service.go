package services

import (
	"api/internal/domain/entities"
	"api/internal/domain/repositories"
)

type PersonService struct {
	repo repositories.PersonRepository
}

func NewPersonService(repo repositories.PersonRepository) *PersonService {
	return &PersonService{repo: repo}
}

func (s *PersonService) Create(person *entities.Person) (*entities.Person, error) {
	return s.repo.Create(person)
}

func (s *PersonService) FindById(id string) (*entities.Person, error) {
	return s.repo.FindById(id)
}

func (s *PersonService) FindAll() ([]*entities.Person, error) {
	return s.repo.FindAll()
}

func (s *PersonService) Update(person *entities.Person) (*entities.Person, error) {
	return s.repo.Update(person)
}

func (s *PersonService) Delete(id string) error {
	return s.repo.Delete(id)
}
