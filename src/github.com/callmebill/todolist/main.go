// main
package main

import (
	"github.com/callmebill/todolist/app"
	"github.com/codegangsta/negroni"
)

func main() {
	router := app.NewRouter()
	midware := negroni.New()
	midware.UseHandler(router)
	midware.Run(":8888")
}
