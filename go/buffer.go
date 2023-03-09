package main

import (
	"bytes"
	"log"
)

func main() {
	var buf bytes.Buffer
	var RoleLevel int

	logger := log.New(&buf, "logger: ", log.Lshortfile)
}
