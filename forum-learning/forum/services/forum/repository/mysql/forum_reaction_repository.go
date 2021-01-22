package mysql

import (
	"database/sql"
	"fmt"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain/model"
)

type forumReactionRepository struct {
	db *sql.DB
}

// NewForumReactionRepository is a constructor of forumReactionRepository
// which implement ForumReactionRepository Interface
func NewForumReactionRepository(db *sql.DB) model.ForumReactionRepository {
	return &forumReactionRepository{
		db: db,
	}
}

func (repo *forumReactionRepository) InsertForumReaction(forumReaction *model.ForumReaction) (int, error) {

	var err error
	var queryString string
	var upVote, downVote int

	queryString = `
	INSERT INTO forum_reaction( user_id, forum_id, up_vote, down_vote )
	VALUES( ?, ?, ?, ? )
	`

	if forumReaction.UpVote {
		upVote = 1
	}

	if forumReaction.DownVote {
		downVote = 1
	}

	statement, err := repo.db.Prepare(queryString)

	if err != nil {
		return -1, err
	}

	res, err := statement.Exec(forumReaction.UserID, forumReaction.ForumID.GetValue(), upVote, downVote)

	if err != nil {
		return -1, err
	}

	rowAffected, err := res.RowsAffected()

	if err != nil {
		return -1, err
	}

	return int(rowAffected), nil

}

func (repo *forumReactionRepository) GetForumReactionByUserIDAndForumID(userID int, forumID domain.UUID) (*model.ForumReaction, error) {

	var err error
	var queryString string
	var upVote, downVote int

	queryString = `
		SELECT user_id, forum_id, up_vote, down_vote 
		FROM forum_reaction
		WHERE user_id=? AND forum_id=?
	`

	forumQueryResult := repo.db.QueryRow(queryString, userID, forumID.GetValue())

	var forumIDString string
	forumReactionData := &model.ForumReaction{}

	err = forumQueryResult.Scan(&forumReactionData.UserID, &forumIDString, &upVote, &downVote)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Forum Reaction Not Found")
		}

		return nil, err
	}

	if upVote > 0 {
		forumReactionData.UpVote = true
	}

	if downVote > 0 {
		forumReactionData.DownVote = true
	}

	forumIDUUID, err := domain.NewUUID(forumIDString)

	if err != nil {
		return nil, fmt.Errorf("Forum ID Format Invalid")
	}

	forumReactionData.ForumID = *forumIDUUID

	return forumReactionData, nil

}

func (repo *forumReactionRepository) UpdateForumReactionByUserIDAndForumID(userID int, forumID domain.UUID, upVote bool, downVote bool) (int, error) {

	var err error
	var queryString string
	var rowAffected int64
	var statement *sql.Stmt
	var res sql.Result

	queryString = `UPDATE forum_reaction SET up_vote=?, down_vote=?, updated_at=NOW() WHERE user_id=? AND forum_id=?`

	statement, err = repo.db.Prepare(queryString)

	if err != nil {
		return -1, err
	}

	var upVoteNumeric, downVoteNumeric int

	if upVote {
		upVoteNumeric = 1
	}

	if downVote {
		downVoteNumeric = 1
	}

	res, err = statement.Exec(upVoteNumeric, downVoteNumeric, userID, forumID.GetValue())

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
