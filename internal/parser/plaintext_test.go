package parser

import "testing"

func TestCannotParseLinesToLeakWithOnlyOneLineWhichIsInvalid(t *testing.T) {
	lines := []string{"fghj2@aaa"}

	_, err := linesToLeakParse(lines)

	if err == nil {
		t.Fatalf("Lines designated by the string below only contains one line which is invalid, but no error was identified\nString: %s", lines)
	}
}

func TestCannotParseLinesToLeakWithMultipleLinesWhichAreInvalid(t *testing.T) {
	lines := []string{"fghj2@aaa", "fghj2,dghf", ",dghf", "fghj2,", ",dghf", "fghj2,", ",dg,hf,"}

	_, err := linesToLeakParse(lines)

	if err == nil {
		t.Fatalf("Lines designated by the string below contains multiple lines which are invalid, but no error was identified\nString: %s", lines)
	}

	if len(lines) != len(err) {
		t.Fatalf("Lines contains %v lines which are invalid, but an error contains %v errors. They must be the same", len(lines), len(err))
	}
}

func TestCanParseLinesToLeakWithSomeInvalidLines(t *testing.T) {
	lines := []string{"fghj2@aaa", "fghj2@aaa;dghf", "fghj2@aaa,dghf"}

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
