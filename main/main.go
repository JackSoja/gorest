package main

import (
	"gorest"
)

func main() {
	app := 	.NewApp()
	app.Get("/json", func(r *gorest.HttpRequest, w gorest.HttpResponse) error {
		a := struct {
			Abc string `json:"abc"`
			Cba string `json:"cba"`
		}{"123", "321"}
		return w.WriteJson(a)
	})

	app.Error(func(err error, r *gorest.HttpRequest, w gorest.HttpResponse) {
		if e, ok := err.(gorest.NoFoundError); ok {
			w.Write([]byte(e.Error()))
		}
		if e, ok := err.(gorest.InternalError); ok {
			w.Write([]byte(e.Error()))
		}
	})
	app.Run(":8081")
}
