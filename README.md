![GitHub Workflow Status](https://img.shields.io/github/workflow/status/juggler434/marvel_server/Go)
[![Coverage Status](https://coveralls.io/repos/github/juggler434/marvel_server/badge.svg?branch=development)](https://coveralls.io/github/juggler434/marvel_server?branch=development)
[![Go Report Card](https://goreportcard.com/badge/github.com/juggler434/marvel_server)](https://goreportcard.com/report/github.com/juggler434/marvel_server)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# Hero Cards

This is the API for running a certain comic themed card game

## Requirements
go 1.12

## Setup
`go build` will create a binary for marvelServer that can be run like any other binary.  There are currently no commands to be run.

## Development

### Testing
To run tests, from project directory run `go test -v ./...`
All tests must pass before a pull request will be merged

### Formatting
All code must conform to gofmt and golint

## Contributing 
Commit message must have a detailed explanation of what the commit contains
Rebase master into personal branches to avoid merge commits

### Test Coverage
90% test coverage is required or justification on why that is not attainable
