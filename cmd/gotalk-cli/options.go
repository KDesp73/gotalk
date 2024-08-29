package main

import (
	"flag"
	"strings"
)

type Options struct {
	Comments bool
	Threads bool
	DeleteThread string
	Users bool
	Backup bool
	Un bool
	Sudo string
	AddTest bool
	AddAdmin bool
	PrintState bool
}

func ParseOptions() Options {
	options := Options{}

	flag.BoolVar(&options.Comments, "comments", options.Comments, "List comments")
	flag.BoolVar(&options.Threads, "threads", options.Threads, "List threads")
	flag.BoolVar(&options.Users, "users", options.Users, "List users")
	flag.BoolVar(&options.Backup, "backup", options.Backup, "Backup state")
	flag.BoolVar(&options.Un, "un", options.Un, "Reverse action")
	flag.StringVar(&options.DeleteThread, "delete-thread", options.DeleteThread, "Deletes threads")
	flag.StringVar(&options.Sudo, "sudo", options.Sudo, "Make user an admin")
	flag.BoolVar(&options.AddTest, "add-test", options.AddTest, "Add a test user")
	flag.BoolVar(&options.AddAdmin, "add-admin", options.AddAdmin, "Add an adminuser")
	flag.BoolVar(&options.PrintState, "print-state", options.PrintState, "Prints the state struct")
	flag.Parse()

	options.Sudo = strings.TrimSpace(options.Sudo)
	options.DeleteThread = strings.TrimSpace(options.DeleteThread)

	return options
}
