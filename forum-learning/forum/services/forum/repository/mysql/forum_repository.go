package mysql

import (
	"database/sql"
	"fmt"

	"github.com/wildangbudhi/pln-pusdiklat/forum-learning/forum/services/forum/domain"
)

type forumRepository struct {
	db *sql.DB
}

// NewForumRepository is a constructor of forumRepository
// which implement ForumRepository Interface
func NewForumRepository(db *sql.DB) domain.ForumRepository {
	return &forumRepository{
		db: db,
	}
}

func (repo *forumRepository) InsertForum(id domain.UUID, title string, question sql.NullString, authorUserID int, categoryID int, status string) (int, error) {

	var err error
	var queryString string

	queryString = `
	INSERT INTO forum( id, title, question, author_user_id, category_id, status )
	VALUES( ?, ?, ?, ?, ?, ? )
	`

	statement, err := repo.db.Prepare(queryString)

	if err != nil {
		return -1, err
	}

	res, err := statement.Exec(id.GetValue(), title, question.String, authorUserID, categoryID, status)

	if err != nil {
		return -1, err
	}

	rowAffected, err := res.RowsAffected()

	if err != nil {
		return -1, err
	}

	return int(rowAffected), nil

}

func (repo *forumRepository) GetForumByIDWithUserReaction(id domain.UUID, userID int) (*domain.Forum, error) {

	var err error
	var queryString string
	var isUpVoted, isDownVoted int

	queryString = `
		SELECT 
			f.id,
			f.title,
			f.question, 
			f.author_user_id,
			ua.full_name,
			ua.username,
			f.status,
			f.category_id,
			c.category_name,
			IFNULL(fr.up_vote, 0) up_vote,
			IFNULL(fr.down_vote, 0) down_vote,
			IFNULL(frp.replies_count, 0) replies_count,
			IFNULL(ufr.up_vote, 0) is_up_vote,
			IFNULL(ufr.down_vote, 0) is_down_vote
		FROM 
			forum f
		LEFT JOIN user_auth ua ON ua.id = f.author_user_id 
		LEFT JOIN category c ON c.id = f.category_id 
		LEFT JOIN (
			SELECT fr.forum_id, SUM( fr.up_vote ) up_vote, SUM( fr.down_vote ) down_vote
			FROM forum_reaction fr 
			GROUP BY fr.forum_id 
		) fr ON fr.forum_id = f.id 
		LEFT JOIN (
			SELECT fr.forum_id, COUNT(1) replies_count FROM forum_replies fr GROUP BY fr.forum_id 
		) frp ON frp.forum_id = f.id 
		LEFT JOIN (
			SELECT fr.forum_id, SUM( fr.up_vote ) up_vote, SUM( fr.down_vote ) down_vote
			FROM forum_reaction fr 
			WHERE fr.user_id = ?
			GROUP BY fr.forum_id 
		) ufr ON ufr.forum_id = f.id 
		WHERE 
			f.id = ?
	`

	forumQueryResult := repo.db.QueryRow(queryString, userID, id.GetValue())

	var forumID string
	forumData := &domain.Forum{}

	err = forumQueryResult.Scan(
		&forumID,
		&forumData.Title,
		&forumData.Question,
		&forumData.AuthorUserID,
		&forumData.AuthoFullName,
		&forumData.AuthoUsername,
		&forumData.Status,
		&forumData.CategoryID,
		&forumData.CategoryName,
		&forumData.UpVote,
		&forumData.DownVote,
		&forumData.RepliesCount,
		&isUpVoted,
		&isDownVoted,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Forum Not Found")
		}

		return nil, err
	}

	if isUpVoted > 0 {
		forumData.IsUpVoted = true
	}

	if isDownVoted > 0 {
		forumData.IsDownVoted = true
	}

	forumIDUUID, err := domain.NewUUIDFromString(forumID)

	if err != nil {
		return nil, fmt.Errorf("ID Format Invalid")
	}

	forumData.ID = *forumIDUUID

	return forumData, nil

}

