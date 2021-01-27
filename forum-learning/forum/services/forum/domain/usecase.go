package domain

// ForumUsecase is an Interface for Forum Usecase
type ForumUsecase interface {
	UserAuthDataChangesEvent(eventAction string, id int, fullName string, email Email, username string) error
	FetchCategory() ([]Category, error)
	CreateForum(title string, question string, requestUserID int, categoryID int) (*UUID, error)
	GetForum(forumID UUID, requestUserID int) (*Forum, error)
	FetchWithPagination(requestUserID int, offset int, limit int, categoriID *int, topForumSort bool) ([]Forum, error)
	FetchWithPaginationByAuthorID(requestUserID int, authorID int, offset int, limit int, topForumSort bool) ([]Forum, error)
	// DeleteForum(forumID UUID) error
	// ReactForum(requestUserID int, forumID UUID, reactionType string) error
	// SearchForum(requestUserID int, query string) ([]Forum, error)
	// ReplyForum(requestUserID int, forumID UUID, answer string) error
	// UpdateForumReplies(requestUserID int, forumRepliesID UUID, answer string) error
	// DeleteForumReplies(requestUserID int, forumRepliesID UUID) error
	// ReactForumReplies(requestUserID int, forumRepliesID UUID, reactionType string) error
	// CloseForum(requestUserID int, forumID UUID) error
}
