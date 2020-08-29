package main

import (
	"os"
	"runtime"

	"github.com/RafilxTenfen/go-chat/cli/user"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	os.Exit(user.Main(os.Args[1:]))
}
