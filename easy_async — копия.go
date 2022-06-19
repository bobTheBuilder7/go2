package main

import (
	"net/http"
	"sync"
)
var wg = sync.WaitGroup{}
var globallist []int

func fetch() {
	
	response, err := http.Get("https://coincap.io/assets/bitcoin")
	if err != nil{
		panic(err)
	}
	globallist = append(globallist, response.StatusCode)
	wg.Done()
}