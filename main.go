package main

import (
	"flag"
	"fmt"

	"github.com/Ungine-Tech/redirector/bootstrap"
	"github.com/Ungine-Tech/redirector/router"
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
