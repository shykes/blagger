import (
	"strings"
)


tempdir: strings.TrimSpace(strings.Exec("mktemp -d"))

_clone: strings.Exec("cd \(tempdir) && git clone https://github.com/dagger/todoapp")

dirs: [path=string]: {
	contents: strings.Exec("ls -l \(path)")
}

dirs: {
	"\(tempdir)": _
	"\(tempdir)/todoapp": _
	"\(tempdir)/todoapp/src": _
}
