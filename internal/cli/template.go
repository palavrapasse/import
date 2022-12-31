package cli

import (
	"fmt"

	"github.com/palavrapasse/damn/pkg/entity"
)

var exampleCommand = fmt.Sprintf(`./import --database-path="path/db.sqlite" --leak-path="path/file.txt" --context="context" --platforms="platform1, platform2" --share-date="%s" --leakers="leaker1, leaker2"`,
	entity.DateFormatLayout)

func CreateAppHelpTemplate(base string) string {

	return fmt.Sprintf(`%s
EXAMPLE: 
	%s

WEBSITE:
	https://github.com/palavrapasse

`, base, exampleCommand)
}
