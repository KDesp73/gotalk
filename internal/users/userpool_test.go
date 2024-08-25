package users

import (
	"gotalk/internal/encryption"
	"gotalk/internal/utils"
	"testing"
)

func TestPushUser(t *testing.T) {
	pool := PoolInit()

	key := pool.PushUser(&User{})
	key = encryption.Hash(key)

	if len(pool.Users) != 1 {
		t.Fatalf("1. One user added and yet %d users found", len(pool.Users))
	}

	if pool.idHashMap[key] != 0 {
		t.Fatalf("1. Id's value should be 0");
	}

	foundKey := pool.Users[pool.idHashMap[key]].Key
	if foundKey != key {
		t.Fatalf("1. Invalid user stored. Expected key: %s | Found key: %s\n", key, foundKey)
	}

	key1 := pool.PushUser(&User{})
	key1 = encryption.Hash(key1)

	if len(pool.Users) != 2 {
		t.Fatalf("2. Two user added and yet %d users found", len(pool.Users))
	}

	if pool.idHashMap[key1] != 1 {
		t.Fatalf("2. Id's value should be 1");
	}

	foundKey1 := pool.Users[pool.idHashMap[key1]].Key
	if foundKey1 != key1 {
		t.Fatalf("2. Invalid user stored. Expected key: %s | Found key: %s\n", key1, foundKey1)
	}
}

func TestGet(t *testing.T) {
	pool := PoolInit()

	name := "test"
	email := "test@test.com"
	time := utils.CurrentTimestamp()
	typ := USER_ADMIN
	key := pool.PushUser(&User{
		Name: name,
		Email: email,
		Type: typ,
		SignUpTime: time,
	})

	encrKey := encryption.Hash(key)
	user := pool.Get(encrKey)

	if user.Name != name || user.Email != email || user.SignUpTime != time || user.Type != typ {
		t.Fatalf("User is not correct")
	}
}

func TestIsAdminNo(t *testing.T) {
	pool := PoolInit()

	key := pool.PushUser(&User{
		Type: USER_DEFAULT,
	})


	if pool.IsAdmin(encryption.Hash(key)) {
		t.Fatalf("User is not an admin")
	}
}

func TestIsAdminYes(t *testing.T) {
	pool := PoolInit()

	key := pool.PushUser(&User{
		Type: USER_ADMIN,
	})

	if !pool.IsAdmin(encryption.Hash(key)) {
		t.Fatalf("User should be an admin")
	}
}

func TestSudo(t *testing.T) {
	pool := PoolInit()

	key := pool.PushUser(&User{
		Type: USER_DEFAULT,
	})
	key = encryption.Hash(key)
	
	pool.Sudo(key, false)

	user := pool.Get(key)

	if user.Type != USER_ADMIN {
		t.Fatalf("User did not become an admin")
	}
}


func TestUnSudo(t *testing.T) {
	pool := PoolInit()

	key := pool.PushUser(&User{
		Type: USER_ADMIN,
	})
	key = encryption.Hash(key)
	
	pool.Sudo(key, true)

	user := pool.Get(key)

	if user.Type != USER_DEFAULT {
		t.Fatalf("User did not become a default")
	}
}
