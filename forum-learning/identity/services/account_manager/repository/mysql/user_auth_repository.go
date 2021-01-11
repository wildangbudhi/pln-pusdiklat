package mysql

import (
	"database/sql"
	"fmt"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/account_manager/domain/model"
)

type userAuthRepository struct {
	db *sql.DB
}

// NewUserAuthRepository is a constructor of userAuthRepository
// which implement UserAuthRepository Interface
func NewUserAuthRepository(db *sql.DB) model.UserAuthRepository {
	return &userAuthRepository{
		db: db,
	}
}

func (repo *userAuthRepository) GetUserAuthByID(id int) (*model.UserAuth, error) {
	var err error
	var queryString string
	userAuth := &model.UserAuth{}

	queryString = "SELECT id, full_name, avatar_file, email, username FROM user_auth WHERE id=?"
	userAuthQueryResult := repo.db.QueryRow(queryString, id)

	err = userAuthQueryResult.Scan(&userAuth.ID, &userAuth.FullName, &userAuth.AvatarFile, &userAuth.Email, &userAuth.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User Not Found")
		}

		return nil, err
	}

	return userAuth, nil
}
