package domain

import (
	"database/sql"
	"encoding/json"
)

// Forum is Struct Mode for forum table in database
type Forum struct {
	ID            UUID
	Title         string
	Question      sql.NullString
	AuthorUserID  int
	AuthoFullName sql.NullString
	AuthoUsername string
	Status        sql.NullString
	CategoryID    int
	CategoryName  string
	UpVote        int
	DownVote      int
	IsUpVoted     bool
	IsDownVoted   bool
	RepliesCount  int
}

// MarshalJSON Custom Marshal Encoder for Forum Object
func (obj *Forum) MarshalJSON() ([]byte, error) {

	newObject := &struct {
		ID            string `json:"id"`
		Title         string `json:"title"`
		Question      string `json:"question"`
		AuthorUserID  int    `json:"author_user_id"`
		AuthoFullName string `json:"author_full_name"`
		AuthoUsername string `json:"author_username"`
		Status        string `json:"status"`
		CategoryID    int    `json:"category_id"`
		CategoryName  string `json:"category_name"`
		UpVote        int    `json:"up_vote"`
		DownVote      int    `json:"down_vote"`
		IsUpVoted     bool   `json:"is_up_voted"`
		IsDownVoted   bool   `json:"is_down_voted"`
		RepliesCount  int    `json:"replies_count"`
	}{
		ID:            obj.ID.GetValue(),
		Title:         obj.Title,
		Question:      obj.Question.String,
		AuthorUserID:  obj.AuthorUserID,
		AuthoFullName: obj.AuthoFullName.String,
		AuthoUsername: obj.AuthoUsername,
		UpVote:        obj.UpVote,
		DownVote:      obj.DownVote,
		IsUpVoted:     obj.IsUpVoted,
		IsDownVoted:   obj.IsDownVoted,
		Status:        obj.Status.String,
		CategoryID:    obj.CategoryID,
		CategoryName:  obj.CategoryName,
		RepliesCount:  obj.RepliesCount,
	}

	return json.Marshal(newObject)

}

// ForumRepository is a contract of ForumRepository
type ForumRepository interface {
	InsertForum(id UUID, title string, question sql.NullString, authorUserID int, categoryID int, status string) (int, error)
	GetForumByIDWithUserReaction(id UUID, userID int) (*Forum, error)
	FetchForumWithUserReaction(offset int, limit int, userID int) ([]Forum, error)
	FetchForumByAuthorIDWithUserReaction(authorID int, offset int, limit int, userID int) ([]Forum, error)
	SearchByTitleAndQuestionWithUserReaction(offset int, limit int, userID int, query string) ([]Forum, error)
	DeleteForumByID(id UUID) (int, error)
	UpdateForumByID(id UUID, title string, question sql.NullString, categoryID int) (int, error)
	UpdateForumStatusByID(id UUID, status string) (int, error)
}
