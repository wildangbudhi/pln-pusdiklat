package mysql

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/identity/services/authentication/domain/model"
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
	userAuth.Roles = make([]model.Roles, 0)

	queryString = "SELECT id, full_name, email, username, password FROM user_auth WHERE id=?"
	userAuthQueryResult := repo.db.QueryRow(queryString, id)

	err = userAuthQueryResult.Scan(&userAuth.ID, &userAuth.FullName, &userAuth.Email, &userAuth.Username, &userAuth.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User Not Found")
		}

		return nil, err
	}

	queryString = "SELECT r.id, r.roles_name FROM user_roles ur JOIN roles r on r.id = ur.role_id WHERE ur.user_id = ?"
	rolesQueryResult, err := repo.db.Query(queryString, id)
	defer rolesQueryResult.Close()

	if err != nil {
		return nil, err
	}

	for rolesQueryResult.Next() {
		role := model.Roles{}
		err := rolesQueryResult.Scan(&role.ID, &role.RoleName)

		if err != nil {
			return nil, err
		}

		userAuth.Roles = append(userAuth.Roles, role)
	}

	return userAuth, nil

}

func (repo *userAuthRepository) GetUserAuthByEmail(email string) (*model.UserAuth, error) {

	var err error
	var queryString string
	userAuth := &model.UserAuth{}
	userAuth.Roles = make([]model.Roles, 0)

	queryString = "SELECT id, full_name, email, username, password FROM user_auth WHERE email=?"
	userAuthQueryResult := repo.db.QueryRow(queryString, email)

	err = userAuthQueryResult.Scan(&userAuth.ID, &userAuth.FullName, &userAuth.Email, &userAuth.Username, &userAuth.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User Not Found")
		}

		return nil, err
	}

	queryString = "SELECT r.id, r.roles_name FROM user_roles ur JOIN roles r on r.id = ur.role_id WHERE ur.user_id = ?"
	rolesQueryResult, err := repo.db.Query(queryString, userAuth.ID)
	defer rolesQueryResult.Close()

	if err != nil {
		return nil, err
	}

	for rolesQueryResult.Next() {
		role := model.Roles{}
		err := rolesQueryResult.Scan(&role.ID, &role.RoleName)

		if err != nil {
			return nil, err
		}

		userAuth.Roles = append(userAuth.Roles, role)
	}

	return userAuth, nil

}

func (repo *userAuthRepository) GetUserAuthByUsername(username string) (*model.UserAuth, error) {

	var err error
	var queryString string
	userAuth := &model.UserAuth{}
	userAuth.Roles = make([]model.Roles, 0)

	queryString = "SELECT id, full_name, email, username, password FROM user_auth WHERE username=?"
	userAuthQueryResult := repo.db.QueryRow(queryString, username)

	err = userAuthQueryResult.Scan(&userAuth.ID, &userAuth.FullName, &userAuth.Email, &userAuth.Username, &userAuth.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User Not Found")
		}

		return nil, err
	}

	queryString = "SELECT r.id, r.roles_name FROM user_roles ur JOIN roles r on r.id = ur.role_id WHERE ur.user_id = ?"
	rolesQueryResult, err := repo.db.Query(queryString, userAuth.ID)
	defer rolesQueryResult.Close()

	if err != nil {
		return nil, err
	}

	for rolesQueryResult.Next() {
		role := model.Roles{}
		err := rolesQueryResult.Scan(&role.ID, &role.RoleName)

		if err != nil {
			return nil, err
		}

		userAuth.Roles = append(userAuth.Roles, role)
	}

	return userAuth, nil

}

func (repo *userAuthRepository) InsertUserAuth(userAuth *model.UserAuth) (int64, error) {

	var err error
	var queryString string

	if len(userAuth.Roles) == 0 {
		return -1, fmt.Errorf("Roles Cannot Be Empty")
	}

	queryString = "INSERT INTO user_auth( full_name, email, username, password ) VALUES( ?, ?, ?, ? )"

	statement, err := repo.db.Prepare(queryString)

	if err != nil {
		return -1, err
	}

	res, err := statement.Exec(userAuth.FullName, userAuth.Email, userAuth.Username, userAuth.Password)

	if err != nil {
		return -1, err
	}

	userAuthID, err := res.LastInsertId()

	if err != nil {
		return -1, err
	}

	queryString = "INSERT INTO user_roles( user_id, role_id ) VALUES "
	values := []interface{}{}

	for i := 0; i < len(userAuth.Roles); i++ {
		if i != 0 {
			queryString += ","
		}

		queryString += "( ?, ? )"
		values = append(values, userAuthID, userAuth.Roles[i].ID)
	}

	statement, err = repo.db.Prepare(queryString)

	if err != nil {
		return -1, err
	}

	res, err = statement.Exec(values...)

	if err != nil {
		return -1, err
	}

	return userAuthID, nil
}

func (repo *userAuthRepository) CountRoleActivitiesPermission(method string, url string, roleIDList []int) (int, error) {

	var err error
	var queryString string

	queryString = `SELECT COUNT(1) FROM role_activities ra JOIN activities a on ra.activities_id = a.id  
	WHERE a.method = ? AND REGEXP_LIKE(?, a.url_regex ) AND ra.role_id in (`

	for i := 0; i < len(roleIDList); i++ {
		queryString += strconv.Itoa(roleIDList[i])

		if i != len(roleIDList)-1 {
			queryString += ","
		} else {
			queryString += ")"
		}

	}

	userAuthQueryResult := repo.db.QueryRow(queryString, method, url)

	var countPermission int

	err = userAuthQueryResult.Scan(&countPermission)

	if err != nil {
		if err == sql.ErrNoRows {
			return -1, fmt.Errorf("Activities Data Now Found")
		}

		return -1, err
	}

	return countPermission, nil

}
