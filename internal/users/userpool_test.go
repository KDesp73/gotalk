package users

import (
	"gotalk/internal/utils"
	"testing"
)

func TestPushUser(t *testing.T) {
	pool := PoolInit()

	key := pool.PushUser(&User{})

	if len(pool.Items) != 1 {
		t.Fatalf("1. One user added and yet %d users found", len(pool.Items))
	}

	if pool.IdHashMap[key] != 0 {
		t.Fatalf("1. Id's value should be 0");
	}

	foundKey := pool.Items[pool.IdHashMap[key]].Key
	if foundKey != key {
		t.Fatalf("1. Invalid user stored. Expected key: %s | Found key: %s\n", key, foundKey)
	}

	key1 := pool.PushUser(&User{})

	if len(pool.Items) != 2 {
		t.Fatalf("2. Two user added and yet %d users found", len(pool.Items))
	}

	if pool.IdHashMap[key1] != 1 {
		t.Fatalf("2. Id's value should be 1");
	}

	foundKey1 := pool.Items[pool.IdHashMap[key1]].Key
	if foundKey1 != key1 {
		t.Fatalf("2. Invalid user stored. Expected key: %s | Found key: %s\n", key1, foundKey1)
	}
}

func TestGet(t *testing.T) {
	pool := PoolInit()

	name := "test"
	email := "test@test.com"
	time := utils.CurrentTimestamp()
	typ := ADMIN
	key := pool.PushUser(&User{
		Name: name,
		Email: email,
		Type: typ,
		SignUpTime: time,
	})

	user := pool.Get(key)

	if user.Name != name || user.Email != email || user.SignUpTime != time || user.Type != typ {
		t.Fatalf("User is not correct")
	}
}

func TestIsAdminNo(t *testing.T) {
	pool := PoolInit()

	key := pool.PushUser(&User{
		Type: DEFAULT,
	})


	if pool.IsAdmin(key) {
		t.Fatalf("User is not an admin")
	}
}

func TestIsAdminYes(t *testing.T) {
	pool := PoolInit()

	key := pool.PushUser(&User{
		Type: ADMIN,
	})

	if !pool.IsAdmin(key) {
		t.Fatalf("User should be an admin")
	}
}

func TestSudo(t *testing.T) {
	pool := PoolInit()

	key := pool.PushUser(&User{
		Type: DEFAULT,
	})
	
	pool.Sudo(key, false)

	user := pool.Get(key)

	if user.Type != ADMIN {
		t.Fatalf("User did not become an admin")
	}
}


func TestUnSudo(t *testing.T) {
	pool := PoolInit()

	key := pool.PushUser(&User{
		Type: ADMIN,
	})
	
	pool.Sudo(key, true)

	user := pool.Get(key)

	if user.Type != DEFAULT {
		t.Fatalf("User did not become a default")
	}
}
