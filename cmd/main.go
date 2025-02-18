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
	menu := &telebot.ReplyMarkup{ResizeKeyboard: true}
	btnStart := menu.Text("Go quiz")

	options := &telebot.ReplyMarkup{ResizeKeyboard: true}
	btnFirst := options.Text("1")
	btnSecond := options.Text("2")
	btnThird := options.Text("3")
	btnFourth := options.Text("4")

	menu.Reply(
		menu.Row(btnStart),
	)

	options.Reply(
		menu.Row(btnFirst, btnSecond, btnThird, btnFourth),
	)

	b.Handle("/start", func(c telebot.Context) error {

		return c.Send("Hello! This is Quiz Bot! Do you want to test your knowledge?", menu)
	})

	b.Handle(&btnStart, func(c telebot.Context) error {
		c.Send("Lets GO")

		question := "first question"

		return c.Send(question, options)
	})

	b.Start()
}
