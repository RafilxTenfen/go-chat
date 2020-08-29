package user

import (
	"github.com/RafilxTenfen/go-chat/cli"
	"github.com/ibraimgm/libcmd"
	"github.com/rhizomplatform/fs"
	"github.com/rhizomplatform/log"
)

const (
	defaultEmail = "anyexample@gmail.com"
)

// terminalHandler exists only to "pass" the
// cli.Terminal along without requiring a global variable and
type terminalHandler struct {
	term cli.Terminal
}

// Main handles all command line interfaces
func Main(args []string) int {
	h := &terminalHandler{cli.Terminal{}}
	app := libcmd.NewApp("User CLI", "Basic User reads and publish into RabbitMQ queues")
	app.Options.HelpOutput = h.term.Out()
	app.Options.StrictOperands = true

	// loggin setup
	log.Setup(fs.Path("./log/user"), "userlg", 60*24, 60)
	defer log.TearDown()
	log.SetStdoutLevel(log.LevelDebug)

	app.Command("add", "Add a new user into the database", func(cmd *libcmd.Cmd) {
		cmdUserParams(cmd)
		cmd.Run(h.add)
	})

	app.Command("login", "Login with a user", func(cmd *libcmd.Cmd) {
		cmdUserParams(cmd)
		cmd.Run(h.login)
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

func cmdUserParams(cmd *libcmd.Cmd) {
	cmd.String("email", 'e', defaultEmail, "User email")
	cmd.String("password", 'p', "", "User passwor")
}
