package main

import (
	"log"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/zserge/lorca"
)

func main() {
	go runServer()

	err := runUI()
	if err != nil {
		log.Fatal(err)
	}
}

func runUI() error {
	ui, err := lorca.New("", "", 320, 480)
	if err != nil {
		return err
	}

	ui.Load("http://localhost:9000")
	<-ui.Done()
	return nil
}

func runServer() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./app/build", true)))
	router.Run(":9000")
}
