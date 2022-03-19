# json-cli
A simple CLI written in Go that generates JSON file through commands

## Installation
* Install Go from https://go.dev/doc/install 
* Clone this repo
* Change to videos directory that contains go files
```bash
cd videos
```
* Execute command to play around with the CLI
``` bash
go run main.go videos.go --available commands--
```

## Available Commands
- [x] GET ALL
- [x] GET ONE
- [x] ADD
- [ ] UPDATE
- [ ] DELETE

### Get all data from videos.json file
```bash
go run main.go videos.go get --all
```
### Get data with specific id from videos.json file
```bash
go run main.go videos.go get --id "video id"
```

### Add data and create videos-updated.json file
```bash
go run main.go videos.go add --id "id" --title "title" --imageUrl "image url" --url "video url" --desc "video description"
```
