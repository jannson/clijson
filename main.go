package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	myApp := cli.NewApp()
	myApp.Name = "kcptun"
	myApp.Usage = "kcptun client"
	myApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "localaddr,l",
			Value: ":12948",
			Usage: "local listen address",
		},
		cli.StringFlag{
			Name:  "remoteaddr, r",
			Value: "vps:29900",
			Usage: "kcp server address",
		},
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

	path := "config.json"
	file, err := os.Open(path) // For read access.
	if err != nil {
		return
	}
	defer file.Close()

	var config map[string]interface{}
	if err = json.NewDecoder(file).Decode(&config); err != nil {
		log.Printf("error:%v\n", err)
		return
	}
	for _, f := range myApp.Flags {
		//fValue := reflect.ValueOf(f)
		fValue := reflect.ValueOf(f).Elem()
		//fValue := reflect.Indirect(reflect.ValueOf(f))
		//fValue := reflect.Indirect(reflect.ValueOf(f).Elem())
		typeOfT := fValue.Type()
		for i := 0; i < fValue.NumField(); i++ {
			ff := fValue.Field(i)
			fmt.Printf("%d: %s %s = %v %v\n", i,
				typeOfT.Field(i).Name, ff.Type(), ff.Interface(), ff.CanSet())
		}

		varValue := fValue.FieldByName("Name")
		if !varValue.IsValid() {
			return
		}
		varStr := varValue.String()
		vars := strings.Split(varStr, ",")
		varStr = strings.TrimSpace(vars[0])
		if varJson, err := jsonGetValue(varStr, config); err == nil {
			log.Printf("varValue=%s varJson=%v\n", varStr, varJson)
		}
	}

	//myApp.Before = altsrc.InitInputSourceWithContext(myApp.Flags, NewJSONSourceFromFlagFunc("c"))
	myApp.Run(os.Args)
}
