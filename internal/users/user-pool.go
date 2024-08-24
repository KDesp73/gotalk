package users

import (
	"gotalk/internal/encryption"

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
	hashedId := encryption.Hash(id)

	p.idHashMap[hashedId] = len(p.Users) // each key points to the user's index
	user.Key = hashedId
	p.Users = append(p.Users, user)

	return id
}
