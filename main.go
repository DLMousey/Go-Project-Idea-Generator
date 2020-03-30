package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"math/rand"
	"os"
	"time"
)

var app = cli.NewApp()

func main() {
	app.Name = "App Ideas Generator"
	app.Usage = "A CLI tool for generating app ideas"
	app.Authors = []*cli.Author{
		&cli.Author{
			Name: "Mousey",
			Email: "tom@enderstudy.com",
		},
	}
	app.Version = "1.0.0"
	app.Commands = []*cli.Command{
		{
			Name: "generate",
			Aliases: []string{"g"},
			Usage: "Spits out an app idea",
			Action: func(c *cli.Context) error {
				file, err := os.Open("app-ideas.txt")
				if err != nil {
					log.Fatal(err)
					return nil
				}

				defer file.Close()

				scanner := bufio.NewScanner(file)
				result := []string{}

				count := 0
				for scanner.Scan() {
					count++
					line := scanner.Text()
					result = append(result, line)
				}

				rand.Seed(time.Now().UnixNano())
				randomInt := rand.Intn(count)
				fmt.Println(result[randomInt])

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}