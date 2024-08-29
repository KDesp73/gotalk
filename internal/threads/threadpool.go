package threads

import (
	"github.com/google/uuid"
)

type ThreadPool struct {
	Threads []*Thread
	idHashMap map[string]int
}

func PoolInit() ThreadPool {
	pool := ThreadPool {
		idHashMap: make(map[string]int),
	}

	return pool
}

func (p ThreadPool) idExists(id string) bool {
	_, ok := p.idHashMap[id]
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
	index, exists := p.idHashMap[id]
	if !exists {
		return false
	}

	p.Threads = append(p.Threads[:index], p.Threads[index+1:]...)

	delete(p.idHashMap, id)

	for i := index; i < len(p.Threads); i++ {
		p.idHashMap[p.Threads[i].ID] = i
	}

	return true
}

func (p *ThreadPool) PushThread(thread *Thread) string {
	id := p.GenerateId()

	thread.ID = id
	p.idHashMap[id] = len(p.Threads) // each key points to the thread's index
	p.Threads = append(p.Threads, thread)

	return id
}

func (p *ThreadPool) Get(id string) *Thread {
	index := p.idHashMap[id]
	if index < 0 || index > len(p.Threads) - 1 {
		return nil
	}

	return p.Threads[index]
}

// SearchCommentID searches for a comment by 
// the first 7 characters of its ID and 
// returns its index. Returns -1 if not found.
func (t *ThreadPool) SearchCommentID(id string) int {
	if len(id) < 7 {
		return -1
	}
	for _, thread := range t.Threads {
		for i, comment := range thread.Comments {
			if len(comment.ID) >= 7 && comment.ID[:7] == id[:7] { // Compare the first 7 characters
				return i
			}
		}
	}
	return -1
}
