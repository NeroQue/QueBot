package main

import (
	"fmt"
	"os"

	"github.com/Goscord/goscord/goscord"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/gateway"
	"github.com/Goscord/goscord/goscord/gateway/event"

	"github.com/joho/godotenv"
)

var client *gateway.Session

func main() {
	fmt.Println("Starting...")
	godotenv.Load()

	client := goscord.New(&gateway.Options{
		Token:   os.Getenv("BOT_TOKEN"),
		Intents: gateway.IntentGuilds | gateway.IntentGuildMessages | gateway.IntentGuildMembers | gateway.IntentGuildVoiceStates,
	})

	client.On(event.EventReady, func() {
		fmt.Println("Logged in as " + client.Me().Tag())
	})

	client.On(event.EventMessageCreate, func(msg *discord.Message) {
		fmt.Println(msg.Content)
		if msg.Content == "ping" {
			client.Channel.SendMessage(msg.ChannelId, "Pong ! 🏓")
		}
	})

	client.Login()

	select {}
}
