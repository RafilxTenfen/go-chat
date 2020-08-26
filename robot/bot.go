package robot

// Bot represents a Robot to listen and publish at N queues
type Bot struct {
	queues   map[string]uint16
	settings Settings
}