func (repo *forumRepository) FetchForumWithUserReaction(offset int, limit int, userID int, topForumSort bool) ([]domain.Forum, error) {

	var err error
	var queryString string
	var isUpVoted, isDownVoted int

	forumList := make([]domain.Forum, 0)

	queryString = `
	SELECT 
		f.id,
		f.title,
		f.question, 
		f.author_user_id,
		ua.full_name,
		ua.username,
		f.status,
		f.category_id,
		c.category_name,
		IFNULL(fr.up_vote, 0) up_vote,
		IFNULL(fr.down_vote, 0) down_vote,
		IFNULL(frp.replies_count, 0) replies_count,
		IFNULL(ufr.up_vote, 0) is_up_vote,
		IFNULL(ufr.down_vote, 0) is_down_vote
	FROM 
		forum f
	LEFT JOIN user_auth ua ON ua.id = f.author_user_id 
	LEFT JOIN category c ON c.id = f.category_id 
	LEFT JOIN (
		SELECT fr.forum_id, SUM( fr.up_vote ) up_vote, SUM( fr.down_vote ) down_vote
		FROM forum_reaction fr 
		GROUP BY fr.forum_id 
	) fr ON fr.forum_id = f.id 
	LEFT JOIN (
		SELECT fr.forum_id, COUNT(1) replies_count FROM forum_replies fr GROUP BY fr.forum_id 
	) frp ON frp.forum_id = f.id 
	LEFT JOIN (
		SELECT fr.forum_id, SUM( fr.up_vote ) up_vote, SUM( fr.down_vote ) down_vote
		FROM forum_reaction fr 
		WHERE fr.user_id = ?
		GROUP BY fr.forum_id 
	) ufr ON ufr.forum_id = f.id 
	ORDER BY f.created_at DESC`

	if topForumSort {
		queryString += `, up_vote DESC, replies_count DESC`
	}

	queryString += "\nLIMIT ? OFFSET ?"

	forumQueryResult, err := repo.db.Query(queryString, userID, limit, offset)
	defer forumQueryResult.Close()

	if err != nil {
		return nil, err
	}

	for forumQueryResult.Next() {

		var forumID string
		forumData := domain.Forum{}

		err = forumQueryResult.Scan(
			&forumID,
			&forumData.Title,
			&forumData.Question,
			&forumData.AuthorUserID,
			&forumData.AuthoFullName,
			&forumData.AuthoUsername,
			&forumData.Status,
			&forumData.CategoryID,
			&forumData.CategoryName,
			&forumData.UpVote,
			&forumData.DownVote,
			&forumData.RepliesCount,
			&isUpVoted,
			&isDownVoted,
		)

		if err != nil {
			return nil, err
		}

		if isUpVoted > 0 {
			forumData.IsUpVoted = true
		}

		if isDownVoted > 0 {
			forumData.IsDownVoted = true
		}

		forumIDUUID, err := domain.NewUUIDFromString(forumID)

		if err != nil {
			return nil, fmt.Errorf("ID Format Invalid")
		}

		forumData.ID = *forumIDUUID

		forumList = append(forumList, forumData)

	}

	return forumList, nil

}

func (repo *forumRepository) FetchForumByAuthorIDWithUserReaction(authorID int, offset int, limit int, userID int, topForumSort bool) ([]domain.Forum, error) {

	var err error
	var queryString string
	var isUpVoted, isDownVoted int

	forumList := make([]domain.Forum, 0)

	queryString = `
	SELECT 
		f.id,
		f.title,
		f.question, 
		f.author_user_id,
		ua.full_name,
		ua.username,
		f.status,
		f.category_id,
		c.category_name,
		IFNULL(fr.up_vote, 0) up_vote,
		IFNULL(fr.down_vote, 0) down_vote,
		IFNULL(frp.replies_count, 0) replies_count,
		IFNULL(ufr.up_vote, 0) is_up_vote,
		IFNULL(ufr.down_vote, 0) is_down_vote
	FROM 
		forum f
	LEFT JOIN user_auth ua ON ua.id = f.author_user_id 
	LEFT JOIN category c ON c.id = f.category_id 
	LEFT JOIN (
		SELECT fr.forum_id, SUM( fr.up_vote ) up_vote, SUM( fr.down_vote ) down_vote
		FROM forum_reaction fr 
		GROUP BY fr.forum_id 
	) fr ON fr.forum_id = f.id 
	LEFT JOIN (
		SELECT fr.forum_id, COUNT(1) replies_count FROM forum_replies fr GROUP BY fr.forum_id 
	) frp ON frp.forum_id = f.id 
	LEFT JOIN (
		SELECT fr.forum_id, SUM( fr.up_vote ) up_vote, SUM( fr.down_vote ) down_vote
		FROM forum_reaction fr 
		WHERE fr.user_id = ?
		GROUP BY fr.forum_id 
	) ufr ON ufr.forum_id = f.id 
	WHERE f.author_user_id = ?
	ORDER BY f.created_at DESC`

	if topForumSort {
		queryString += `, up_vote DESC, replies_count DESC`
	}

	queryString += "\nLIMIT ? OFFSET ?"

	forumQueryResult, err := repo.db.Query(queryString, userID, authorID, limit, offset)
	defer forumQueryResult.Close()

	if err != nil {
		return nil, err
	}

	for forumQueryResult.Next() {

		var forumID string
		forumData := domain.Forum{}

		err = forumQueryResult.Scan(
			&forumID,
			&forumData.Title,
			&forumData.Question,
			&forumData.AuthorUserID,
			&forumData.AuthoFullName,
			&forumData.AuthoUsername,
			&forumData.Status,
			&forumData.CategoryID,
			&forumData.CategoryName,
			&forumData.UpVote,
			&forumData.DownVote,
			&forumData.RepliesCount,
			&isUpVoted,
			&isDownVoted,
		)

		if err != nil {
			return nil, err
		}

		if isUpVoted > 0 {
			forumData.IsUpVoted = true
		}

		if isDownVoted > 0 {
			forumData.IsDownVoted = true
		}

		forumIDUUID, err := domain.NewUUIDFromString(forumID)

		if err != nil {
			return nil, fmt.Errorf("ID Format Invalid")
		}

		forumData.ID = *forumIDUUID

		forumList = append(forumList, forumData)

	}

	return forumList, nil

}

