package domain

import "fmt"

// ForumRepliesReactionType is an Object to handle reaction of Forum Replies also format checking
type ForumRepliesReactionType struct {
	isUpVoteToggled   bool
	isDownVoteToggled bool
	isAgreeToggled    bool
	isSkepticToggled  bool
}

// NewForumRepliesReactionType is an Constructor for ForumRepliesReactionType Object
func NewForumRepliesReactionType(reactionType string, toggled bool) (*ForumRepliesReactionType, error) {

	reactionObj := &ForumRepliesReactionType{
		isUpVoteToggled:   false,
		isDownVoteToggled: false,
	}

	if reactionType == "up_vote" {
		reactionObj.isUpVoteToggled = true
	} else if reactionType == "down_vote" {
		reactionObj.isDownVoteToggled = true
	} else if reactionType == "agree" {
		reactionObj.isAgreeToggled = true
	} else if reactionType == "skeptic" {
		reactionObj.isSkepticToggled = true
	} else {
		return nil, fmt.Errorf("Forum Reaction Type Not Allowed")
	}

	return reactionObj, nil

}

// IsUpVoteToggled it a Getter Function to Get if UpVote Toggled
func (obj *ForumRepliesReactionType) IsUpVoteToggled() bool {
	return obj.isUpVoteToggled
}

// IsDownVoteToggled it a Getter Function to Get if DownVote Toggled
func (obj *ForumRepliesReactionType) IsDownVoteToggled() bool {
	return obj.isDownVoteToggled
}

// IsAgreeToggled it a Getter Function to Get if Agree Toggled
func (obj *ForumRepliesReactionType) IsAgreeToggled() bool {
	return obj.isAgreeToggled
}

// IsSkepticToggled it a Getter Function to Get if Skeptic Toggled
func (obj *ForumRepliesReactionType) IsSkepticToggled() bool {
	return obj.isSkepticToggled
}

// IsAllUntoggled is a Function to Check if All Untoggled
func (obj *ForumRepliesReactionType) IsAllUntoggled() bool {
	return obj.isUpVoteToggled && obj.isDownVoteToggled && obj.isAgreeToggled && obj.isSkepticToggled
}
