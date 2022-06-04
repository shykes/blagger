import (
	"strings"
)

pipeline: [
	strings.Exec("sleep 3; touch foo"),
	strings.Exec("cp foo bar"),
	strings.Exec("cp bar baz"),
]
