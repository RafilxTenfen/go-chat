package robot

// Run the robot that listen to all commands
func (b *Bot) Run(useShell bool) error {
	b.ConsumeAllQueues()

	if useShell {
		shell := b.Shell()
		shell.Run()
	} else {
		forever := make(chan bool)
		<-forever
	}

	defer b.Exit()
	return nil
}
