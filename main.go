package main

import (
	"fmt"
)

func main() {
	w, err := Fetch_address("34xp4vRoCGJym3xR7yCVPFHoCNxv4Twseo")
	if err != nil {
		panic(err)
	}
	fmt.Println(w)
}
