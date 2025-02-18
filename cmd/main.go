package main

import (
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pref := telebot.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	slog.Info("starting bot")

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c telebot.Context) error {
		return c.Send("Hello!")
	})

	b.Handle("/info", func(c telebot.Context) error {
		return c.Send("This is quiz bot")
	})

	b.Start()
}
