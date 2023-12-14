package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN","xoxb-5474558710455-6328726628551-34mrG1i7qk84P95cfF1kZk0Z")
	//NOTE: CHANNEL_IO en este caso es el id del canal donde se quiere subir el archivo
	os.Setenv("CHANNEL_IO","C05ED2U81K4")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_IO")}
	fileArr := []string{"RSL.pdf"}

	for i := 0; i < len(channelArr); i++{
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return 
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
