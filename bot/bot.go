package bot

import (
	"fmt"

	"github.com/Raghavk4u/discord-ping/config"
	"github.com/bwmarrin/discordgo"
)

var botID string
var goBot *discordgo.Session

func Start(){

	goBot, err := discordgo.New("Bot " + config.Token)       //in the discordgo.New function we have to add one space after the string "Bot" for no reason but necessary

	if err != nil{
		fmt.Println(err.Error())
	   return
	}

	u, err := goBot.User("@me")

	if err != nil{
		fmt.Println(err.Error())
	}

	botID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if (err != nil){
		fmt.Println(err.Error())
	}

	fmt.Println("Bot is running ...")
}


func messageHandler(s *discordgo.Session , m *discordgo.MessageCreate){
	if m.Author.ID == botID{
		return
	}

	if(m.Content == "ping"){
      	_,_	= s.ChannelMessageSend(m.ChannelID, "Pong!")
	}


}