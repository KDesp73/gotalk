package users

import (
	"github.com/google/uuid"
)

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


func (p *UserPool) NameExists(name string) bool {
	for _, user := range p.Users {
		if(name == user.Name) {
			return true
		}
	}

	return false
}

func (p *UserPool) EmailExists(email string) bool {
	for _, user := range p.Users {
		if(email == user.Email) {
			return true
		}
	}

	return false
}

func (p *UserPool) Get(id string) *User {
	return p.Users[p.idHashMap[id]]
}

func (p *UserPool) GenerateId() string {
	var id uuid.UUID = uuid.New()
	for{
		if !p.idExists(id.String()) {
			break
		}
		id = uuid.New()
	}
	return id.String()
}

func (p *UserPool) PushUser(user *User) string {
	id := p.GenerateId()

	p.idHashMap[id] = len(p.Users) // each key points to the user's index
	user.Key = id 
	p.Users = append(p.Users, user)

	return id
}

func (p *UserPool) IsAdmin(id string) bool {
	return p.Get(id).Type == ADMIN
}

func (p *UserPool) Sudo(id string, undo bool) bool {
	index, exists := p.idHashMap[id]

	if !exists {
		return false
	}

	if undo {
		p.Users[index].Type = DEFAULT
	} else {
		p.Users[index].Type = ADMIN
	}
	return true
}
