package mysql

import (
	"database/sql"
	"fmt"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type forumRepliesReactionRepository struct {
	db *sql.DB
}

// NewForumRepliesReactionRepository is a constructor of forumRepliesReactionRepository
// which implement ForumRepliesReactionRepository Interface
func NewForumRepliesReactionRepository(db *sql.DB) domain.ForumRepliesReactionRepository {
	return &forumRepliesReactionRepository{
		db: db,
	}
}

func (repo *forumRepliesReactionRepository) InsertForumRepliesReaction(forumRepliesReaction *domain.ForumRepliesReaction) (int, error) {

	var err error
	var queryString string
	var upVote, downVote, agree, skeptic int

	queryString = `
	INSERT INTO forum_replies_reaction( user_id, forum_replies_id, up_vote, down_vote, agree, skeptic )
	VALUES( ?, ?, ?, ?, ?, ? )
	`
	if forumRepliesReaction.UpVote {
		upVote = 1
	}

	if forumRepliesReaction.DownVote {
		downVote = 1
	}

	if forumRepliesReaction.Agree {
		agree = 1
	}

	if forumRepliesReaction.Skeptic {
		skeptic = 1
	}

	statement, err := repo.db.Prepare(queryString)

	if err != nil {
		return -1, err
	}

	res, err := statement.Exec(
		forumRepliesReaction.UserID,
		forumRepliesReaction.ForumRepliesID.GetValue(),
		upVote,
		downVote,
		agree,
		skeptic,
	)

	if err != nil {
		return -1, err
	}

	rowAffected, err := res.RowsAffected()

	if err != nil {
		return -1, err
	}

	return int(rowAffected), nil

}

func (repo *forumRepliesReactionRepository) GetForumRepliesReactionByUserIDAndForumRepliesID(userID int, forumRepliesID domain.UUID) (*domain.ForumRepliesReaction, error) {

	var err error
	var queryString string

	queryString = `
		SELECT user_id, forum_replies_id, up_vote, down_vote, agree, skeptic 
		FROM forum_replies_reaction
		WHERE user_id=? AND forum_replies_id=?
	`

	forumRepliesReactionQueryResult := repo.db.QueryRow(queryString, userID, forumRepliesID.GetValue())

	var forumRepliesIDString string
	var upVote, downVote, agree, skeptic int

	forumRepliesReactionData := &domain.ForumRepliesReaction{}

	err = forumRepliesReactionQueryResult.Scan(
		&forumRepliesReactionData.UserID,
		&forumRepliesIDString,
		&upVote,
		&downVote,
		&agree,
		&skeptic,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Forum Replies Reaction Not Found")
		}

		return nil, err
	}

	if upVote > 0 {
		forumRepliesReactionData.UpVote = true
	}

	if downVote > 0 {
		forumRepliesReactionData.DownVote = true
	}

	if agree > 0 {
		forumRepliesReactionData.Agree = true
	}

	if skeptic > 0 {
		forumRepliesReactionData.Skeptic = true
	}

	forumRepliesReactionData.ForumRepliesID = forumRepliesID

	return forumRepliesReactionData, nil

}

func (repo *forumRepliesReactionRepository) UpdateForumRepliesReactionByUserIDAndForumRepliesID(userID int, forumRepliesID domain.UUID, userReaction domain.ForumRepliesReactionType) (int, error) {

	var err error
	var queryString string
	var rowAffected int64
	var statement *sql.Stmt
	var res sql.Result

	queryString = `
	UPDATE forum_replies_reaction SET up_vote=?, down_vote=?, agree=?, skeptic=?, updated_at=NOW() 
	WHERE user_id=? AND forum_replies_id=?`

	statement, err = repo.db.Prepare(queryString)

	if err != nil {
		return -1, err
	}

	var upVoteNumeric, downVoteNumeric, agreeNumeric, skepticNumerik int

	if userReaction.IsUpVoteToggled() {
		upVoteNumeric = 1
	}

	if userReaction.IsDownVoteToggled() {
		downVoteNumeric = 1
	}

	if userReaction.IsAgreeToggled() {
		agreeNumeric = 1
	}

	if userReaction.IsSkepticToggled() {
		skepticNumerik = 1
	}

	res, err = statement.Exec(upVoteNumeric, downVoteNumeric, agreeNumeric, skepticNumerik, userID, forumRepliesID.GetValue())

	if err != nil {
		return -1, err
	}

	rowAffected, err = res.RowsAffected()

	if err != nil {
		return -1, err
	}

	if rowAffected == 0 {
		return -1, fmt.Errorf("Forum Reaction With Specific Forum ID and User ID Not Found")
	}

	return int(rowAffected), nil

}
