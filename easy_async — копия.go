package main

import (
	"fmt"
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


func main() {
	n := 1000
	wg.Add(n)
	for i := 0; i < n; i++ {
		go fetch()
	}
wg.Wait()
fmt.Print(globallist)

}
