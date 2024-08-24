package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"gotalk/api/state"
	"gotalk/api/v1/middleware"
	"gotalk/api/v1/routing"
	"gotalk/internal/encryption"
	"gotalk/internal/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const StateFile = "gotalk.state"

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

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	var port = 8080
	var generate_key = false
	flag.IntVar(&port, "port", port, "Specify the port")
	flag.BoolVar(&generate_key, "generate-key", generate_key, "Generate an encryption key")
	flag.Parse()

	if generate_key {
		key, err := encryption.GenerateKey(state.KeySize)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERRO: Could not generate key\n")
			os.Exit(1)
		}
		fmt.Printf("Key: %s\n", hex.EncodeToString(key))
		os.Exit(0)
	}

	decodedKey, err := getKey()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERRO: %s\n", err.Error())
		os.Exit(1)
	}
	
	if utils.FileExists(StateFile) {
		s, err := state.LoadState(StateFile, decodedKey)

		if err != nil {
			fmt.Fprintf(os.Stderr, "ERRO: Could not load state from file\n")
			state.Instance = state.StateInit()
		} else {
			fmt.Println("INFO: Loading state...")
			state.Instance = s
		}
	} else {
		state.Instance = state.StateInit()
	}


	router := routing.Router()
	adminRouter := routing.AdminRouter()

	router.Handle("/", middleware.EnsureAdmin(adminRouter))

	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))

	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.IsAuthenticated,
	)

	server := http.Server {
		Addr: fmt.Sprintf(":%d", port),
		Handler: stack(router),
	}

	log.Printf("Starting server on port %d", port)
	go server.ListenAndServe()

	<-sigChan
	println()

	fmt.Println("INFO: Saving state...")
	err = state.SaveState(state.Instance, StateFile, decodedKey)

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERRO: Could not save the state (%v)\n", err);
	}

	fmt.Println("INFO: Terminating...")
}
