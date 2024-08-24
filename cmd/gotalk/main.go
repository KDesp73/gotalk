package main

import (
	"encoding/hex"
	"fmt"
	"gotalk/api/state"
	"gotalk/api/v1/middleware"
	"gotalk/api/v1/routing"
	"gotalk/internal/options"
	"gotalk/internal/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func getKey() ([]byte, error) {
	key := os.Getenv("ENCR_KEY")
	if utils.StrEmpty(key) {
		return nil, fmt.Errorf("ENCR_KEY environment variable is not set")
	}

	decodedKey, err := hex.DecodeString(key)
	if err != nil {
		return nil, fmt.Errorf("Could not decode key")
	}

	return decodedKey, nil
}

func save(s *state.State, key []byte) {
	fmt.Println("INFO: Saving state...")
	err := state.SaveState(s, state.StateFile, key)

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERRO: Could not save the state (%v)\n", err);
	}
}

func backup(s *state.State, key []byte) {
	fmt.Println("INFO: Backing up state...")
	err := state.SaveState(s, state.StateFileBackup, key)

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERRO: Could not save the state (%v)\n", err);
	}
}


func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	saveTicker := time.NewTicker(state.SaveInterval)
	defer saveTicker.Stop()
	backupTicker := time.NewTicker(state.BackupInterval)
	defer backupTicker.Stop()

	options := options.ParseOptions()
	options.HandleOptions()

	// decodedKey, err := getKey()
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "ERRO: %s\n", err.Error())
	// 	os.Exit(1)
	// }
	var decodedKey []byte = nil
	
	if utils.FileExists(state.StateFile) {
		s, err := state.LoadState(state.StateFile, decodedKey)

		if err != nil {
			fmt.Fprintf(os.Stderr, "ERRO: Could not load state from file\n")
			utils.CopyFile(state.StateFile, state.StateFileOld)
			state.Instance = state.StateInit()
		} else {
			fmt.Println("INFO: Loading state...")
			state.Instance = s
		}
	} else {
		state.Instance = state.StateInit()
	}



	router := routing.Router()
	stack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server {
		Addr: fmt.Sprintf(":%d", options.Port),
		Handler: stack(router),
	}

	log.Printf("Starting server on port %d", options.Port)
	
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case _ = <-saveTicker.C:
				save(state.Instance, decodedKey)
			case _ = <-backupTicker.C:
				backup(state.Instance, decodedKey)
			}
		}
	}()

	go server.ListenAndServe()


	<-sigChan
	println()

	save(state.Instance, decodedKey)

	fmt.Println("INFO: Terminating...")
}
