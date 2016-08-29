package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/urfave/cli/altsrc"
)

func main() {
	myApp := cli.NewApp()
	myApp.Name = "kcptun"
	myApp.Usage = "kcptun client"
	myApp.Flags = []cli.Flag{
		altsrc.NewStringFlag(cli.StringFlag{
			Name:  "localaddr,l",
			Value: ":12948",
			Usage: "local listen address",
		}),
		altsrc.NewStringFlag(cli.StringFlag{
			Name:  "remoteaddr, r",
			Value: "vps:29900",
			Usage: "kcp server address",
		}),
		cli.StringFlag{
			Name:  "c", // configuration file support
			Usage: "config.json file support",
			Value: "",
		},
	}
	myApp.Action = func(c *cli.Context) error {
		log.Println("remote address:", c.String("remoteaddr"))
		log.Println("local address:", c.String("localaddr"))
		return nil
	}

	myApp.Before = altsrc.InitInputSourceWithContext(myApp.Flags, NewJSONSourceFromFlagFunc("c"))
	myApp.Run(os.Args)
}
