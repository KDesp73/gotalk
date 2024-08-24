package state

import (
	"bufio"
	"gotalk/internal/threads"
	"gotalk/internal/users"
	"os"
)

type State struct {
	Threads threads.ThreadPool
	Users users.UserPool
}
var Instance *State

func StateInit() *State {
	return &State{
		Threads: threads.PoolInit(),
		Users: users.PoolInit(),
	}
}

func (s *State) SaveState(file string) error {
	outFile, err := os.Create(file)
	if err != nil {
		return err
    }
    defer outFile.Close()

	writer := bufio.NewWriter(outFile)
    defer writer.Flush()

	return nil
}
