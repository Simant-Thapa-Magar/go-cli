package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// command for getting videos
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	// input for get videos command
	// all videos
	getAll := getCmd.Bool("all", false, "Get all vidos")

	// video by id
	getId := getCmd.String("id", "", "Video ID")

	// command args validation
	if len(os.Args) < 2 {
		fmt.Println("expected 'get' or 'save' command")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		// handle get command
		handleGet(getCmd, getAll, getId)
		break
	case "save":
		// handle save command
	default:
	}
}

func handleGet(getCmd *flag.FlagSet, getAll *bool, getId *string) {
	fmt.Println("I will handle get command")
}
