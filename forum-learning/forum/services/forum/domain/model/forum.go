package model

import (
	"database/sql"
	"encoding/json"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

// Forum is Struct Mode for forum table in database
type Forum struct {
	ID            domain.UUID
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
		CategoryName  string `json:"category_name"`
		UpVote        int    `json:"up_vote"`
		DownVote      int    `json:"down_vote"`
		IsUpVote      bool   `json:"is_up_vote"`
		IsDownVote    bool   `json:"is_down_vote"`
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
		IsUpVote:      obj.IsUpVoted,
		IsDownVote:    obj.IsDownVoted,
		Status:        obj.Status.String,
		CategoryName:  obj.CategoryName,
		RepliesCount:  obj.RepliesCount,
	}

	return json.Marshal(newObject)

}

// ForumRepository is a contract of ForumRepository
type ForumRepository interface {
	InsertForum(title string, question sql.NullString, authorUserID int, categoryID int, status string) (int, error)
	GetForumByID(id domain.UUID) (*Forum, error)
	GetForumByIDWithUserReaction(id domain.UUID, userID int) (*Forum, error)
	FetchForum(offset int, limit int) ([]Forum, error)
	FetchForumWithUserReaction(offset int, limit int, userID int) ([]Forum, error)
	FetchForumByAuthorID(authorID int, offset int, limit int) ([]Forum, error)
	DeleteForumByID(id domain.UUID) (int, error)
	UpdateForumByID(id domain.UUID, title string, question sql.NullString, categoryID int) (int, error)
	UpdateForumStatusByID(id domain.UUID, status string) (int, error)
}
