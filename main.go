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

	// command for saving video
	saveCmd := flag.NewFlagSet("save", flag.ExitOnError)

	// input for save video command
	// id
	addId := saveCmd.String("id", "", "Video ID")

	// title
	addTitle := saveCmd.String("title", "", "Video Title")

	// description
	addDescription := saveCmd.String("description", "", "Video Description")

	// preview url
	addPreview := saveCmd.String("previewUrl", "", "Video Image URL")

	// url
	addUrl := saveCmd.String("url", "", "Video URL")

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
		handleSave(saveCmd, addId, addTitle, addDescription, addPreview, addUrl)
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
		videos := getVideos()
		found := false
		for _, video := range videos {
			if video.Id == *id {
				found = true
				fmt.Println("Video ")
				fmt.Printf("ID \tVIDEO TITLE \tIMAGE PREVIEW \tVIDEO URL \tVIDEO DESCRIPTION \n")
				fmt.Printf("%s \t%s \t%s \t%s \t%s \n", video.Id, video.Title, video.PreviewUrl, video.Url, video.Description)
			}
		}

		if !found {
			fmt.Println("No video found with id ", *id)
		}
	}

}

func handleSave(saveCmd *flag.FlagSet, id *string, title *string, description *string, previewUrl *string, url *string) {
	saveCmd.Parse(os.Args[2:])

	if *id == "" {
		fmt.Println("Provide video id for saving")
		saveCmd.PrintDefaults()
		os.Exit(1)
	}

	if *title == "" {
		fmt.Println("Provide video title for saving")
		saveCmd.PrintDefaults()
		os.Exit(1)
	}

	if *description == "" {
		fmt.Println("Provide video description for saving")
		saveCmd.PrintDefaults()
		os.Exit(1)
	}

	if *previewUrl == "" {
		fmt.Println("Provide video preview url for saving")
		saveCmd.PrintDefaults()
		os.Exit(1)
	}

	if *url == "" {
		fmt.Println("Provide video url for saving")
		saveCmd.PrintDefaults()
		os.Exit(1)
	}

	video := video{
		Id:          *id,
		Title:       *title,
		Description: *description,
		PreviewUrl:  *previewUrl,
		Url:         *url,
	}

	videos := getVideos()
	videos = append(videos, video)

	saveVideos(videos)
}