func (repo *forumRepository) FetchForumByCategoryIDWithUserReaction(categoryID int, offset int, limit int, userID int, topForumSort bool) ([]domain.Forum, error) {

	var err error
	var queryString string
	var isUpVoted, isDownVoted int

	forumList := make([]domain.Forum, 0)

	queryString = `
	SELECT 
		f.id,
		f.title,
		f.question, 
		f.author_user_id,
		ua.full_name,
		ua.username,
		f.status,
		f.category_id,
		c.category_name,
		IFNULL(fr.up_vote, 0) up_vote,
		IFNULL(fr.down_vote, 0) down_vote,
		IFNULL(frp.replies_count, 0) replies_count,
		IFNULL(ufr.up_vote, 0) is_up_vote,
		IFNULL(ufr.down_vote, 0) is_down_vote
	FROM 
		forum f
	LEFT JOIN user_auth ua ON ua.id = f.author_user_id 
	LEFT JOIN category c ON c.id = f.category_id 
	LEFT JOIN (
		SELECT fr.forum_id, SUM( fr.up_vote ) up_vote, SUM( fr.down_vote ) down_vote
		FROM forum_reaction fr 
		GROUP BY fr.forum_id 
	) fr ON fr.forum_id = f.id 
	LEFT JOIN (
		SELECT fr.forum_id, COUNT(1) replies_count FROM forum_replies fr GROUP BY fr.forum_id 
	) frp ON frp.forum_id = f.id 
	LEFT JOIN (
		SELECT fr.forum_id, SUM( fr.up_vote ) up_vote, SUM( fr.down_vote ) down_vote
		FROM forum_reaction fr 
		WHERE fr.user_id = ?
		GROUP BY fr.forum_id 
	) ufr ON ufr.forum_id = f.id 
	WHERE f.category_id = ?
	ORDER BY f.created_at DESC`

	if topForumSort {
		queryString += `, up_vote DESC, replies_count DESC`
	}

	queryString += "\nLIMIT ? OFFSET ?"

	forumQueryResult, err := repo.db.Query(queryString, userID, categoryID, limit, offset)
	defer forumQueryResult.Close()

	if err != nil {
		return nil, err
	}

	for forumQueryResult.Next() {

		var forumID string
		forumData := domain.Forum{}

		err = forumQueryResult.Scan(
			&forumID,
			&forumData.Title,
			&forumData.Question,
			&forumData.AuthorUserID,
			&forumData.AuthoFullName,
			&forumData.AuthoUsername,
			&forumData.Status,
			&forumData.CategoryID,
			&forumData.CategoryName,
			&forumData.UpVote,
			&forumData.DownVote,
			&forumData.RepliesCount,
			&isUpVoted,
			&isDownVoted,
		)

		if err != nil {
			return nil, err
		}

		if isUpVoted > 0 {
			forumData.IsUpVoted = true
		}

		if isDownVoted > 0 {
			forumData.IsDownVoted = true
		}

		forumIDUUID, err := domain.NewUUIDFromString(forumID)

		if err != nil {
			return nil, fmt.Errorf("ID Format Invalid")
		}

		forumData.ID = *forumIDUUID

		forumList = append(forumList, forumData)

	}

	return forumList, nil

}

