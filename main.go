package main

import (
	"log"
	"os"
	"time"

	tg "gopkg.in/telebot.v4"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	err := os.Setenv("TOKEN", "7743496303:AAFdoSLRcTEyXSknBQbPtBwCbgbBlfrpN4U")
	if err != nil {
		log.Fatalln("Ошибка создания переменной окружения", err)
	}
	pref := tg.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tg.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tg.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(c tg.Context) error {
		return c.Send("Привет!")
	})

	b.Start()
}
