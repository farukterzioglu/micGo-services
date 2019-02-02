package models

// CommandMessage represents messages come from producers
type CommandMessage struct {
	CommandType string
	CommandData []byte
}
