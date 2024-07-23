package main

import (
	"fmt"

	"github.com/Raghavk4u/discord-ping/bot"
	"github.com/Raghavk4u/discord-ping/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	bot.Start()
	//make(chan struct{}) creates a channel that can send and receive messages of type struct{}.
	<-make(chan struct{}) //the program is waiting to receive a message from the newly created channel.

	return //Since the channel is empty and nothing will ever be sent to it, the program will wait forever.
	//This effectively keeps the program running without doing anything else.
}
