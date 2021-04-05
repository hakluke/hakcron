package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/robfig/cron"
)

func main() {
	// parse CLI flags
	frequency := flag.String("f", "NONE", "how often would you like to run the command?")
	command := flag.String("c", "NONE", "what is the command you would like to run?")
	flag.Parse()

	// Do some sanity checks
	if *frequency == "NONE" || *command == "NONE" {
		usage()
	}

	freqString := "@" + *frequency

	log.Println("Executing", *command, *frequency, "\nPress return to exit.")

	// set up the cron
	c := cron.New()
	c.AddFunc(freqString, func() { execCommand(*command) })
	c.Start()
	fmt.Scanln() // prevent the program from exiting
	c.Stop()     // Stop the scheduler (does not stop any jobs already running).
	log.Println("Stopped scheduled task.")
}

func execCommand(command string) {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		log.Println("Error running command:", err)
		return
	}
	fmt.Println(string(output))
}

func usage() {
	fmt.Println(`
	hakcron allows you to run a command at specific intervals. 
	It was written with the intention of being able to quickly set up a cronjob in a tmux 
	session or similar, without having to actually edit the crontab.

	For example, to run a command daily, you could do:
	hakcron -f "daily" -c "curl hakluke.com/dostuff.php"

	To run a command hourly, you could do:
	hakcron -f "hourly" -c "curl hakluke.com/dostuff.php"

	-f can be set to yearly, montly, weekly, daily and hourly:

	Entry                  | Description                                | Equivalent To
	-----                  | -----------                                | -------------
	yearly (or annually)   | Run once a year, midnight, Jan. 1st        | 0 0 0 1 1 *
	monthly                | Run once a month, midnight, first of month | 0 0 0 1 * *
	weekly                 | Run once a week, midnight between Sat/Sun  | 0 0 0 * * 0
	daily (or @midnight)   | Run once a day, midnight                   | 0 0 0 * * *
	hourly                 | Run once an hour, beginning of hour        | 0 0 * * * *

	You can also use intervals like this:

	every 1h30m
	every 5s

	hakcron implements robfig's cron library, so for more details see here: 
	https://pkg.go.dev/github.com/robfig/cron#hdr-Intervals
	`)
	os.Exit(1)
}
