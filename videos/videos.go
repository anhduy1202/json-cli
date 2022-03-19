package main

import (
	"encoding/json"
	"io/ioutil"
)

type video struct {
	Id          string
	Title       string
	Description string
	Imageurl    string
	Url         string
}

func getVideos(filepath string) (videos []video) {
	fileBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(fileBytes, &videos) //covert fileBytes to json
	if err != nil {
		panic(err)
	}
	return videos
}

func saveVideos(videos []video) {
	videoBytes, err := json.Marshal(videos)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("./videos-updated.json", videoBytes, 0644)
	if err != nil {
		panic(err)
	}
}
