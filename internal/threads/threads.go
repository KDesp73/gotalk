package threads

import (
	"gotalk/internal/comments"
	"gotalk/internal/utils"
	"strings"
)


type Thread struct {
	ID string
	Comments []*comments.Comment
}

func (t *Thread) PushComment(author string, content string) {
	comment := &comments.Comment{
		Author: author,
		Content: content,
		ThreadID: t.ID,
		Timestamp: utils.CurrentTimestamp(),
	}
	comment.GenerateID(len(t.Comments))
	t.Comments = append(t.Comments, comment)
}

func (t *Thread) SearchCommentAuthor(author string) []int {
	var indices []int
	for i, comment := range t.Comments {
		if strings.EqualFold(comment.Author, author) {
			indices = append(indices, i)
		}
	}
	return indices
}


// SearchCommentID searches for a comment by 
// the first 7 characters of its ID and 
// returns its index. Returns -1 if not found.
func (t *Thread) SearchCommentID(id string) int {
	if len(id) < 7 {
		return -1
	}
	for i, comment := range t.Comments {
		if len(comment.ID) >= 7 && comment.ID[:7] == id[:7] { // Compare the first 7 characters
			return i
		}
	}
	return -1
}

// SearchCommentContent searches for a comment 
// by its content and returns its index. Returns -1 if not found.
func (t *Thread) SearchCommentContent(content string) int {
	for i, comment := range t.Comments {
		if strings.Contains(comment.Content, content) { 
			return i
		}
	}
	return -1
}

func (t *Thread) RemoveComment(index int) bool {
	if index < 0 {
		return false
	}

	t.Comments = append(t.Comments[:index], t.Comments[index+1:]...)

	return true
}

