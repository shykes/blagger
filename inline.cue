import "strings"


helloWorld: hello({inputs: name: "world"})

hello: {
	name: string
	message: "hello, \(name)!"
}
