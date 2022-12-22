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
	NumberPositions  = 2
)

type PlainTextLeakParser struct {
	FilePath string
}

func (p PlainTextLeakParser) Parse() (entity.LeakParse, []error) {
	var errors []error

	lines, err := getFileLines(p.FilePath)

	if err != nil {
		errors = append(errors, err)
		return nil, errors
	}

	return linesToLeakParse(lines)
}

func lineToUserCredential(line string, separator string) (entity.User, entity.Credentials, error) {

	if !strings.Contains(line, separator) {
		err := fmt.Errorf("Input incorrect. Line %v should the separator (%v)", line, separator)
		return entity.User{}, entity.Credentials{}, err
	}

	lineSplit := strings.Split(line, separator)

	if len(lineSplit) != NumberPositions {
		err := fmt.Errorf("Input incorrect. Line %v should contain email and password information", line)
		return entity.User{}, entity.Credentials{}, err
	}

	email := string(lineSplit[EmailPosition])
	u, err := entity.NewUser(email)

	if err != nil {
		return entity.User{}, entity.Credentials{}, err
	}

	password := string(lineSplit[PasswordPosition])
	p, err := entity.NewPassword(password)

	if err != nil {
		return entity.User{}, entity.Credentials{}, err
	}

	c := entity.NewCredentials(p)

	return u, c, nil
}

func findSeparator(line string) (string, error) {

	for _, separator := range separatorsSupported {
		if strings.Contains(line, separator) {
			return separator, nil
		}
	}

	err := fmt.Errorf("Input incorrect. Line %v should contain a valid separator (%v)", line, strings.Join(separatorsSupported, " "))
	return "", err
}

func linesToLeakParse(lines []string) (entity.LeakParse, []error) {
	var errors []error
	leak := entity.LeakParse{}

	if len(lines) == 0 {
		errors = append(errors, fmt.Errorf("Can't process empty leak"))
		return leak, errors
	}

	separator, err := findSeparator(lines[0])

	if err != nil {
		errors = append(errors, err)
		return leak, errors
	}

	for _, line := range lines {

		user, credential, err := lineToUserCredential(line, separator)

		if err == nil {
			leak[user] = credential
		} else {
			errors = append(errors, err)
		}
	}

	return leak, errors
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
