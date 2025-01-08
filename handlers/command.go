package handlers

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func Start(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Привет",
	})
}

func Help(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Взаимодействие с ботом должно выполняться через команды или inline кнопки(пока что отсуствуют), Список команд для взаимодействия с ботом: \n /start - запускает бота \n /help - помощь по боту и список команд",
	})
}
