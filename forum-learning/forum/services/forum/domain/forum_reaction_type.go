package domain

import "fmt"

// ForumReactionType is an Object to handle reaction of Forum also format checking
type ForumReactionType struct {
	isUpVoteToggled   bool
	isDownVoteToggled bool
}

// NewForumReactionType is an Constructor for ForumReactionType Object
func NewForumReactionType(reactionType string, toggled bool) (*ForumReactionType, error) {

	reactionObj := &ForumReactionType{
		isUpVoteToggled:   false,
		isDownVoteToggled: false,
	}

	if reactionType == "up_vote" {
		reactionObj.isUpVoteToggled = toggled
	} else if reactionType == "down_vote" {
		reactionObj.isDownVoteToggled = toggled
	} else {
		return nil, fmt.Errorf("Forum Reaction Type Not Allowed")
	}

	return reactionObj, nil

}

// IsUpVoteToggled is a Getter Function to Get if UpVote Toggled
func (obj *ForumReactionType) IsUpVoteToggled() bool {
	return obj.isUpVoteToggled
}

// IsDownVoteToggled is a Getter Function to Get if DownVote Toggled
func (obj *ForumReactionType) IsDownVoteToggled() bool {
	return obj.isDownVoteToggled
}

// IsAllUntoggled is a Function to Check if All Untoggled
// func (obj *ForumReactionType) IsAllUntoggled() bool {
// 	return obj.isUpVoteToggled && obj.isDownVoteToggled
// }
