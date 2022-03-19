package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	//GET COMMAND
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	//get all video
	getAll := getCmd.Bool("all", false, "Get all videos")

	//get a video by id
	getId := getCmd.String("id", "", "Youtube video id")

	//ADD COMMAND
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	//add id
	addId := addCmd.String("id", "", "Youtube video id")

	//add title
	addTitle := addCmd.String("title", "", "Youtube video title")

	//add url
	addUrl := addCmd.String("url", "", "Youtube video url")

	//add imageUrl
	addImageUrl := addCmd.String("imageUrl", "", "Youtube video Image Url")

	//add description
	addDesc := addCmd.String("desc", "", "Youtube video Description")

	if len(os.Args) < 2 {
		fmt.Println("expected 'get or 'add' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		handleGet(getCmd, getAll, getId)
	case "add":
		handleAdd(addCmd, addId, addTitle, addUrl, addImageUrl, addDesc)
	default:

	}

}

func handleGet(getCmd *flag.FlagSet, all *bool, id *string) {
	//Parse command after get
	getCmd.Parse(os.Args[2:])

	//Input validation
	if !*all && *id == "" {
		fmt.Print("id is required or specify --all for all videos")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	//If --all
	if *all {
		//return all videos
		videos := getVideos()
		fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
		for _, video := range videos {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \t \n", video.Id, video.Title, video.Url, video.Imageurl, video.Description)
		}
		return
	}

	//If --id
	if *id != "" {
		videos := getVideos()
		id := *id

		for _, video := range videos {
			//Find video id
			if id == video.Id {
				fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
				fmt.Printf("%v \t %v \t %v \t %v \t %v \t \n", video.Id, video.Title, video.Url, video.Imageurl, video.Description)
			}
		}
		return
	}
}

func ValidateVideoInput(addCmd *flag.FlagSet, id *string, title *string, imageUrl *string, url *string, desc *string) {

	//Ignore the add flag, take the rest (starting from -id)
	addCmd.Parse(os.Args[2:])

	if *id == "" || *title == "" || *url == "" || *imageUrl == "" || *desc == "" {
		fmt.Print("all fields are required for adding a video")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
}

func handleAdd(addCmd *flag.FlagSet, id *string, title *string, imageUrl *string, url *string, desc *string) {
	ValidateVideoInput(addCmd, id, title, imageUrl, url, desc)

	//Create new video object
	video := video{
		Id:          *id,
		Title:       *title,
		Description: *desc,
		Imageurl:    *imageUrl,
		Url:         *url,
	}

	videos := getVideos()
	videos = append(videos, video)

	saveVideos(videos)
}
