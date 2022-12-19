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
	ContextPosition = 0
	DatePosition    = 1
)

type PlainTextLeaksParser struct {
	FilePath string
}

func (p PlainTextLeaksParser) ParseLeaks() ([]entity.Leak, error) {
	lines, err := getFileLines(p.FilePath)

	if err != nil {
		return nil, err
	}

	size := len(lines)
	leaks := make([]entity.Leak, size)

	for i, line := range lines {
		leak, err := lineToLeak(line)

		if err != nil {
			leaks[i] = leak
		}
	}

	return leaks, nil
}

func lineToLeak(line string) (entity.Leak, error) {

	containsSeparator := false
	separator := ""

	for _, separator = range separatorsSupported {
		if strings.Contains(line, separator) {
			containsSeparator = true
			break
		}
	}

	if !containsSeparator {
		return entity.Leak{}, fmt.Errorf("Input incorrect. Line %v should contain a valid separator (%v)", line, strings.Join(separatorsSupported, " "))
	}

	lineSplit := strings.Split(line, separator)

	context := string(lineSplit[ContextPosition])
	date := string(lineSplit[DatePosition])
	ds, err := entity.NewDateInSeconds(date)

	if err != nil {
		return entity.Leak{}, err
	}

	leak, err := entity.NewLeak(context, ds)

	if err != nil {
		return entity.Leak{}, err
	}

	return leak, nil
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
