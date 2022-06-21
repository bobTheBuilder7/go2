package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/jbowtie/gokogiri"
	"github.com/jbowtie/gokogiri/xpath"
)

type Wallet struct {
	balance_in_btc string
	balance_in_usd string
}

func fetch_address(address string) (Wallet, error) {
	response, err := http.Get(fmt.Sprintf("https://www.blockchain.com/btc/address/%s", address))
	if err != nil {
		return Wallet{}, errors.New(string(response.StatusCode))
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return Wallet{}, err
	}
	doc, err := gokogiri.ParseHtml(body)
	if err != nil {
		return Wallet{}, err
	}
	defer doc.Free()
	xps := xpath.Compile("/html/body/div[1]/div[3]/div[2]/div/div/div[1]/div/div[3]/span")
	ss, err := doc.Search(xps)
	if err != nil {
		return Wallet{}, err
	}
	data := ss[0].String()
	ind := strings.Index(data, "address is")
	data = data[ind+len("address is") : len(data)-8]
	data_sl := strings.Split(data, "BTC")

	data_sl[0] = data_sl[0][1 : len(data_sl[0])-1]
	data_sl[1] = data_sl[1][2 : len(data_sl[1])-2]
	w := Wallet{
		balance_in_btc: data_sl[0],
		balance_in_usd: data_sl[1]}
	return w, nil
}
