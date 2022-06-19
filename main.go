package main

import (
	"fmt"
	"log"

	_ "os"
	"time"

	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func main() {
	pref := tele.Settings{
		Token:  "5571023482:AAHglns0AvWfe6ysclpZG9JfE3NKK0lJMGw",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}
	b.Use(middleware.Logger())
	b.Handle("/w", func(c tele.Context) error {
		tags := c.Args()
		if len(tags) == 1 {
			w, err := Fetch_address(tags[0])
			if err != nil {
				return c.Send("No Such Address")
			}
			msg := fmt.Sprintf("BTC:%s  ---  USD:%s", w.balance_in_btc, w.balance_in_usd)
			return c.Send(msg)
		}
		return c.Send("No Address")
	})
	b.Start()
}
