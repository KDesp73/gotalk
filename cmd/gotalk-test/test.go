package main

import (
	"fmt"
	"gotalk/internal/threads"
)

func main() {
	pool := threads.PoolInit()

	thread := &threads.Thread{}
	pool.PushThread(thread)

	thread.AppendComment("KDesp73", "Hello")
	thread.AppendComment("Christina", "Hi")


	thread1 := &threads.Thread{}
	pool.PushThread(thread1)
	thread1.AppendComment("KDesp73", "How are you in a different thread")

	fmt.Printf("Pool length: %d\n", len(pool.Threads))

	for _, comment := range pool.Threads[0].Comments {
		comment.Log()
	}

	for _, comment := range pool.Threads[1].Comments {
		comment.Log()
	}
}
