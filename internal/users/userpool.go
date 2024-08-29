package users

import (
	"github.com/google/uuid"
)

type UserPool struct {
	Items []*User
	IdHashMap map[string]int
}

func PoolInit() UserPool {
	pool := UserPool {
		IdHashMap: make(map[string]int),
	}

	return pool
}

func (p UserPool) idExists(id string) bool {
	_, ok := p.IdHashMap[id]
	return ok
}


func (p *UserPool) NameExists(name string) bool {
	for _, user := range p.Items {
		if(name == user.Name) {
			return true
		}
	}

	return false
}

func (p *UserPool) EmailExists(email string) bool {
	for _, user := range p.Items {
		if(email == user.Email) {
			return true
		}
	}

	return false
}

func (p *UserPool) Get(id string) *User {
	return p.Items[p.IdHashMap[id]]
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

	p.IdHashMap[id] = len(p.Items) // each key points to the user's index
	user.Key = id 
	p.Items = append(p.Items, user)

	return id
}

func (p *UserPool) IsAdmin(id string) bool {
	return p.Get(id).Type == ADMIN
}

func (p *UserPool) Sudo(id string, undo bool) bool {
	index, exists := p.IdHashMap[id]

	if !exists {
		return false
	}

	if undo {
		p.Items[index].Type = DEFAULT
	} else {
		p.Items[index].Type = ADMIN
	}
	return true
}

func (p *UserPool) RemoveUser(id string) bool {
	index, exists := p.IdHashMap[id]
	if !exists {
		return false
	}

	p.Items = append(p.Items[:index], p.Items[index+1:]...)

	delete(p.IdHashMap, id)

	for i := index; i < len(p.Items); i++ {
		p.IdHashMap[p.Items[i].Key] = i
	}

	return true
}
