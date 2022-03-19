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

	//UPDTAE COMMAND
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)

	//add id
	updateId := updateCmd.String("id", "", "Youtube video id")

	//update title
	updateTitle := updateCmd.String("title", "", "Youtube video title")

	//update url
	updateUrl := updateCmd.String("url", "", "Youtube video url")

	//update imageUrl
	updateImageUrl := updateCmd.String("imageUrl", "", "Youtube video Image Url")

	//update description
	updateDesc := updateCmd.String("desc", "", "Youtube video Description")

	if len(os.Args) < 2 {
		fmt.Println("expected 'get or 'add' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		handleGet(getCmd, getAll, getId)
	case "add":
		handleAdd(addCmd, addId, addTitle, addUrl, addImageUrl, addDesc)
	case "update":
		handleUpdate(updateCmd, updateId, updateTitle, updateUrl, updateImageUrl, updateDesc)
	default:

	}

}

func checkFileExist() string {
	_, err := os.Stat("videos-updated.json")
	if err != nil {
		panic(err)
	}
	if os.IsNotExist(err) {
		return "videos.json"
	} else {
		return "videos-updated.json"
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
		pathname := checkFileExist()
		videos := getVideos(pathname)
		fmt.Printf("ID \t Title \t URL \t ImageURL \t Description \n")
		for _, video := range videos {
			fmt.Printf("%v \t %v \t %v \t %v \t %v \t \n", video.Id, video.Title, video.Url, video.Imageurl, video.Description)
		}
		return
	}

	//If --id
	if *id != "" {
		videos := getVideos("./videos.json")
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

func ValidateVideoInput(validateCmd *flag.FlagSet, id *string, title *string, imageUrl *string, url *string, desc *string) {

	//Ignore the add flag, take the rest (starting from -id)
	validateCmd.Parse(os.Args[2:])

	if *id == "" || *title == "" || *url == "" || *imageUrl == "" || *desc == "" {
		fmt.Print("Missing required fields \n")
		validateCmd.PrintDefaults()
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

	pathname := checkFileExist()
	videos := getVideos(pathname)
	videos = append(videos, video)

	saveVideos(videos)
}

func handleUpdate(updateCmd *flag.FlagSet, id *string, title *string, imageUrl *string, url *string, desc *string) {
	ValidateVideoInput(updateCmd, id, title, imageUrl, url, desc)
	//Find the video's id and update
	//We want to get videos-updated.json if it exists
	pathname := checkFileExist()
	videos := getVideos(pathname)

	for i, video := range videos {
		if *id == video.Id {
			//update video
			videos[i].Title = *title
			videos[i].Imageurl = *imageUrl
			videos[i].Url = *url
			videos[i].Description = *desc
		}
	}
	saveVideos(videos)
}
