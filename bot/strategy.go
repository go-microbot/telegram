package bot

import "github.com/go-microbot/telegram/models"

const defaultUpdatesLimit = 20

// UpdatesStrategy represents a strategy interface
// for determining how a Bot will receive new messages.
type UpdatesStrategy interface {
	Updates() chan models.Update
	Errors() chan error
	Listen()
	Stop()
}
