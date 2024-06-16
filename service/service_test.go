// Assuming the service package and person.go are already defined.

package service_test

import (
    "testing"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/require"
    "myapp/service"
    "myapp/repository/mocks"
)

type serviceTest struct {
    repo   *mocks.PersonRepository
    service *service.PersonService
}

func setup() *serviceTest {
    repo := &mocks.PersonRepository{}
    svc := service.NewPersonService(repo)
    return &serviceTest{
        repo: repo,
        service: svc,
    }
}

func TestCreatePerson(t *testing.T) {
    s := setup()
    s.repo.On("Create", mock.Anything).Return(nil) // Assuming Create returns error only
    err := s.service.CreatePerson(service.Person{})
    require.NoError(t, err)
    s.repo.AssertExpectations(t)
}