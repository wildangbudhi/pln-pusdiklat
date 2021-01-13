package model

import "github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"

// ForumRepliesReaction is Struct Mode for forum_replies_reaction table in database
type ForumRepliesReaction struct {
	UserID         int         `json:"user_id"`
	ForumID        domain.UUID `json:"forum_id"`
	ForumRepliesID domain.UUID `json:"forum_replies_id"`
	UpVote         bool        `json:"up_vote"`
	DownVote       bool        `json:"down_vote"`
	Agree          bool        `json:"agree"`
	Skeptic        bool        `json:"skeptic"`
}

// ForumRepliesReactionRepository is a contract of ForumRepliesReactionRepository
type ForumRepliesReactionRepository interface {
	InsertForumRepliesReaction(forumRepliesReaction *ForumRepliesReaction) (int, error)
	GetForumRepliesReactionByUserIDAndForumRepliesID(userID int, forumRepliesID domain.UUID) (*ForumRepliesReaction, error)
	UpdateForumRepliesReactionByUserIDAndForumRepliesID(userID int, forumRepliesID domain.UUID, upVote bool, downVote bool, agree bool, skeptic bool) (int, error)
}
