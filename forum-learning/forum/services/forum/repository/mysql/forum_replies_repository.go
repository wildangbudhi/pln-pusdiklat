package mysql

import (
	"database/sql"
	"fmt"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type forumRepliesRepository struct {
	db *sql.DB
}

// NewForumRepliesRepository is a constructor of forumRepliesRepository
// which implement ForumRepliesRepository Interface
func NewForumRepliesRepository(db *sql.DB) domain.ForumRepliesRepository {
	return &forumRepliesRepository{
		db: db,
	}
}

func (repo *forumRepliesRepository) InsertForumReplies(id domain.UUID, userID int, forumID domain.UUID, answer string) (int, error) {

	var err error
	var queryString string

	queryString = `
	INSERT INTO forum_replies( id, forum_id, author_user_id, answer )
	VALUES( ?, ?, ?, ? )
	`

	statement, err := repo.db.Prepare(queryString)

	if err != nil {
		return -1, err
	}

	res, err := statement.Exec(id.GetValue(), forumID.GetValue(), userID, answer)

	if err != nil {
		return -1, err
	}

	rowAffected, err := res.RowsAffected()

	if err != nil {
		return -1, err
	}

	return int(rowAffected), nil

}

func (repo *forumRepliesRepository) GetForumRepliesByIDWithUserReaction(id domain.UUID, userID int) (*domain.ForumReplies, error) {

	var err error
	var queryString string

	queryString = `
		SELECT 
			fr.id,
			fr.forum_id,
			fr.author_user_id,
			ua.full_name,
			ua.username,
			fr.answer,
			IFNULL(frr.up_vote, 0) up_vote,
			IFNULL(frr.down_vote, 0) down_vote,
			IFNULL(frr.agree, 0) agree,
			IFNULL(frr.skeptic, 0) skeptic,
			IFNULL(ufrr.up_vote, 0) is_up_vote,
			IFNULL(ufrr.down_vote, 0) is_down_vote,
			IFNULL(ufrr.agree, 0) is_agree_toggled,
			IFNULL(ufrr.skeptic, 0) is_skeptic_toggled
		FROM 
			forum_replies fr 
		LEFT JOIN user_auth ua ON ua.id = fr.author_user_id 
		LEFT JOIN (
			SELECT frr.forum_replies_id, SUM( frr.up_vote ) up_vote, SUM( frr.down_vote ) down_vote, SUM( frr.agree ) agree, SUM( frr.skeptic ) skeptic
			FROM forum_replies_reactions frr 
			GROUP BY frr.forum_replies_id  
		) frr ON frr.forum_replies_id = fr.id 
		LEFT JOIN (
			SELECT frr.forum_replies_id, SUM( frr.up_vote ) up_vote, SUM( frr.down_vote ) down_vote, SUM( frr.agree ) agree, SUM( frr.skeptic ) skeptic
			FROM forum_replies_reactions frr 
			WHERE frr.user_id = ?
			GROUP BY frr.forum_replies_id
		) ufrr ON ufrr.forum_replies_id = fr.id 
		WHERE 
			fr.id = ?
	`

	forumRepliesQueryResut := repo.db.QueryRow(queryString, userID, id.GetValue())

	var isUpVoted, isDownVoted, IsAgreeToggled, IsSkepticToggled int
	var idString, forumIDString string

	forumRepliesData := &domain.ForumReplies{}

	err = forumRepliesQueryResut.Scan(
		&idString,
		&forumIDString,
		&forumRepliesData.AuthorUserID,
		&forumRepliesData.AuthoFullName,
		&forumRepliesData.AuthoUsername,
		&forumRepliesData.Answer,
		&forumRepliesData.UpVote,
		&forumRepliesData.DownVote,
		&forumRepliesData.Agree,
		&forumRepliesData.Skeptic,
		&isUpVoted,
		&isDownVoted,
		&IsAgreeToggled,
		&IsSkepticToggled,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Forum Replies Not Found")
		}

		return nil, err
	}

	if isUpVoted > 0 {
		forumRepliesData.IsUpVoted = true
	}

	if isDownVoted > 0 {
		forumRepliesData.IsDownVoted = true
	}

	if IsAgreeToggled > 0 {
		forumRepliesData.IsAgreeToggled = true
	}

	if IsSkepticToggled > 0 {
		forumRepliesData.IsSkepticToggled = true
	}

	forumRepliesID, err := domain.NewUUIDFromString(idString)

	if err != nil {
		return nil, fmt.Errorf("ID Format Invalid")
	}

	forumRepliesData.ID = *forumRepliesID

	forumID, err := domain.NewUUIDFromString(forumIDString)

	if err != nil {
		return nil, fmt.Errorf("ID Format Invalid")
	}

	forumRepliesData.ForumID = *forumID

	return forumRepliesData, nil

}

