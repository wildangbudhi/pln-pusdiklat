package domain

// ForumUsecase is an Interface for Forum Usecase
type ForumUsecase interface {
	FetchCategory() ([]Category, error)
	CreateForum(title string, question string, requestUserID int, categoryID int) (*UUID, error)
	GetForum(forumID UUID, requestUserID int) (*Forum, error)
	UpdateForum(requestUserID int, forumID UUID, title string, question string, categoriID int, requestUserRoles []string) error
	DeleteForum(forumID UUID, requestUserID int, requestUserRoles []string) error
	FetchWithPagination(requestUserID int, offset int, limit int, categoriID *int, topForumSort bool, timelineTimeFrame *TimelineTimeFrame) ([]Forum, error)
	FetchWithPaginationByAuthorID(requestUserID int, authorID int, offset int, limit int, topForumSort bool) ([]Forum, error)
	SearchForum(offset int, limit int, requestUserID int, query string) ([]Forum, error)
	ReactForum(requestUserID int, forumID UUID, userReaction ForumReactionType) error
	ReplyForum(requestUserID int, forumID UUID, answer string) (*UUID, error)
	UpdateForumReplies(requestUserID int, forumRepliesID UUID, answer string, requestUserRoles []string) error
	DeleteForumReplies(requestUserID int, requestUserRoles []string, forumRepliesID UUID) error
	ReactForumReplies(requestUserID int, forumRepliesID UUID, userReaction ForumRepliesReactionType) error
	FetchReplyByForumIDWithPagination(requestUserID int, offset int, limit int, forumID UUID) ([]ForumReplies, error)
	CloseForum(requestUserID int, requestUserRoles []string, forumID UUID) error
}
