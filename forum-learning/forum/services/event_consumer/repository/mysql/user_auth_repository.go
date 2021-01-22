package mysql

import (
	"database/sql"
	"fmt"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/event_consumer/domain/model"
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

func (repo *userAuthRepository) UpdateUserAuthByID(id int, fullName string) (int, error) {

	queryString := "UPDATE user_auth SET full_name=? WHERE id=?"

	statement, err := repo.db.Prepare(queryString)

	if err != nil {
		return -1, err
	}

	res, err := statement.Exec(fullName, id)

	if err != nil {
		return -1, err
	}

	rowAffected, err := res.RowsAffected()

	if err != nil {
		return -1, err
	}

	return int(rowAffected), nil

}

func (repo *userAuthRepository) InsertUserAuth(userAuth *model.UserAuth) (int64, error) {

	var err error
	var queryString string

	queryString = "INSERT INTO user_auth( id, full_name, email, username ) VALUES( ?, ?, ?, ? )"

	statement, err := repo.db.Prepare(queryString)

	if err != nil {
		return -1, err
	}

	res, err := statement.Exec(userAuth.ID, userAuth.FullName, userAuth.Email, userAuth.Username)

	if err != nil {
		return -1, err
	}

	userAuthID, err := res.LastInsertId()

	if err != nil {
		return -1, err
	}

	return userAuthID, nil
}

func (repo *userAuthRepository) DeleteUserAuthByID(id int) (int, error) {

	var err error
	var queryString string

	queryString = "DELETE FROM user_auth WHERE id=?"

	statement, err := repo.db.Prepare(queryString)

	if err != nil {
		return -1, err
	}

	res, err := statement.Exec(id)

	if err != nil {
		return -1, err
	}

	rowAffected, err := res.RowsAffected()

	return int(rowAffected), nil

}