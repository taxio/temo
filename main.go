package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/zserge/lorca"
)

var staticDir = "./app/build"

func main() {
	fmt.Println(staticDir)
	go runServer()

	err := runUI()
	if err != nil {
		log.Fatal(err)
	}
}

// app/build/以下を配信するサーバー
func runServer() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile(staticDir, true)))
	err := router.Run(":9000")
	if err != nil {
		log.Fatal(err)
	}
}

// lorca起動
func runUI() error {
	ui, err := lorca.New("", "", 320, 480)
	if err != nil {
		return err
	}

	ui.Load("http://localhost:9000")
	<-ui.Done()
	return nil
}
