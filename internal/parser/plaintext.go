package parser

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"sync"

	"github.com/palavrapasse/damn/pkg/entity/query"
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

const MaxLinesOfGoroutine = 5000

var supportedSeparators = []string{ColonSeparator, CommaSeparator, SemiColonSeparator}

type PlainTextLeakParser struct {
	FilePath string
}

type linesParseResult struct {
	query.LeakParse
	errors []error
}

func (p PlainTextLeakParser) Parse(ecb ...OnParseErrorCallback) (query.LeakParse, []error) {
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

	err := fmt.Errorf("input incorrect. Line %v should contain a valid separator (%v)", line, strings.Join(supportedSeparators, " "))
	return "", err
}

func lineToUser(line string, separator string) (query.User, error) {

	if !strings.Contains(line, separator) {
		err := fmt.Errorf("input incorrect. Line %v should the separator (%v)", line, separator)
		return query.User{}, err
	}

	lineSplit := strings.Split(line, separator)

	if len(lineSplit) < NumberPositions {
		err := fmt.Errorf("input incorrect. Line %v should contain email and password information", line)
		return query.User{}, err
	}

	emailString := string(lineSplit[EmailPosition])
	email, err := query.NewEmail(emailString)

	if err != nil {
		return query.User{}, err
	}

	u := query.NewUser(email)

	password := string(strings.Join(lineSplit[PasswordPosition:], separator))
	_, err = query.NewPassword(password)

	if err != nil {
		return query.User{}, err
	}

	return u, nil
}

func linesToLeakParse(lines []string, ecb ...OnParseErrorCallback) (query.LeakParse, []error) {
	var errors []error
	leak := query.LeakParse{}

	if len(lines) == 0 {
		err := fmt.Errorf("can't process empty leak")

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

	nlines := len(lines)

	ngoroutines := 1
	if nlines > MaxLinesOfGoroutine {
		ngoroutines = int(math.Ceil(float64(nlines) / float64(MaxLinesOfGoroutine)))
	}

	linesParseResultChan := make(chan linesParseResult)

	var wg sync.WaitGroup

	wg.Add(ngoroutines)

	for i := 0; i < ngoroutines; i++ {

		init := i * MaxLinesOfGoroutine
		end := (i + 1) * MaxLinesOfGoroutine
		if end > nlines {
			end = nlines
		}

		go func(lines []string) {

			defer wg.Done()
			linesParseResultChan <- routineLinesToLeakParse(lines, separator)

		}(lines[init:end])
	}

	go func() {
		wg.Wait()
		close(linesParseResultChan)
	}()

	for s := range linesParseResultChan {
		errors = append(errors, s.errors...)
		leak = append(leak, s.LeakParse...)
	}

	return leak, errors
}

func routineLinesToLeakParse(lines []string, separator string, ecb ...OnParseErrorCallback) linesParseResult {
	leak := query.LeakParse{}
	var errors []error

	for _, line := range lines {

		user, err := lineToUser(line, separator)

		if err == nil {
			leak = append(leak, user)
		} else {
			processOnParseError(err, ecb...)
			errors = append(errors, err)
		}
	}

	return linesParseResult{
		LeakParse: leak,
		errors:    errors,
	}
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
