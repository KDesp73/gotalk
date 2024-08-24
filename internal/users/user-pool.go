package users

import "github.com/google/uuid"

type UserPool struct {
	Users []*User
	idHashMap map[string]int
}

func PoolInit() UserPool {
	pool := UserPool {
		idHashMap: make(map[string]int),
	}

	return pool
}

func (p UserPool) idExists(id string) bool {
	_, ok := p.idHashMap[id]
	return ok
}

func (p *UserPool) PushUser(thread *User) {
	var id uuid.UUID = uuid.New()
	for{
		if !p.idExists(id.String()) {
			break
		}
		id = uuid.New()
	}

	p.idHashMap[id.String()] = len(p.Users) // each key points to the user's index
	p.Users = append(p.Users, thread)
}
