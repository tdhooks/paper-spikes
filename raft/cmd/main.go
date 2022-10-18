package main

import (
	"fmt"

	"github.com/tdhooks/paper-spikes/raft/internal/hello"
)

func main() {
	fmt.Print(hello.SayHello("Dayne"))
}
