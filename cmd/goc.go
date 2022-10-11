package main

import (
	"log"
	"os"

	"github.com/LeanderGangso/goc"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func main() {
	info()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func info() {
	app.Name = "Google Calendar CLI"
	app.Usage = "A simple CLI for tracking hours into Google Calendar"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:   "setup",
			Usage:  "Setup Google calendar",
			Action: goc.GoogleSetup,
		},
		{
			Name:      "start",
			Aliases:   []string{"s"},
			Usage:     "Start tracking new task",
			ArgsUsage: "'Task name'",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "t",
					Usage: "Set start time for task (HH:MM)",
				},
			},
			Action: goc.StartTask,
		},
		{
			Name:    "end",
			Aliases: []string{"e"},
			Usage:   "End the currently tracked task",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "t",
					Usage: "Set end time for task (HH:MM)",
				},
			},
			Action: goc.EndTask,
		},
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "Update the current task",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "n",
					Usage: "Set new task name",
				},
				cli.StringFlag{
					Name:  "t",
					Usage: "Set new task time (HH:MM)",
				},
			},
			Action: goc.EditCurrentTask,
		},
		{
			Name:    "status",
			Aliases: []string{"st"},
			Usage:   "Get current task status.",
			Action:  goc.TaskStatus,
		},
	}
}