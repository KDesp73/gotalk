package main

import (
	"flag"
	"fmt"
	"gotalk/internal/state"
	"gotalk/internal/utils"
	"gotalk/internal/users"
	"os"
	"strconv"
	"strings"
)

func dashes(count int) string {
	return strings.Repeat("-", count)
}

func listComments(s *state.State) {
	if(len(s.Threads.Threads) == 0){
		fmt.Printf("No threads found\n");
		os.Exit(0)
	}

	allEmpty := true
	for _, thread := range s.Threads.Threads{
		if len(thread.Comments) != 0 {
			allEmpty = false
		}
	}

	if allEmpty {
		fmt.Printf("No comments found\n")
		os.Exit(0)
	}

	idDashes := 16
	authorDashes := 20
	contentDashes := 30
	threadIdDashes := 16
	timeDashes := 19
	lineFmt := "| %*s | %*s | %*s | %*s | %*s |\n"
	borderFmt := "+ %s + %s + %s + %s + %s +\n"

	fmt.Printf(borderFmt, dashes(idDashes), dashes(authorDashes), dashes(contentDashes), dashes(threadIdDashes), dashes(timeDashes))
	fmt.Printf(lineFmt, -idDashes, "ID", -authorDashes, "Author", -contentDashes, "Content", -threadIdDashes, "ThreadID", -timeDashes, "Timestamp");
	fmt.Printf(borderFmt, dashes(idDashes), dashes(authorDashes), dashes(contentDashes), dashes(threadIdDashes), dashes(timeDashes))

	for _, thread := range s.Threads.Threads {
		if thread == nil {
			continue
		}
		for _, comment := range thread.Comments {
			if comment == nil {
				continue
			}
			fmt.Printf(lineFmt, 
				-idDashes, utils.ShortenString(comment.ID, idDashes), 
				-authorDashes, utils.ShortenString(comment.Author, authorDashes), 
				-contentDashes, utils.ShortenString(comment.Content, contentDashes), 
				-threadIdDashes, utils.ShortenString(comment.ThreadID, threadIdDashes), 
				-timeDashes, comment.Timestamp,
			)
		}
	}
	fmt.Printf(borderFmt, dashes(idDashes), dashes(authorDashes), dashes(contentDashes), dashes(threadIdDashes), dashes(timeDashes))
}

func listThreads(s *state.State) {
	if(len(s.Threads.Threads) == 0){
		fmt.Printf("No threads found\n");
		os.Exit(0)
	}

	idDashes := 36
	commentCountDashes := 13
	lineFmt := "| %*s | %*s |\n"

	fmt.Printf("+ %s + %s +\n", dashes(idDashes), dashes(commentCountDashes))
	fmt.Printf(lineFmt, -idDashes, "ID", -commentCountDashes, "Comment Count")
	fmt.Printf("+ %s + %s +\n", dashes(idDashes), dashes(commentCountDashes))

	for _, thread := range s.Threads.Threads {
		fmt.Printf(lineFmt, 
			-idDashes, thread.ID,
			-commentCountDashes, strconv.Itoa(len(thread.Comments)),
		)
	}
	fmt.Printf("+ %s + %s +\n", dashes(idDashes), dashes(commentCountDashes))
}

func listUsers(s *state.State){
	if len(s.Users.Users) == 0 {
		fmt.Printf("No users found\n")
		os.Exit(0)
	}

	lineFmt := "| %*s | %*s | %*s | %*s | %*s |\n"
	borderFmt := "+ %s + %s + %s + %s + %s +\n"

	nameDashes := 20
	emailDashes := 24
	keyDashes := 36
	typeDashes := 20
	timeDashes := 19
	fmt.Printf(borderFmt, dashes(nameDashes), dashes(emailDashes), dashes(keyDashes), dashes(typeDashes), dashes(timeDashes))
	fmt.Printf(lineFmt, -nameDashes, "Name", -emailDashes, "Email", -keyDashes, "Key", -typeDashes, "Type", -timeDashes, "Timestamp");
	fmt.Printf(borderFmt, dashes(nameDashes), dashes(emailDashes), dashes(keyDashes), dashes(typeDashes), dashes(timeDashes))

	for _, user := range s.Users.Users {
		fmt.Printf(lineFmt, 
			-nameDashes, utils.ShortenString(user.Name, nameDashes), 
			-emailDashes, utils.ShortenString(user.Email, emailDashes), 
			-keyDashes, utils.ShortenString(user.Key, keyDashes), 
			-typeDashes, utils.ShortenString(user.Type, typeDashes), 
			-timeDashes, user.SignUpTime,
		)
	}
	fmt.Printf(borderFmt, dashes(nameDashes), dashes(emailDashes), dashes(keyDashes), dashes(typeDashes), dashes(timeDashes))

}

func sudo(s *state.State, name string, un bool) {
	found := false
	for _, user := range s.Users.Users {
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

func main() {
	var commentsFlag = false
	var threadsFlag = false
	var usersFlag = false
	var backupFlag = false
	var unFlag = false
	var sudoFlag = ""
	flag.BoolVar(&commentsFlag, "comments", commentsFlag, "List comments")
	flag.BoolVar(&threadsFlag, "threads", threadsFlag, "List threads")
	flag.BoolVar(&usersFlag, "users", usersFlag, "List users")
	flag.BoolVar(&backupFlag, "backup", backupFlag, "Backup state")
	flag.BoolVar(&unFlag, "un", unFlag, "Reverse action")
	flag.StringVar(&sudoFlag, "sudo", sudoFlag, "Make user an admin")
	flag.Parse()
	
	var s *state.State
	if utils.FileExists(state.StateFile) {
		var err error
		s, err = state.LoadState(state.StateFile, nil)

		if err != nil {
			fmt.Fprintf(os.Stderr, "ERRO: Could not load state from file\n")
			os.Exit(1)
		} 
	} else {
		fmt.Fprintf(os.Stderr, "ERRO: State file not found\n")
		os.Exit(1)
	}

	if(len(os.Args) < 2) {
		fmt.Fprintln(os.Stderr, "ERRO: At least 1 argument is required")
		os.Exit(1)
	}

	if backupFlag {
		utils.CopyFile(state.StateFile, state.StateFileBackup)
		fmt.Printf("Backup created at %s\n", state.StateFileBackup)
	}

	if commentsFlag {
		listComments(s)
	} else if threadsFlag {
		listThreads(s)
	} else if usersFlag {
		listUsers(s)
	}

	if strings.TrimSpace(sudoFlag) != "" {
		sudo(s, strings.TrimSpace(sudoFlag), unFlag)
	}
}
