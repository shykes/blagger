import (
	"dagger"
)

hello: dagger.Shell({
	script: "echo hello world | cat -en"
})
