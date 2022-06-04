import "strings"

uname: strings.Exec("uname -a")

helloWorld: hello({inputs: name: "world"})

hello: {
	name: string
	message: "hello, \(name)!"
}
