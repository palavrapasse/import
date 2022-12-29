package cli

import "fmt"

func CreateAppHelpTemplate(base string) string {

	// Append to an existing template
	return fmt.Sprintf(`%s
EXAMPLE: 
	%s

WEBSITE:
	https://github.com/palavrapasse

`, base, exampleCommand)
}
