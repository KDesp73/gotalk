package options

import (
	"encoding/hex"
	"flag"
	"fmt"
	"gotalk/internal/state"
	"gotalk/internal/encryption"
	"os"
)

type Options struct {
	Port int
	GenerateKey bool
}

func ParseOptions() *Options {
	var port = 8080
	var generate_key = false
	flag.IntVar(&port, "port", port, "Specify the port")
	flag.BoolVar(&generate_key, "generate-key", generate_key, "Generate an encryption key")
	flag.Parse()

	return &Options{
		Port: port,
		GenerateKey: generate_key,
	}
}

// Handles all cases where the program needs to exit
func (o *Options)HandleOptions() {
	if o.GenerateKey {
		key, err := encryption.GenerateKey(state.KeySize)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERRO: Could not generate key\n")
			os.Exit(1)
		}
		fmt.Printf("Key: %s\n", hex.EncodeToString(key))
		os.Exit(0)
	}
}
