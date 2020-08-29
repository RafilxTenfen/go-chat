package main

import (
	"os"
	"runtime"

	"github.com/RafilxTenfen/go-chat/cli/chatapi"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	os.Exit(chatapi.Main(os.Args[1:]))
}
