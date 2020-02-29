package main

import (
	"fmt"
	"net/http"

	"gopkg.in/macaron.v1"
)

func main() {
	ma := macaron.Classic()
	ma.Get("/", handler)
	ma.Get("/*", handler)

	// ma.Run()

	fmt.Println("Server is running...")
	fmt.Println(http.ListenAndServe("localhost:8000", ma))
}

func handler(ctx *macaron.Context) string {
	return "You request is " + ctx.Req.RequestURI
}
