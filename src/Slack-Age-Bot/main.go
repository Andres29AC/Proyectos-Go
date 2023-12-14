package main

import (
	"fmt"
	"context"
	"log"
	"os"
	"strconv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Evenets")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
//NOTE: analyticsChannel es un canal que recibe eventos de comandos
func main() {
	os.Setenv("SLACK_BOT_TOKEN","xoxb-5474558710455-6349869728673-fbWgVXwdkHmxs9mLaBgu0KTo")
	os.Setenv("SLACK_APP_TOKEN","xapp-1-A06AKUUV5UY-6340273777092-a13dd2b1a11e7ec683f3ec1cf507dec93e469300c568c51e2b785c52f258c03c")
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator. Example: 'my yob is 2020'",	
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2023 - yob
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	//NOTE: WithContext() sirve para crear un contexto con un valor especifico
	//NOTE: Listen() es un metodo que recibe un contexto y un error 

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

//NOTE: Reply() sirve para responder al usuario que hizo la peticion 