func (repo *forumRepository) SearchByTitleAndQuestionWithUserReaction(offset int, limit int, userID int, query string) ([]domain.Forum, error) {

	var err error
	var queryString string
	var isUpVoted, isDownVoted int

	forumList := make([]domain.Forum, 0)

	queryString = `
		SELECT 
			f.id,
			f.title,
			f.question, 
			f.author_user_id,
			ua.full_name,
			ua.username,
			f.status,
			f.category_id,
			c.category_name,
			IFNULL(fr.up_vote, 0) up_vote,
			IFNULL(fr.down_vote, 0) down_vote,
			IFNULL(frp.replies_count, 0) replies_count,
			IFNULL(ufr.up_vote, 0) is_up_vote,
			IFNULL(ufr.down_vote, 0) is_down_vote
		FROM 
			forum f
		LEFT JOIN user_auth ua ON ua.id = f.author_user_id 
		LEFT JOIN category c ON c.id = f.category_id 
		LEFT JOIN (
			SELECT fr.forum_id, SUM( fr.up_vote ) up_vote, SUM( fr.down_vote ) down_vote
			FROM forum_reaction fr 
			GROUP BY fr.forum_id 
		) fr ON fr.forum_id = f.id 
		LEFT JOIN (
			SELECT fr.forum_id, COUNT(1) replies_count FROM forum_replies fr GROUP BY fr.forum_id 
		) frp ON frp.forum_id = f.id 
		LEFT JOIN (
			SELECT fr.forum_id, SUM( fr.up_vote ) up_vote, SUM( fr.down_vote ) down_vote
			FROM forum_reaction fr 
			WHERE fr.user_id = ?
			GROUP BY fr.forum_id 
		) ufr ON ufr.forum_id = f.id 
		WHERE 
			f.title LIKE '%?%' or f.question LIKE '%?%'
		ORDER BY up_vote DESC, replies_count DESC, f.created_at DESC
		LIMIT ? OFFSET ?
	`

	forumQueryResult, err := repo.db.Query(queryString, userID, query, query, limit, offset)
	defer forumQueryResult.Close()

	if err != nil {
		return nil, err
	}

	for forumQueryResult.Next() {

		var forumID string
		forumData := domain.Forum{}

		err = forumQueryResult.Scan(
			&forumID,
			&forumData.Title,
			&forumData.Question,
			&forumData.AuthorUserID,
			&forumData.AuthoFullName,
			&forumData.AuthoUsername,
			&forumData.Status,
			&forumData.CategoryID,
			&forumData.CategoryName,
			&forumData.UpVote,
			&forumData.DownVote,
			&forumData.RepliesCount,
			&isDownVoted,
			&isDownVoted,
		)

		if err != nil {
			return nil, err
		}

		if isUpVoted > 0 {
			forumData.IsUpVoted = true
		}

		if isDownVoted > 0 {
			forumData.IsDownVoted = true
		}

		forumIDUUID, err := domain.NewUUIDFromString(forumID)

		if err != nil {
			return nil, fmt.Errorf("ID Format Invalid")
		}

		forumData.ID = *forumIDUUID

		forumList = append(forumList, forumData)

	}

	return forumList, nil

}

func (repo *forumRepository) DeleteForumByID(id domain.UUID) (int, error) {

	var err error
	var queryString string
	var rowAffected int64
	var statement *sql.Stmt
	var res sql.Result

	queryString = `DELETE FROM forum WHERE id=?`

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
		return -1, fmt.Errorf("Forum With Specific ID Not Found")
	}

	return int(rowAffected), nil

}

func (repo *forumRepository) UpdateForumByID(id domain.UUID, title string, question sql.NullString, categoryID int) (int, error) {

	var err error
	var queryString string
	var rowAffected int64
	var statement *sql.Stmt
	var res sql.Result

	queryString = `UPDATE forum SET title=?, question=?, category_id=?, updated_at=NOW() WHERE id=?`

	statement, err = repo.db.Prepare(queryString)

	if err != nil {
		return -1, err
	}

	res, err = statement.Exec(title, question.String, categoryID, id.GetValue())

	if err != nil {
		return -1, err
	}

	rowAffected, err = res.RowsAffected()

	if err != nil {
		return -1, err
	}

	if rowAffected == 0 {
		return -1, fmt.Errorf("Forum With Specific ID Not Found")
	}

	return int(rowAffected), nil

}

func (repo *forumRepository) UpdateForumStatusByID(id domain.UUID, status string) (int, error) {

	var err error
	var queryString string
	var rowAffected int64
	var statement *sql.Stmt
	var res sql.Result

	queryString = `UPDATE forum SET status=?, updated_at=NOW() WHERE id=?`

	statement, err = repo.db.Prepare(queryString)

	if err != nil {
		return -1, err
	}

	res, err = statement.Exec(status, id.GetValue())

	if err != nil {
		return -1, err
	}

	rowAffected, err = res.RowsAffected()

	if err != nil {
		return -1, err
	}

	if rowAffected == 0 {
		return -1, fmt.Errorf("Forum With Specific ID Not Found")
	}

	return int(rowAffected), nil

}
