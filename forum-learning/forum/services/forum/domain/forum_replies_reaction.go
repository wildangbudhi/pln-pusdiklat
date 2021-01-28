package domain

// ForumRepliesReaction is Struct Mode for forum_replies_reaction table in database
type ForumRepliesReaction struct {
	UserID         int  `json:"user_id"`
	ForumRepliesID UUID `json:"forum_replies_id"`
	UpVote         bool `json:"up_vote"`
	DownVote       bool `json:"down_vote"`
	Agree          bool `json:"agree"`
	Skeptic        bool `json:"skeptic"`
}

// ForumRepliesReactionRepository is a contract of ForumRepliesReactionRepository
type ForumRepliesReactionRepository interface {
	InsertForumRepliesReaction(forumRepliesReaction *ForumRepliesReaction) (int, error)
	GetForumRepliesReactionByUserIDAndForumRepliesID(userID int, forumRepliesID UUID) (*ForumRepliesReaction, error)
	UpdateForumRepliesReactionByUserIDAndForumRepliesID(userID int, forumRepliesID UUID, userReaction ForumRepliesReactionType) (int, error)
}
