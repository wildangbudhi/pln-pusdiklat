package domain

import (
	"database/sql"
	"encoding/json"
)

// ForumReplies is Struct Model for user_auth forum_replies in database
type ForumReplies struct {
	ID               UUID
	ForumID          UUID
	AuthorUserID     int
	AuthoFullName    sql.NullString
	AuthoUsername    string
	Answer           sql.NullString
	UpVote           int
	DownVote         int
	Agree            int
	Skeptic          int
	IsUpVoted        bool
	IsDownVoted      bool
	IsAgreeToggled   bool
	IsSkepticToggled bool
}

// MarshalJSON Custom Marshal Encoder for ForumReplies Object
func (obj *ForumReplies) MarshalJSON() ([]byte, error) {

	newObject := &struct {
		ID               string `json:"id"`
		ForumID          string `json:"forum_id"`
		AuthorUserID     int    `json:"author_user_id"`
		AuthoFullName    string `json:"author_full_name"`
		AuthoUsername    string `json:"author_username"`
		Answer           string `json:"answer"`
		UpVote           int    `json:"up_vote"`
		DownVote         int    `json:"down_vote"`
		Agree            int    `json:"agree"`
		Skeptic          int    `json:"skepic"`
		IsUpVoted        bool   `json:"is_up_vote"`
		IsDownVoted      bool   `json:"is_down_vote"`
		IsAgreeToggled   bool   `json:"is_agree_toggled"`
		IsSkepticToggled bool   `json:"is_skeptic_toggled"`
	}{
		ID:               obj.ID.GetValue(),
		ForumID:          obj.ForumID.GetValue(),
		AuthorUserID:     obj.AuthorUserID,
		AuthoFullName:    obj.AuthoFullName.String,
		AuthoUsername:    obj.AuthoUsername,
		Answer:           obj.Answer.String,
		UpVote:           obj.UpVote,
		DownVote:         obj.DownVote,
		Agree:            obj.Agree,
		Skeptic:          obj.Skeptic,
		IsUpVoted:        obj.IsUpVoted,
		IsDownVoted:      obj.IsDownVoted,
		IsAgreeToggled:   obj.IsAgreeToggled,
		IsSkepticToggled: obj.IsSkepticToggled,
	}

	return json.Marshal(newObject)

}

// ForumRepliesRepository is a contract of ForumRepliesRepository
type ForumRepliesRepository interface {
	InsertForumReplies(id UUID, userID int, forumID UUID, answer string) (int, error)
	GetForumRepliesByIDWithUserReaction(id UUID, userID int) (*ForumReplies, error)
	FetchForumRepliesByForumIDWithUserReaction(offset int, limit int, forumID UUID, userID int) ([]ForumReplies, error)
	UpdateForumRepliesByID(id UUID, answer string) (int, error)
	DeleteForumRepliesByID(id UUID) (int, error)
}