func (repo *forumRepliesRepository) FetchForumRepliesByForumIDWithUserReaction(offset int, limit int, forumID domain.UUID, userID int) ([]domain.ForumReplies, error) {

	var err error
	var queryString string

	forumRepliesList := make([]domain.ForumReplies, 0)

	queryString = `
		SELECT 
			fr.id,
			fr.forum_id,
			fr.author_user_id,
			ua.full_name,
			ua.username,
			fr.answer,
			IFNULL(frr.up_vote, 0) up_vote,
			IFNULL(frr.down_vote, 0) down_vote,
			IFNULL(frr.agree, 0) agree,
			IFNULL(frr.skeptic, 0) skeptic,
			IFNULL(ufrr.up_vote, 0) is_up_vote,
			IFNULL(ufrr.down_vote, 0) is_down_vote,
			IFNULL(ufrr.agree, 0) is_agree_toggled,
			IFNULL(ufrr.skeptic, 0) is_skeptic_toggled
		FROM 
			forum_replies fr 
		LEFT JOIN user_auth ua ON ua.id = fr.author_user_id 
		LEFT JOIN (
			SELECT frr.forum_replies_id, SUM( frr.up_vote ) up_vote, SUM( frr.down_vote ) down_vote, SUM( frr.agree ) agree, SUM( frr.skeptic ) skeptic
			FROM forum_replies_reactions frr 
			GROUP BY frr.forum_replies_id  
		) frr ON frr.forum_replies_id = fr.id 
		LEFT JOIN (
			SELECT frr.forum_replies_id, SUM( frr.up_vote ) up_vote, SUM( frr.down_vote ) down_vote, SUM( frr.agree ) agree, SUM( frr.skeptic ) skeptic
			FROM forum_replies_reactions frr 
			WHERE frr.user_id = ?
			GROUP BY frr.forum_replies_id
		) ufrr ON ufrr.forum_replies_id = fr.id 
		WHERE fr.forum_id = ?
		ORDER BY up_vote DESC, agree DESC, fr.created_at DESC
		LIMIT ? OFFSET ?
	`

	forumRepliesQueryResult, err := repo.db.Query(queryString, userID, forumID.GetValue(), limit, offset)
	defer forumRepliesQueryResult.Close()

	if err != nil {
		return nil, err
	}

	for forumRepliesQueryResult.Next() {

		var isUpVoted, isDownVoted, IsAgreeToggled, IsSkepticToggled int
		var idString, forumIDString string

		forumRepliesData := domain.ForumReplies{}

		err = forumRepliesQueryResult.Scan(
			&idString,
			&forumIDString,
			&forumRepliesData.AuthorUserID,
			&forumRepliesData.AuthoFullName,
			&forumRepliesData.AuthoUsername,
			&forumRepliesData.Answer,
			&forumRepliesData.UpVote,
			&forumRepliesData.DownVote,
			&forumRepliesData.Agree,
			&forumRepliesData.Skeptic,
			&isUpVoted,
			&isDownVoted,
			&IsAgreeToggled,
			&IsSkepticToggled,
		)

		if err != nil {
			return nil, err
		}

		if isUpVoted > 0 {
			forumRepliesData.IsUpVoted = true
		}

		if isDownVoted > 0 {
			forumRepliesData.IsDownVoted = true
		}

		if IsAgreeToggled > 0 {
			forumRepliesData.IsAgreeToggled = true
		}

		if IsSkepticToggled > 0 {
			forumRepliesData.IsSkepticToggled = true
		}

		forumRepliesID, err := domain.NewUUIDFromString(idString)

		if err != nil {
			return nil, fmt.Errorf("ID Format Invalid")
		}

		forumRepliesData.ID = *forumRepliesID

		forumID, err := domain.NewUUIDFromString(forumIDString)

		if err != nil {
			return nil, fmt.Errorf("ID Format Invalid")
		}

		forumRepliesData.ForumID = *forumID

		forumRepliesList = append(forumRepliesList, forumRepliesData)

	}

	return forumRepliesList, nil

}

func (repo *forumRepliesRepository) UpdateForumRepliesByID(id domain.UUID, answer string) (int, error) {

	var err error
	var queryString string
	var rowAffected int64
	var statement *sql.Stmt
	var res sql.Result

	queryString = `UPDATE forum_replies SET answer=?, updated_at=NOW() WHERE id=?`

	statement, err = repo.db.Prepare(queryString)

	if err != nil {
		return -1, err
	}

	res, err = statement.Exec(answer, id.GetValue())

	if err != nil {
		return -1, err
	}

	rowAffected, err = res.RowsAffected()

	if err != nil {
		return -1, err
	}

	if rowAffected == 0 {
		return -1, fmt.Errorf("Forum Replies With Specific ID Not Found")
	}

	return int(rowAffected), nil

}

func (repo *forumRepliesRepository) DeleteForumRepliesByID(id domain.UUID) (int, error) {

	var err error
	var queryString string
	var rowAffected int64
	var statement *sql.Stmt
	var res sql.Result

	queryString = `DELETE FROM forum_replies WHERE id=?`

	statement, err = repo.db.Prepare(queryString)

	if err != nil {
		return -1, err
	}

	res, err = statement.Exec(id.GetValue())

	if err != nil {
		return -1, err
	}

	rowAffected, err = res.RowsAffected()

	if err != nil {
		return -1, err
	}

	if rowAffected == 0 {
		return -1, fmt.Errorf("Forum Replies With Specific ID Not Found")
	}

	return int(rowAffected), nil

}
