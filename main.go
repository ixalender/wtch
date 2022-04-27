package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gosuri/uilive"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "watch",
		Usage:     "Repeats specified command line",
		Version:   "1.0",
		UsageText: `watch [global options] help | "system command with args"`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "interval",
				Aliases: []string{"n"},
				Value:   "1000",
				Usage:   "interval in milliseconds",
			},
			&cli.StringFlag{
				Name:    "repeat",
				Aliases: []string{"r"},
				Value:   "0",
				Usage:   "times to repeat (default: 0 - infinity)",
			},
		},
		Action: func(c *cli.Context) error {
			command := ""
			if c.NArg() > 0 {
				command = c.Args().Get(0)
			} else {
				cli.ShowAppHelp(c)
				return nil
			}

			interval := c.Int("interval")
			repeat := c.Int("repeat")

			err := execCommand(command, interval, repeat)
			if err != nil {
				log.Println(err)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func execCommand(command string, interval, repeat int) error {
	log.Println(fmt.Sprintf("'%s' every %d millisecond", command, interval))

	outWriter := uilive.New()
	outWriter.Start()

	repeatCount := 0

	for range time.Tick(time.Millisecond * time.Duration(interval)) {
		var args []string = []string{}

		commandParts := strings.Split(command, " ")

		if len(commandParts) > 1 {
			args = commandParts[1:]
		}

		cmd := exec.Command(commandParts[0], args...)
		outBuff, _ := cmd.CombinedOutput()

		fmt.Fprintf(outWriter, fmt.Sprintf("%s", string(outBuff)))

		repeatCount++
		if repeatCount == repeat {
			break
		}
	}

	outWriter.Stop()

	return nil
}
