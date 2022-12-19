package internal

import (
	"strings"
	"testing"
)

func TestNewBadActorWithIdentifierEmpty(t *testing.T) {
	identifier := ""

	_, err := NewBadActor(identifier)

	if err == nil {
		t.Fatalf("Identifier designated by the string below is empty, but no error was identified\nString: %s", identifier)
	}
}

func TestNewBadActorWithIdentifierWithOnlySpaces(t *testing.T) {
	identifier := "   "

	_, err := NewBadActor(identifier)

	if err == nil {
		t.Fatalf("Identifier designated by the string below contains only spaces, but no error was identified\nString: %s", identifier)
	}
}

func TestNewBadActorWithIdentifierThatExceeds30Characters(t *testing.T) {
	identifier := strings.Repeat("x", 31)

	_, err := NewBadActor(identifier)

	if err == nil {
		t.Fatalf("Identifier designated by the string below exceeds 30 characters, but no error was identified\nString: %s", identifier)
	}
}

func TestNewBadActorWithIdentifierThatMatches30Characters(t *testing.T) {
	identifier := strings.Repeat("x", 30)

	_, err := NewBadActor(identifier)

	if err != nil {
		t.Fatalf("Identifier designated by the string below matches 30 characters, but an error was identified\nString: %s", identifier)
	}
}

func TestNewBadActorWithIdentifierThatDoesNotExceed30Characters(t *testing.T) {
	identifier := strings.Repeat("x", 20)

	_, err := NewBadActor(identifier)

	if err != nil {
		t.Fatalf("Identifier designated by the string below does not exceed 30 characters, but an error was identified\nString: %s", identifier)
	}
}

func TestNewBadActorTrimsIdentifierSpaces(t *testing.T) {
	identifier := " identifier    "

	ba, _ := NewBadActor(identifier)

	if len(ba.Identifier) == len(identifier) {
		t.Fatalf("Original identifier string contains unneeded spaces, and should be trimmed, but output summary still contains those spaces")
	}
}

func TestCopyBadActor(t *testing.T) {
	ba, _ := NewBadActor("identifier")
	key := AutoGenKey(10)
	copy := ba.Copy(key)

	if copy.Identifier != ba.Identifier {
		t.Fatalf("Copy identifier is different from the original Identifier")
	}

	if copy.BaId != key {
		t.Fatalf("Copy BaId is different from the provided\nKey: %v", key)
	}
}
