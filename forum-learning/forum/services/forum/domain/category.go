package domain

// Category is Struct Mode for category table in database
type Category struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
}

// CategoryRepository is a contract of CategoryRepository
type CategoryRepository interface {
	FetchCategory() ([]Category, error)
}
