package main

import (
	"encoding/json"
	"io/ioutil"
)

type video struct {
	Id          string
	Title       string
	Description string
	PreviewUrl  string
	Url         string
}

func getVideos() (videos []video) {
	fileBytes, err := ioutil.ReadFile("./videos.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &videos)

	if err != nil {
		panic(err)
	}

	return videos
}

func saveVideos(video []video) {

}
