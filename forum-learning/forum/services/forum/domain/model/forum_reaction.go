package model

import "github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"

// ForumReaction is Struct Mode for forum_reaction table in database
type ForumReaction struct {
	UserID   int         `json:"user_id"`
	ForumID  domain.UUID `json:"forum_id"`
	UpVote   bool        `json:"up_vote"`
	DownVote bool        `json:"down_vote"`
}

// ForumReactionRepository is a contract of ForumReactionRepository
type ForumReactionRepository interface {
	InsertForumReaction(forumReaction *ForumReaction) (int, error)
	GetForumReactionByUserIDAndForumID(userID int, forumID domain.UUID) (*ForumReaction, error)
	UpdateForumReactionByUserIDAndForumID(userID int, forumID domain.UUID, upVote bool, downVote bool) (int, error)
}
