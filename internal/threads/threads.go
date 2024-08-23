package threads

import "gotalk/internal/comments"


type Thread struct {
	ID string
	Comments []*comments.Comment
}



func (t *Thread) AppendComment(author string, content string) {
	comment := &comments.Comment{
		Author: author,
		Content: content,
		ThreadID: t.ID,
	}
	comment.GenerateID(len(t.Comments))
	t.Comments = append(t.Comments, comment)
}
