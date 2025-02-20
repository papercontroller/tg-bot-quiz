package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v4"
)

var (
	current int
	score   int
)

var questions = []string{
	"Question 1",
	"Question 2",
	"Question 3",
	"Question 4",
	"Question 5",
	"Question 6",
	"Question 7",
	"Question 8",
	"Question 9",
	"Question 10",
}

var options = []string{
	"1",
	"2",
	"3",
	"4",
}

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

	menu.Reply(
		menu.Row(btnStart),
	)

	b.Handle("/start", func(c telebot.Context) error {

		return c.Send("Hello! This is Quiz Bot! Do you want to test your knowledge?", menu)
	})

	b.Handle(&btnStart, func(c telebot.Context) error {
		c.Send("Lets Go!")
		sendQuestion(b, int(c.Chat().ID))
		return nil

	})

	b.Handle(telebot.OnText, func(c telebot.Context) error {
		score++
		current++
		sendQuestion(b, int(c.Chat().ID))
		return nil
	})

	b.Start()
}

func sendQuestion(bot *telebot.Bot, chatID int) {
	if current >= len(questions) {
		bot.Send(&telebot.Chat{ID: int64(chatID)}, fmt.Sprintf("Your score is :%d", score))
		return
	}

	var buttons []telebot.ReplyButton
	for _, option := range options {
		buttons = append(buttons, telebot.ReplyButton{Text: fmt.Sprintf("%v", option)})
	}
	keyboard := &telebot.ReplyMarkup{
		ReplyKeyboard:  [][]telebot.ReplyButton{buttons},
		ResizeKeyboard: true,
	}

	question := questions[current]
	bot.Send(&telebot.Chat{ID: int64(chatID)}, question, &telebot.SendOptions{ReplyMarkup: keyboard})

}
