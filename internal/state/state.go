package state

import (
	"bytes"
	"encoding/gob"
	_ "gotalk/internal/encryption"
	"gotalk/internal/threads"
	"gotalk/internal/users"
	"os"
	"time"
)

type State struct {
	Threads threads.ThreadPool `json:"threads"`
	Users users.UserPool `json:"users"`
}

var Instance *State
const KeySize = 16
const SaveInterval = 2 * time.Minute
const BackupInterval = 10 * time.Minute
const StateFile = "gotalk.gob"
const StateFileOld = "gotalk.gob.old"
const StateFileBackup = "gotalk.gob.backup"
const StateFileTest = "gotalk.gob.test"

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

	// Fill hashmaps
	for i, user := range state.Users.Items {
		state.Users.IdHashMap[user.Key] = i
	}

	for i, thread := range state.Threads.Items {
		state.Threads.IdHashMap[thread.ID] = i
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
