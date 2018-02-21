package main

import "net/http"

func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
