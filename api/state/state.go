package state

import (
	"bytes"
	"encoding/gob"
	_ "gotalk/internal/encryption"
	"gotalk/internal/threads"
	"gotalk/internal/users"
	"os"
)

type State struct {
	Threads threads.ThreadPool
	Users users.UserPool
}

var Instance *State
var KeySize = 16

func StateInit() *State {
	return &State{
		Threads: threads.PoolInit(),
		Users: users.PoolInit(),
	}
}

func LoadState(filename string, key []byte) (*State, error) {
	var state *State = StateInit()

	data, err := os.ReadFile(filename)
	if err != nil {
		return state, err
	}

	// decryptedData, err := encryption.Decrypt(string(data), key)
	// if err != nil {
	// 	return state, err
	// }

	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)
	if err := decoder.Decode(&state); err != nil {
		return state, err
	}

	return state, nil
}

func SaveState(state *State, filename string, key []byte) error {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(state); err != nil {
		return err
	}

	// encryptedData, err := encryption.Encrypt(buf.Bytes(), key)
	// if err != nil {
	// 	return err
	// }

	return os.WriteFile(filename, buf.Bytes(), 0644)
}
