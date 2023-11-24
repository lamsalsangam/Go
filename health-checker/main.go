package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Health Checker",
		Usage: "Tool thak check if the website is running or not",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "domain",
				// Aliases are used as "go run -d ______.com" ≡ "go run --domain _____.com" where _____.com is replace the "d".
				Aliases:  []string{"d"},
				Usage:    "Domain name to check",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port number to check",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			if c.String("port") == "" {
				port = "80"
			}
			status := Check(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}
	// os.Args helps to pass the arguments from the command line
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
