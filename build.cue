package main

import (
	"dagger.io/dagger"
	"dagger.io/dagger/core"
	"universe.dagger.io/go"
)

dagger.#Plan & {
	_source: core.#Source & {
		path: "."
	}
	actions: build: go.#Build & {
		source: _source.output
		package: "./cmd/cue"
	}
}
