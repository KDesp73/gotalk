package threads

import (
	"github.com/google/uuid"
)

type ThreadPool struct {
	Items []*Thread
	IdHashMap map[string]int
}

func PoolInit() ThreadPool {
	pool := ThreadPool {
		IdHashMap: make(map[string]int),
	}

	return pool
}


func (p *ThreadPool) TitleExists(title string) bool {
	for _, thread := range p.Items {
		if title == thread.Title {
			return true
		}
	}
	return false
}

func (p ThreadPool) idExists(id string) bool {
	_, ok := p.IdHashMap[id]
	return ok
}

func (p *ThreadPool) GenerateId() string {
	var id uuid.UUID = uuid.New()
	for{
		if !p.idExists(id.String()) {
			break
		}
		id = uuid.New()
	}
	return id.String()
}

func (p *ThreadPool) RemoveThread(id string) bool {
	index, exists := p.IdHashMap[id]
	if !exists {
		return false
	}

	p.Items = append(p.Items[:index], p.Items[index+1:]...)

	delete(p.IdHashMap, id)

	for i := index; i < len(p.Items); i++ {
		p.IdHashMap[p.Items[i].ID] = i
	}

	return true
}

func (p *ThreadPool) PushThread(thread *Thread) string {
	id := p.GenerateId()

	thread.ID = id
	p.IdHashMap[id] = len(p.Items) // each key points to the thread's index
	p.Items = append(p.Items, thread)

	return id
}

func (p *ThreadPool) Get(id string) *Thread {
	index := p.IdHashMap[id]
	if index < 0 || index > len(p.Items) - 1 {
		return nil
	}

	return p.Items[index]
}

// SearchCommentID searches for a comment by 
// the first 7 characters of its ID and 
// returns its index. Returns -1 if not found.
func (t *ThreadPool) SearchCommentID(id string) int {
	if len(id) < 7 {
		return -1
	}
	for _, thread := range t.Items {
		for i, comment := range thread.Comments {
			if len(comment.ID) >= 7 && comment.ID[:7] == id[:7] { // Compare the first 7 characters
				return i
			}
		}
	}
	return -1
}
