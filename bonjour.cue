package bonjour

import (
	"strings"
)

hello: strings.Exec("echo \(message)")

message: strings.Exec("docker ps")

