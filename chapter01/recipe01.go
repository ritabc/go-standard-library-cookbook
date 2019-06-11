package main

import (
	"log"
	"runtime"
)

const info = `Application %s starting.
The binary was built by GO: %s`

func main() {
	log.Printf(info, "Example", runtime.Version())
}
