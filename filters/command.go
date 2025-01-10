package filters

import "github.com/go-telegram/bot/models"

const (
	start = "/start"
	help  = "/help"
	login = "/login"
)

func IsStart(update *models.Update) bool {
	return update.Message != nil && update.Message.Text == start
}
func IsHelp(update *models.Update) bool {
	return update.Message != nil && update.Message.Text == help
}
func IsLogin(update *models.Update) bool {
	return update.Message != nil && update.Message.Text == login
}
