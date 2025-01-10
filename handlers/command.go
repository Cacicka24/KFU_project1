package handlers

import (
	"context"
	//библиотека для бота
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	//для логов
	"github.com/rs/zerolog/log"
)

func Start(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Приветственное сообщение или информационное",
	})
	log.Info().Msg("Пользователь запустил бота")
}

func Help(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Список команд для взаимодействия с ботом: \n /start - запускает бота \n /help - список команд",
	})
	log.Info().Msg("Пользователь воспользовался командой /help")
}
