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

func (p *ThreadPool) PushThread(thread *Thread) {
	var id uuid.UUID = uuid.New()
	for{
		if !p.idExists(id.String()) {
			break
		}
		id = uuid.New()
	}

	thread.ID = id.String()
	p.idHashMap[id.String()] = len(p.Threads) // each key points to the thread's index
	p.Threads = append(p.Threads, thread)
}
