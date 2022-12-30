package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/palavrapasse/import/internal/entity"
)

const (
	EmailPosition    = 0
	PasswordPosition = 1
	NumberPositions  = 2
)

const (
	CommaSeparator     = ","
	ColonSeparator     = ":"
	SemiColonSeparator = ";"
)

var supportedSeparators = []string{ColonSeparator, CommaSeparator, SemiColonSeparator}

type PlainTextLeakParser struct {
	FilePath string
}

func (p PlainTextLeakParser) Parse(ecb ...OnParseErrorCallback) (entity.LeakParse, []error) {
	var errors []error

	lines, err := getFileLines(p.FilePath)

	if err != nil {
		processOnParseError(err, ecb...)
		errors = append(errors, err)

		return nil, errors
	}

	return linesToLeakParse(lines, ecb...)
}

func findSeparator(line string) (string, error) {

	for _, separator := range supportedSeparators {
		if strings.Contains(line, separator) {
			return separator, nil
		}
	}

	err := fmt.Errorf("Input incorrect. Line %v should contain a valid separator (%v)", line, strings.Join(supportedSeparators, " "))
	return "", err
}

func lineToUserCredential(line string, separator string) (entity.User, entity.Credentials, error) {

	if !strings.Contains(line, separator) {
		err := fmt.Errorf("Input incorrect. Line %v should the separator (%v)", line, separator)
		return entity.User{}, entity.Credentials{}, err
	}

	lineSplit := strings.Split(line, separator)

	if len(lineSplit) < NumberPositions {
		err := fmt.Errorf("Input incorrect. Line %v should contain email and password information", line)
		return entity.User{}, entity.Credentials{}, err
	}

	email := string(lineSplit[EmailPosition])
	u, err := entity.NewUser(email)

	if err != nil {
		return entity.User{}, entity.Credentials{}, err
	}

	password := string(strings.Join(lineSplit[PasswordPosition:], separator))
	p, err := entity.NewPassword(password)

	if err != nil {
		return entity.User{}, entity.Credentials{}, err
	}

	c := entity.NewCredentials(p)

	return u, c, nil
}

func linesToLeakParse(lines []string, ecb ...OnParseErrorCallback) (entity.LeakParse, []error) {
	var errors []error
	leak := entity.LeakParse{}

	if len(lines) == 0 {
		err := fmt.Errorf("Can't process empty leak")

		processOnParseError(err, ecb...)
		errors = append(errors, err)

		return leak, errors
	}

	separator, err := findSeparator(lines[0])

	if err != nil {
		processOnParseError(err, ecb...)
		errors = append(errors, err)

		return leak, errors
	}

	for _, line := range lines {

		user, credential, err := lineToUserCredential(line, separator)

		if err == nil {
			leak[user] = credential
		} else {
			processOnParseError(err, ecb...)
			errors = append(errors, err)
		}
	}

	return leak, errors
}

func getFileLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
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
