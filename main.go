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
		break
	default:
		fmt.Println("expected 'get' or 'save' command")
	}
}

func handleGet(getCmd *flag.FlagSet, all *bool, id *string) {
	getCmd.Parse(os.Args[2:])

	if *all == false && *id == "" {
		fmt.Println("Specify id or all to retrieve all videos")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *all {
		videos := getVideos()
		fmt.Println("Videos ")
		fmt.Printf("ID \tVIDEO TITLE \tIMAGE PREVIEW \tVIDEO URL \tVIDEO DESCRIPTION \n")
		for _, video := range videos {
			fmt.Printf("%s \t%s \t%s \t%s \t%s \n", video.Id, video.Title, video.PreviewUrl, video.Url, video.Description)
		}
	}

	if *id != "" {

	}

}
