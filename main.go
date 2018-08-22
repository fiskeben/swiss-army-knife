package main

import "github.com/majestic-fox/swiss-army-knife/cmd"

var (
	Version   string
	BuildDate string
)

func main() {
	cmd.Execute(Version, BuildDate)
}
