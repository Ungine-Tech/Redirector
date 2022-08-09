package main

import (
	"flag"
	"fmt"

	"github.com/Ungine-Tech/redirector/bootstrap"
	"github.com/Ungine-Tech/redirector/router"
	"github.com/gin-gonic/gin"
)

var (
	target string
	port   string
	dev    bool
)

func init() {
	flag.StringVar(&target, "t", "http://localhost", "target url")
	flag.StringVar(&port, "p", ":80", "port")
	flag.BoolVar(&dev, "D", false, "dev mode")
	flag.Parse()

	if !dev {
		gin.SetMode(gin.ReleaseMode)
	}

	bootstrap.Init(target)
}

func main() {
	r := router.InitRouter()

	fmt.Printf("Server is running on port %s", port)
	err := r.Run(port)
	if err != nil {
		fmt.Print("error: ", err)
		return
	}
}
