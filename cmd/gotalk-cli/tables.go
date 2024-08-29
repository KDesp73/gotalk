package main

import (
	"fmt"
	"gotalk/internal/state"
	"gotalk/internal/utils"
	"os"
	"strconv"
	"strings"
)

func dashes(count int) string {
	return strings.Repeat("-", count)
}

func listComments(s *state.State) {
	if(len(s.Threads.Items) == 0){
		fmt.Printf("No threads found\n");
		os.Exit(0)
	}

	allEmpty := true
	for _, thread := range s.Threads.Items{
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

	for _, thread := range s.Threads.Items {
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
	if(len(s.Threads.Items) == 0){
		fmt.Printf("No threads found\n");
		os.Exit(0)
	}

	titleDashes := 20
	idDashes := 36
	commentCountDashes := 13
	lineFmt := "| %*s | %*s | %*s |\n"

	fmt.Printf("+ %s + %s + %s +\n", dashes(titleDashes), dashes(idDashes), dashes(commentCountDashes))
	fmt.Printf(lineFmt, -titleDashes, "Title", -idDashes, "ID", -commentCountDashes, "Comment Count")
	fmt.Printf("+ %s + %s + %s +\n", dashes(titleDashes), dashes(idDashes), dashes(commentCountDashes))

	for _, thread := range s.Threads.Items {
		fmt.Printf(lineFmt, 
			-titleDashes, thread.Title,
			-idDashes, thread.ID,
			-commentCountDashes, strconv.Itoa(len(thread.Comments)),
		)
	}
	fmt.Printf("+ %s + %s + %s +\n", dashes(titleDashes), dashes(idDashes), dashes(commentCountDashes))
}

func listUsers(s *state.State){
	if len(s.Users.Items) == 0 {
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

	for _, user := range s.Users.Items {
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
