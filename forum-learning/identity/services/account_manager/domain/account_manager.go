package domain

// AccountManagerUsecase is an Interface for Account Manager Usecase
type AccountManagerUsecase interface {
	GetUserData(userID int) (*GetUserDataUsecaseResponse, error)
	UpdateUserData(userID int, fullName string, requestUserID int, requestUserRoles []string) (bool, error)
}

type GetUserDataUsecaseResponse struct {
	ID         int    `json:"id"`
	FullName   string `json:"full_name"`
	AvatarFile string `json:"avatar_file"`
	Username   string `json:"username"`
	EmployeeNo string `json:"employee_no"`
	IsEmployee bool   `json:"is_employee"`
}
