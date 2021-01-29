package usecase

import (
	"database/sql"
	"fmt"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

func (usecase *forumUsecase) CreateForum(title string, question string, requestUserID int, categoryID int) (*domain.UUID, error) {

	forumID := domain.NewUUID()
	forumStatus := "OPEN"

	if title == "" {
		return nil, fmt.Errorf("Title Cannot Be Empty")
	}

	_, err := usecase.forumRepository.InsertForum(
		*forumID,
		title,
		sql.NullString{String: question, Valid: true},
		requestUserID,
		categoryID,
		forumStatus,
	)

	if err != nil {
		return nil, err
	}

	return forumID, nil
}
