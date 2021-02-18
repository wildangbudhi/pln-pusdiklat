package domain

import (
	"io"
	"net/http"
)

type ForumRepository interface {
	FetchCategory(requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	CreateForum(requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	UpdateForum(forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	CloseForum(forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	DeleteForum(forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	FetchWithPagination(requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	GetForum(forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	FetchWithPaginationByAuthorID(authorID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	SearchForum(requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	ReactForum(forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	ReplyForum(forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	UpdateForumReplies(forumReplyID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	DeleteForumReplies(forumReplyID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	ReactForumReplies(forumReplyID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	FetchReplyByForumIDWithPagination(forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
}

type ForumUsecase interface {
	FetchCategory(requestURI string, requestMethod string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	CreateForum(requestURI string, requestMethod string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	UpdateForum(requestURI string, requestMethod string, forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	CloseForum(requestURI string, requestMethod string, forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	DeleteForum(requestURI string, requestMethod string, forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	FetchWithPagination(requestURI string, requestMethod string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	GetForum(requestURI string, requestMethod string, forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	FetchWithPaginationByAuthorID(requestURI string, requestMethod string, authorID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	SearchForum(requestURI string, requestMethod string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	ReactForum(requestURI string, requestMethod string, forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	ReplyForum(requestURI string, requestMethod string, forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	UpdateForumReplies(requestURI string, requestMethod string, forumReplyID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	DeleteForumReplies(requestURI string, requestMethod string, forumReplyID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	ReactForumReplies(requestURI string, requestMethod string, forumReplyID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
	FetchReplyByForumIDWithPagination(requestURI string, requestMethod string, forumID string, requestBody io.Reader, urlQuery string, requestHeader map[string][]string) (*http.Response, error)
}
