package main

import (
	"fmt"
	"log"

	_ "os"
	"time"
	"github.com/go-co-op/gocron"
	tele "gopkg.in/telebot.v3"
)
var btc_price string


func init() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(60).Seconds().Do(func(){
		btc_price = get_btc_price()
	})
	s.StartAsync()
}

func main() {
	pref := tele.Settings{
		Token:  "5571023482:AAHglns0AvWfe6ysclpZG9JfE3NKK0lJMGw",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/btc", func(c tele.Context) error {
		tags := c.Args()
		if len(tags) == 1 {
			w, err := fetch_address(tags[0])
			if err != nil {
				return c.Send("No Such Address")
			}
			msg := fmt.Sprintf("BTC:%s  ---  USD:%s", w.balance_in_btc, w.balance_in_usd)
			return c.Send(msg)
		}
		return c.Send("No Address")
	})

	b.Handle("price", func(c tele.Context) error {
		return c.Send(btc_price)
	})
	b.Start()
}
