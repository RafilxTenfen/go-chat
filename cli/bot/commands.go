package bot

import (
	"github.com/RafilxTenfen/go-chat/cli"
	"github.com/ibraimgm/libcmd"
	"github.com/rhizomplatform/fs"
	"github.com/rhizomplatform/log"
)

const (
	queues = "mqueue,otherqueue"
)

// terminalHandler exists only to "pass" the
// cli.Terminal along without requiring a global variable and
type terminalHandler struct {
	term cli.Terminal
}

// Main handles all command line interfaces
func Main(args []string) int {
	h := &terminalHandler{cli.Terminal{}}
	app := libcmd.NewApp("Bot Service", "Bot that reads and publish into RabbitMQ queues")
	app.Options.HelpOutput = h.term.Out()
	app.Options.StrictOperands = true

	// loggin setup
	log.Setup(fs.Path("./log/bot"), "bot service", 30, 60)
	defer log.TearDown()
	log.SetStdoutLevel(log.LevelDebug)

	app.Command("run", "Run the bot service", func(cmd *libcmd.Cmd) {
		cmd.String("queues", 'q', queues, "Queues name comma separated")
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
