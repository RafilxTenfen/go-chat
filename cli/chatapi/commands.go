package chatapi

import (
	"github.com/RafilxTenfen/go-chat/cli"
	"github.com/ibraimgm/libcmd"
	"github.com/rhizomplatform/fs"
	"github.com/rhizomplatform/log"
)

const (
	port = uint16(4444)
)

// terminalHandler exists only to "pass" the
// cli.Terminal along without requiring a global variable and
type terminalHandler struct {
	term cli.Terminal
}

// Main handles all command line interfaces
func Main(args []string) int {
	h := &terminalHandler{cli.Terminal{}}
	app := libcmd.NewApp("Chat API Server", "Chat API that wraps publish users methods")
	app.Options.HelpOutput = h.term.Out()
	app.Options.StrictOperands = true

	// loggin setup
	log.Setup(fs.Path("./log/api"), "apilg", 60*24, 60)
	defer log.TearDown()
	log.SetStdoutLevel(log.LevelDebug)

	app.Command("run", "Run the API service", func(cmd *libcmd.Cmd) {
		cmd.Uint16("port", 'p', port, "Port of the API")
		cmd.Run(h.run)
	})

	// root - no command specified
	app.Run(func(cmd *libcmd.Cmd) error {
		app.PrintHelp(h.term.Out())
		return nil
	})

	if err := app.ParseArgs(args); err != nil {
		h.term.Printf("%v\n", cli.UserError(err))
		return 1
	}

	return 0
}
