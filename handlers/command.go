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
		Text:   "Бот запущен! Ознакомиться со списком всех команд можно по команде /help. Если вы не авторизованны, то это стоит сделать нажав сюда: /login", //FIXME
	})
	log.Info().Msg("Пользователь запустил бота")
}

func Help(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Список команд для взаимодействия с ботом: \n /start - запускает бота \n /help - список команд \n /login - для регистрации",
	})
	log.Info().Msg("Пользователь воспользовался командой /help")
}

func Login(ctx context.Context, b *bot.Bot, update *models.Update) {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "Яндекс ID", CallbackData: "Авторизация через Яндекс ID"},
			}, {
				{Text: "Github", CallbackData: "Авторизация через Github"},
			}, {
				{Text: "Код", CallbackData: "Авторизация через код"},
			},
		},
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Выберите способ авторизации:",
		ReplyMarkup: kb,
	})
	log.Info().Msg("Пользователь воспользовался командой /login")
}
