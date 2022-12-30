package parser

import (
	"fmt"
	"testing"

	"github.com/palavrapasse/import/internal/entity"
)

func TestCannotParseEmptyLines(t *testing.T) {
	lines := []string{}

	_, err := linesToLeakParse(lines)

	if err == nil {
		t.Fatalf("Lines designated by the string below contains multiple lines which are invalid, but no error was identified\nString: %s", lines)
	}

	if len(err) != 1 {
		t.Fatalf("Lines is empty so it should contain one error")
	}
}

func TestCannotParseLinesToLeakWithOnlyOneLineWhichIsInvalid(t *testing.T) {
	lines := []string{"fghj2@aaa"}

	_, err := linesToLeakParse(lines)

	if err == nil {
		t.Fatalf("Lines designated by the string below only contains one line which is invalid, but no error was identified\nString: %s", lines)
	}
}

func TestCannotParseLinesToLeakWithMultipleLinesWhichAreInvalid(t *testing.T) {
	lines := []string{"fghj2@aaa,", "fghj2,dghf", ",dghf", "fghj2,", ",dghf", "fghj2,", ",dg,hf,"}

	_, err := linesToLeakParse(lines)

	if err == nil {
		t.Fatalf("Lines designated by the string below contains multiple lines which are invalid, but no error was identified\nString: %s", lines)
	}

	if len(lines) != len(err) {
		t.Fatalf("Lines contains %v lines which are invalid, but an error contains %v errors. They must be the same", len(lines), len(err))
	}
}

func TestCannotParseLinesToLeakWithFirstLineWithoutValidSeparator(t *testing.T) {
	lines := []string{"fghj2@aaa", "fghj2,dghf", ",dghf", "fghj2,", ",dghf", "fghj2,", ",dg,hf,"}

	_, err := linesToLeakParse(lines)

	if err == nil {
		t.Fatalf("Lines designated by the string below contains multiple lines which are invalid, but no error was identified\nString: %s", lines)
	}

	if len(err) != 1 {
		t.Fatalf("The first line of Lines designated by the string below does not contain a valid separator so it should contain one error\nString: %s", lines)
	}
}

func TestCanParseLinesToLeakWithSomeInvalidLines(t *testing.T) {
	lines := []string{"fghj2@aaa,", "fghj2@aaa;dghf", "fghj2@aaa,dghf"}

	leak, err := linesToLeakParse(lines)

	if err == nil {
		t.Fatalf("Lines designated by the string below contains some invalid lines, but no error was identified\nString: %s", lines)
	}

	if len(leak) == 0 {
		t.Fatalf("Lines designated by the string below contains some valid lines, but Leak is empty\nString: %s", lines)
	}
}

func TestCanParseLinesToLeakWithOnlyValidLines(t *testing.T) {
	lines := []string{"test@aaa,dghf", "fghj2@aaa,dghf"}

	leak, err := linesToLeakParse(lines)

	if err != nil {
		t.Fatalf("Lines designated by the string below does not contain invalid lines, but an error was identified\nString: %s", lines)
	}

	if len(leak) == 0 {
		t.Fatalf("Lines designated by the string below contains some valid lines, but Leak is empty\nString: %s", lines)
	}
}

func TestCanParseLinesToLeakWithPasswordThatContainsSeparator(t *testing.T) {
	lines := []string{"test@aaa:dghf:aaa"}

	leak, err := linesToLeakParse(lines)

	panicOnError(err)

	if len(leak) != 1 {
		t.Fatalf("Lines designated by the string below contains some valid lines, but Leak is empty\nString: %s", lines)
	}
}

func TestCanParsePasswordThatDoesNotContainSeparator(t *testing.T) {
	email := "test@aaa"
	password := "dghfaaa,;.78"
	lines := []string{fmt.Sprintf("%s:%s", email, password)}

	leak, err := linesToLeakParse(lines)

	panicOnError(err)

	if len(leak) != 1 {
		t.Fatalf("Lines designated by the string below contains some valid lines, but Leak is empty\nString: %s", lines)
	}

	user := entity.User{Email: entity.Email(email)}
	result := leak[user].Password

	if result != entity.Password(password) {
		t.Fatalf("Password designated by the string below should be the same as '%s'\nString: %s", password, result)
	}
}

func TestCanParsePasswordThatContainsSeparator(t *testing.T) {
	email := "test@aaa"
	password := "dghf:aaa,;.78"
	lines := []string{fmt.Sprintf("%s:%s", email, password)}

	leak, err := linesToLeakParse(lines)

	panicOnError(err)

	if len(leak) != 1 {
		t.Fatalf("Lines designated by the string below contains some valid lines, but Leak is empty\nString: %s", lines)
	}

	user := entity.User{Email: entity.Email(email)}
	result := leak[user].Password

	if result != entity.Password(password) {
		t.Fatalf("Password designated by the string below should be the same as '%s'\nString: %s", password, result)
	}
}

func TestCannotFindUnknownSeparator(t *testing.T) {
	line := "my.username-my.password"

	sep, err := findSeparator(line)

	if err == nil {
		t.Fatalf("No valid separator is present in line, but a separator was found: %s\n", sep)
	}
}

func TestCanFindCommaSeparator(t *testing.T) {
	testSep := CommaSeparator
	line := fmt.Sprintf("my.username%smy.password", testSep)

	sep, _ := findSeparator(line)

	if sep != testSep {
		t.Fatalf("A comma separator is present in line, but a different separator was found (%s)\n", testSep)
	}
}

func TestCanFindColonSeparator(t *testing.T) {
	testSep := ColonSeparator
	line := fmt.Sprintf("my.username%smy.password", testSep)

	sep, _ := findSeparator(line)

	if sep != testSep {
		t.Fatalf("A colon separator is present in line, but a different separator was found (%s)\n", testSep)
	}
}

func TestCanFindSemiColonSeparator(t *testing.T) {
	testSep := SemiColonSeparator
	line := fmt.Sprintf("my.username%smy.password", testSep)

	sep, _ := findSeparator(line)

	if sep != testSep {
		t.Fatalf("A semicolon separator is present in line, but a different separator was found (%s)\n", testSep)
	}
}

func panicOnError(err []error) {
	if len(err) != 0 {
		panic(err)
	}
}
