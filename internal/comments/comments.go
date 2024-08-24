package comments

import (
	"fmt"
	"gotalk/internal/encryption"
	"gotalk/internal/utils"
	"strconv"
)

type Comment struct {
	ID string
	Author string
	Content string
	ThreadID string
	Timestamp string
}

func CreateComment(author string, content string) *Comment {
	return &Comment{
		Author: author,
		Content: content,
		Timestamp: utils.CurrentTimestamp(),
	}
}

func (c *Comment) Log() {
	fmt.Printf("ID: %s, Author: %s, Content: %s, ThreadID: %s\n", c.ID, c.Author, c.Content, c.ThreadID)
}

func (c *Comment) GenerateID(index int) string {
	id := encryption.Hash(c.Author + c.Content + c.Timestamp + strconv.Itoa(index))
	c.ID = id
	return id
}
