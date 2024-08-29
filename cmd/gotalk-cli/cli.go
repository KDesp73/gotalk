package main

import (
	"fmt"
	"gotalk/internal/state"
	"gotalk/internal/users"
	"gotalk/internal/utils"
	"os"
)


func sudo(s *state.State, name string, un bool) {
	found := false
	for _, user := range s.Users.Items {
		if name == user.Name {
			if un {
				user.Type = users.DEFAULT
			} else {
				user.Type = users.ADMIN
			}
			found = true
			break
		}
	}

	if !found {
		fmt.Fprintf(os.Stderr, "%s not found", name)
	} else if un {
		state.SaveState(s, state.StateFile, nil)
		fmt.Printf("User %s is now not an admin\n", name)
	} else {
		state.SaveState(s, state.StateFile, nil)
		fmt.Printf("User %s is now an admin\n", name)
	}
}

func loadState(path string) *state.State {
	var s *state.State
	if utils.FileExists(path) {
		var err error
		s, err = state.LoadState(path, nil)

		if err != nil || s == nil {
			fmt.Fprintf(os.Stderr, "ERRO: Could not load state from file: %s\n", path)
			os.Exit(1)
		} 
	} else {
		fmt.Fprintf(os.Stderr, "ERRO: State file '%s' not found\n", path)
		os.Exit(1)
	}
	return s
}

func deleteThread(s *state.State, thread string) {
	if thread != "all" && s.Threads.Get(thread) == nil {
		fmt.Fprintf(os.Stderr, "Thread '%s' doesn't exist\n", thread)
		os.Exit(1)
	}

	if thread != "all" {
		s.Threads.RemoveThread(thread)
	} else {
		for _, t := range state.Instance.Threads.Items {
			state.Instance.Threads.RemoveThread(t.ID)
		}
	}
	state.SaveState(s, state.StateFile, nil)
}

func main() {
	options := ParseOptions()	
	s := loadState(options.StatePath)
	
	if(len(os.Args) < 2) {
		fmt.Fprintln(os.Stderr, "ERRO: At least 1 argument is required")
		os.Exit(1)
	}

	if options.PrintState {
		fmt.Println(utils.JsonToString(s))
		os.Exit(0)
	}

	if options.AddTest {
		s.Users.PushUser(&users.User{
			Name: "test",
			Email: "test@email.com",
			Type: users.DEFAULT,
			SignUpTime: utils.CurrentTimestamp(),
		})
		state.SaveState(s, state.StateFile, nil)
	}

	if options.AddAdmin {
		s.Users.PushUser(&users.User{
			Name: "admin",
			Email: "admin@email.com",
			Type: users.ADMIN,
			SignUpTime: utils.CurrentTimestamp(),
		})
		state.SaveState(s, state.StateFile, nil)
	}

	if options.DeleteThread != "" {
		thread := options.DeleteThread
		deleteThread(s, thread)
	}

	if options.Backup {
		utils.CopyFile(state.StateFile, state.StateFileBackup)
		fmt.Printf("Backup created at %s\n", state.StateFileBackup)
	}

	if options.Comments {
		listComments(s)
	} else if options.Threads {
		listThreads(s)
	} else if options.Users {
		listUsers(s)
	}

	if options.Sudo != "" {
		sudo(s, options.Sudo, options.Un)
	}
}
