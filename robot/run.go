package robot

// Run the robot that listen to all commands
func (b *Bot) Run() error {
	b.ConsumeAllQueues()

	// just to stay alive
	forever := make(chan bool)
	<-forever

	return nil
}
