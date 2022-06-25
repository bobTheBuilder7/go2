package main

import (
	"fmt"
	"log"

	_ "os"
	"time"

	"github.com/go-co-op/gocron"
	tele "gopkg.in/telebot.v3"
)

var (
	firstPage     = &tele.ReplyMarkup{}
	fToFirstPage  = firstPage.Data("· 1 ·", "b1")
	fToSecondPage = firstPage.Data("2", "b2")
)

// Creating first page keyboard
func init() {
	firstPage.Inline(
		firstPage.Row(fToFirstPage, fToSecondPage),
	)
}

var (
	secondPage    = &tele.ReplyMarkup{}
	sToFirstPage  = secondPage.Data("1", "b3")
	sToSecondPage = secondPage.Data("· 2 ·", "b4")
)

// Creating second page keyboard
func init() {
	secondPage.Inline(
		secondPage.Row(sToFirstPage, sToSecondPage),
	)
}

var btc_price string

// Go Cron for checking btc price once in 60 seconds
func init() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(60).Seconds().Do(func() {
		btc_price = get_btc_price()
		log.Println(btc_price)
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

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Первая страница", firstPage)
	})

	b.Handle(&fToFirstPage, func(c tele.Context) error {
		return c.Edit("Первая страница", firstPage)
	})
	b.Handle(&fToSecondPage, func(c tele.Context) error {
		return c.Edit("Вторая страница", secondPage)
	})
	b.Handle(&sToFirstPage, func(c tele.Context) error {
		return c.Edit("Первая страница", firstPage)
	})
	b.Handle(&sToSecondPage, func(c tele.Context) error {
		return c.Edit("Вторая страница", secondPage)
	})

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

	b.Handle("/price", func(c tele.Context) error {
		return c.Send(btc_price)
	})
	b.Start()
}
