package mysql

import (
	"database/sql"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type categoryRepository struct {
	db *sql.DB
}

// NewCategoryRepository is a constructor of categoryRepository
// which implement CategoryRepository Interface
func NewCategoryRepository(db *sql.DB) domain.CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (repo *categoryRepository) FetchCategory() ([]domain.Category, error) {

	var err error
	var queryString string

	queryString = "SELECT id, category_name FROM category"

	categoryQueryResult, err := repo.db.Query(queryString)
	defer categoryQueryResult.Close()

	if err != nil {
		return nil, err
	}

	categoryList := make([]domain.Category, 0)

	for categoryQueryResult.Next() {
		category := domain.Category{}
		err := categoryQueryResult.Scan(&category.ID, &category.CategoryName)

		if err != nil {
			return nil, err
		}

		categoryList = append(categoryList, category)
	}

	return categoryList, nil

}
