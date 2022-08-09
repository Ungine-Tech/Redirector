package main

import (
	"flag"
	"fmt"
	"github.com/Ungine-Tech/redirector/bootstrap"
	"github.com/Ungine-Tech/redirector/router"
	"github.com/gin-gonic/gin"
)

var (
	target      string
	port        string
	dev         bool
	enableHttps bool
	httpsPort   string
	pemFile     string
	keyFile     string
)

func init() {
	flag.BoolVar(&dev, "D", false, "dev mode")
	flag.StringVar(&target, "t", "http://localhost", "target url")

	flag.StringVar(&port, "p", ":80", "port")

	flag.BoolVar(&enableHttps, "s", false, "enable https")
	flag.StringVar(&httpsPort, "sp", ":443", "https port")
	flag.StringVar(&pemFile, "pem", "./cert.pem", "pem file")
	flag.StringVar(&keyFile, "key", "./cert.key", "key file")

	flag.Parse()

	if !dev {
		gin.SetMode(gin.ReleaseMode)
	}

	bootstrap.Init(target)
}

func main() {
	r := router.InitRouter()

	go func(port string) {
		if enableHttps {
			if pemFile == "" || keyFile == "" {
				panic("pem file or key file is empty")
				return
			}

			fmt.Printf("TLS Server is running on port %s\n", port)
			err := r.RunTLS(port, pemFile, keyFile)
			if err != nil {
				fmt.Print("error: ", err)
				return
			}
		}
	}(httpsPort)

	fmt.Printf("Server is running on port %s\n", port)
	err := r.Run(port)
	if err != nil {
		fmt.Print("error: ", err)
		return
	}
}
