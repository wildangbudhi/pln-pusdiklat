package domain

// AccountManagerUsecase is an Interface for Account Manager Usecase
type AccountManagerUsecase interface {
	GetUserData(userID int) (*GetUserDataUsecaseResponse, error)
}

type GetUserDataUsecaseResponse struct {
	ID         int    `json:"id"`
	FullName   string `json:"full_name"`
	AvatarFile string `json:"avatar_file"`
	Email      string `json:"email"`
	Username   string `json:"username"`
}
