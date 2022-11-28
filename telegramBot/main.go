package main

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	pref := tele.Settings{
		Token:  "5555228041:AAHhFBdYy3c4pSzimKC7gx1_pFUeWbGBM8I",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})
	b.Handle(tele.OnText, func(c tele.Context) error {

		// Use full-fledged bot's functions
		// only if you need a result:
		// Instead, prefer a context short-hand:
		return c.Send("хули тут непонятного долбаеб)")
	})
	b.Handle(tele.OnPhoto, func(c tele.Context) error {
		// Photos only.
		return c.Send("Твоя картинка хуйня полная")
	})
	b.Start()
}
