package main

import (
	"fmt"
	"log"
	app2 "stay-server/app"
)

func init() {
	log.Print("start init")
}

func main() {
	var app = app2.App{}
	fmt.Print(app.Id)

}
