package filters

import "github.com/go-telegram/bot/models"

const (
	start = "/start"
)

func IsStart(update *models.Update) bool {
	return update.Message.Text == start
}
