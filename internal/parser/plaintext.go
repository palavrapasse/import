package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/palavrapasse/import/internal/entity"
)

var separatorsSupported = []string{",", ";"}

const (
	EmailPosition    = 0
	PasswordPosition = 1
)

type PlainTextLeaksParser struct {
	FilePath string
}

func (p PlainTextLeaksParser) Parse() (entity.LeakParse, []error) {
	var errors []error

	lines, err := getFileLines(p.FilePath)

	if err != nil {
		errors = append(errors, err)
		return nil, errors
	}

	leaks := make(map[entity.User]entity.Credentials)

	for _, line := range lines {
		user, credential, err := lineToUserCredential(line)

		if err == nil {
			leaks[user] = credential
		} else {
			errors = append(errors, err)
		}
	}

	return leaks, errors
}

func lineToUserCredential(line string) (entity.User, entity.Credentials, error) {

	containsSeparator := false
	separator := ""

	for _, separator = range separatorsSupported {
		if strings.Contains(line, separator) {
			containsSeparator = true
			break
		}
	}

	if !containsSeparator {
		return entity.User{}, entity.Credentials{}, fmt.Errorf("Input incorrect. Line %v should contain a valid separator (%v)", line, strings.Join(separatorsSupported, " "))
	}

	lineSplit := strings.Split(line, separator)

	email := string(lineSplit[EmailPosition])
	u, err := entity.NewUser(email)

	if err != nil {
		return entity.User{}, entity.Credentials{}, err
	}

	password := string(lineSplit[PasswordPosition])
	p := entity.NewPassword(password)

	c := entity.NewCredentials(p)

	return u, c, nil
}

func getFileLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error while opening file: %v", err)
		return nil, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines, nil
}
