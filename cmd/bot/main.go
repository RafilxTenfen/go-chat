package main

import (
	"os"
	"runtime"

	"github.com/RafilxTenfen/go-chat/cli/bot"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	os.Exit(bot.Main(os.Args[1:]))
}
