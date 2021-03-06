package main

import (
	//"bufio"
	//"bytes"
	//"flag"
	"fmt"
	//"io/ioutil"
	//"net/http"
	//"net/url"
	"os"
	//"strings"
	//"text/template"
	"time"
)

import (
	//"github.com/hashicorp/hcl"
	"github.com/mattn/go-colorable"
	//"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

import (
  "github.com/spatialcurrent/sgol-cli/sgol"
)

var SGOL_CLI_VERSION = "0.0.1"

func main() {

	start := time.Now()

	if len(os.Args) < 2 {
		fmt.Println("Missing subcommand")
		fmt.Println("Run \"sgol help\" for command line options.")
		os.Exit(1)
	}

	config, err := sgol.LoadConfig(os.Getenv("SGOL_CONFIG_PATH"), true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cmd := os.Args[1]

	commands := map[string]sgol.Command{}
	//commands["help"] = &sgol.HelpCommand{}
	commands["queries"] = sgol.NewQueriesCommand(config)
	commands["formats"] = sgol.NewFormatsCommand(config)
	commands["exec"] = sgol.NewExecCommand(config)
	//commands["lint"] = &sgol.LintCommand{}

	if cmd == "help" || cmd == "--help" || cmd == "-h" || cmd == "-help" {
		fmt.Println("SGOL CLI - Version " + SGOL_CLI_VERSION)
		fmt.Println("Commands:")
		for x, _ := range commands {
			fmt.Println("\t- sgol " + x)
		}
		os.Exit(0)
	}

	if _, ok := commands[cmd]; !ok {
		fmt.Println("Error: Unknown subcommand")
		fmt.Println("Run \"sgol help\" for command line options.")
		os.Exit(1)
	}

	err = commands[cmd].Parse(os.Args[2:])
	if err != nil {
		fmt.Println(err)
		fmt.Println("Run \"sgol "+cmd+" -help\" for command line options.")
		os.Exit(1)
	}

	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())
	var log = logrus.New()

	err = commands[cmd].Run(log, start, SGOL_CLI_VERSION)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
