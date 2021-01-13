package domain

type EventConsumerUsecase interface {
	UserAuthDataChangesEvent(eventAction string, id int, fullName string, email Email, username string) error
}
