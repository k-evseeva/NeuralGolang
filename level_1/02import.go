package main

import (
	foo "fmt"
	bar "net/http"
)

func main_02() {
	foo.Println("Hello, user!")
	resp, err := bar.Get("https://translate.google.com/?sl=ru&tl=en&op=translate")
	if err != nil {
		foo.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	foo.Println("HTTP response status:", resp.Status)
}
