package bot

import (
	"strings"

	"github.com/RafilxTenfen/go-chat/rabbit"
	"github.com/RafilxTenfen/go-chat/robot"
	"github.com/ibraimgm/libcmd"
	"github.com/rhizomplatform/log"
)

func (t *terminalHandler) run(cmd *libcmd.Cmd) error {
	strName := cmd.GetString("queues")
	queues := strings.Split(*strName, ",")

	st := rabbit.LoadSettingsFromEnv()
	log.With(log.F{
		"Settings": st,
		"Queues":   queues,
	}).Debug("Settings Loaded")

	bot, err := robot.NewBot(st, queues)
	if err != nil {
		return err
	}

	b := cmd.GetBool("useshell")
	return bot.Run(*b)
}
