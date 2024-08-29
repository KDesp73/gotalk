package comments

import (
	"fmt"
	"gotalk/internal/encryption"
	"gotalk/internal/utils"
	"strconv"
)

type Comment struct {
	ID string `json:"id"`
	Author string `json:"author"`
	Content string `json:"content"`
	ThreadID string `json:"threadid"`
	Timestamp string `json:"timestamp"`
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
