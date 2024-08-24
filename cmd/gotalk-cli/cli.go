package main

import (
	"flag"
	"fmt"
	"gotalk/api/state"
	"gotalk/internal/utils"
	"os"
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

	idDashes := 8
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
				-idDashes, utils.ShortenString(comment.ID, idDashes+2), 
				-authorDashes, utils.ShortenString(comment.Author, authorDashes+2), 
				-contentDashes, utils.ShortenString(comment.Content, contentDashes+2), 
				-threadIdDashes, utils.ShortenString(comment.ThreadID, threadIdDashes+2), 
				-timeDashes, comment.Timestamp,
			)
		}
	}
	fmt.Printf(borderFmt, dashes(idDashes), dashes(authorDashes), dashes(contentDashes), dashes(threadIdDashes), dashes(timeDashes))
}

func listUsers(s *state.State){
	lineFmt := "| %*s | %*s | %*s | %*s | %*s |\n"
	borderFmt := "+ %s + %s + %s + %s + %s +\n"

	nameDashes := 20
	emailDashes := 30
	keyDashes := 30
	typeDashes := 20
	timeDashes := 19
	fmt.Printf(borderFmt, dashes(nameDashes), dashes(emailDashes), dashes(keyDashes), dashes(typeDashes), dashes(timeDashes))
	fmt.Printf(lineFmt, -nameDashes, "Name", -emailDashes, "Email", -keyDashes, "Key", -typeDashes, "Type", -timeDashes, "Timestamp");
	fmt.Printf(borderFmt, dashes(nameDashes), dashes(emailDashes), dashes(keyDashes), dashes(typeDashes), dashes(timeDashes))

	for _, user := range s.Users.Users {
		fmt.Printf(lineFmt, 
			-nameDashes, utils.ShortenString(user.Name, nameDashes-2), 
			-emailDashes, utils.ShortenString(user.Email, emailDashes-2), 
			-keyDashes, utils.ShortenString(user.Key, keyDashes-2), 
			-typeDashes, utils.ShortenString(user.Type, typeDashes-2), 
			-timeDashes, user.SignUpTime,
		)
	}
	fmt.Printf(borderFmt, dashes(nameDashes), dashes(emailDashes), dashes(keyDashes), dashes(typeDashes), dashes(timeDashes))

}

func main() {
	var comments = false
	var threads = false
	var users = false
	flag.BoolVar(&comments, "comments", comments, "List comments")
	flag.BoolVar(&threads, "threads", threads, "List threads")
	flag.BoolVar(&users, "users", users, "List users")
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
		fmt.Println("At least 1 argument is required")
		os.Exit(1)
	}

	if comments {
		listComments(s)
	} else if threads {
	} else if users {
		listUsers(s)
	}
}
