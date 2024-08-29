package threads

import (
	"fmt"
	"testing"
)

func checkComment(thread *Thread, index int, authorExpected string, contentExpected string) error {
	if thread.Comments[index].ThreadID != thread.ID {
		return fmt.Errorf("Comment did not aquire the parent thread id")
	}

	if thread.Comments[index].Author != authorExpected {
		return fmt.Errorf("Comment doesn't have the correct author. Should be: %s. Is: %s\n", authorExpected, thread.Comments[index].Author)
	}

	if thread.Comments[index].Content != contentExpected{
		return fmt.Errorf("Comment doesn't have the correct content. Should be: %s. Is: %s\n", contentExpected, thread.Comments[index].Author)
	}

	return nil
}

func TestPushComment(t *testing.T) {
	thread := &Thread {
		ID: "test",
	}
	
	thread.PushComment("KDesp73", "Test")

	err := checkComment(thread, 0, "KDesp73", "Test")

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestPushCommentTwice(t *testing.T) {
	thread := &Thread {
		ID: "test",
	}

	var err error

	thread.PushComment("KDesp73", "Test")
	err = checkComment(thread, 0, "KDesp73", "Test")
	if err != nil {
		t.Fatalf(err.Error())
	}

	thread.PushComment("John", "Hello World")
	err = checkComment(thread, 1, "John", "Hello World")
	if err != nil {
		t.Fatalf(err.Error())
	}
}
