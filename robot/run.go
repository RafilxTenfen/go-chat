package robot

// Run the robot that listen to all commands
func (b *Bot) Run() error {
	b.ConsumeAllQueues()

	shell := b.Shell()
	shell.Run()

	defer b.Exit()
	return nil
}
