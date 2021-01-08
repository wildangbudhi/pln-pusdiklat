package rabbitmq

import (
	"github.com/stretchr/testify/mock"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/event"
)

type UserAuthEventRepositoryMock struct {
	mock.Mock
}

func (repo *UserAuthEventRepositoryMock) PublishDataChangesEvent(userAuthEvent *event.UserAuthEvent) error {
	args := repo.Called()
	return args.Error(0)
}
