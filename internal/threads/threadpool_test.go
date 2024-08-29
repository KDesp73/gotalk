package threads

import (
	"testing"
)

func TestRemoveThread(t * testing.T) {
	pool := PoolInit()

	id := pool.PushThread(&Thread{})

	if !pool.RemoveThread(id) {
		t.Fatalf("Failed to remove existing thread")
	}
}

func TestRemoveThreadSecond(t * testing.T) {
	pool := PoolInit()

	id := pool.PushThread(&Thread{})
	id = pool.PushThread(&Thread{})

	if !pool.RemoveThread(id) {
		t.Fatalf("Failed to remove existing thread")
	}
}
