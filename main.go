package main

import (
	"TG_bot/filters"
	"TG_bot/handlers"
	"context"
	"os"
	"os/signal"

	//библиотека для бота
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	//библиотека для логирования
	"github.com/rs/zerolog/log"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New("7743496303:AAFdoSLRcTEyXSknBQbPtBwCbgbBlfrpN4U", opts...)
	if err != nil {
		log.Panic()
	}

	//хендлер отвечающий за привествие по команде "/start"
	b.RegisterHandlerMatchFunc(filters.IsStart, handlers.Start)
	//думаю, что так или иначе в каждом боте должна быть команда /help
	b.RegisterHandlerMatchFunc(filters.IsHelp, handlers.Help)

	b.Start(ctx)
}

//Дефолтный обработчик сообщений, отвечает на каждое сообщение "Нет такой команды", кроме тех сообщений,
//которые прописанны в других хендлерах

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Нет такой команды",
	})
	log.Info().Msg("Пользователь ввел непредусмотренное сообщение")
}
