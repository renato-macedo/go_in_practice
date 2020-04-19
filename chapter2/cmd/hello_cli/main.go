package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "hello_cli"
	app.Usage = "Print hello world"

	app.Flags = []cli.Flag{

		&cli.StringFlag{
			Name:  "name, n",
			Value: "World",
			Usage: "Who to say hello to.",
		},
		&cli.BoolFlag{
			Name:  "spanish, s",
			Usage: "Use Spanish Language.",
		},
	}

	app.Action = func(c *cli.Context) error {

		name := c.GlobalString("name")
		if c.Bool("spanish") == true {
			fmt.Printf("Hola %s!\n", name)
		} else {
			fmt.Printf("Hello %s!\n", name)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
