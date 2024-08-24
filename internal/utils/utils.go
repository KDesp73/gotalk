package utils

import (
	"encoding/json"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

func CurrentTimestamp() string {
	currentTime := time.Now()
	return currentTime.Format("2006/01/02 15:04:05")
}

func JsonToString(jsonStruct any) string {
	jsonData, err := json.Marshal(jsonStruct)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func StrEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}

func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}

	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	sourceInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}
	err = os.Chmod(dst, sourceInfo.Mode())
	if err != nil {
		return err
	}

	return nil
}

func IsValidEmail(email string) bool {
	const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailPattern)

	return re.MatchString(email)
}

func ShortenString(s string, length int) string {
	if len(s) <= length {
		return s
	}
	if length < 3 {
		return s[:length]
	}
	return s[:length-3] + "..."
}
