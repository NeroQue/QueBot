package main

import (
	"fmt"
	"os"

	"github.com/Goscord/goscord/goscord"
	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/gateway"
	"github.com/Goscord/goscord/goscord/gateway/event"

	"github.com/NeroQue/QueBot/command"
	myEvent "github.com/NeroQue/QueBot/event"

	"github.com/joho/godotenv"
)

var (
	client *gateway.Session
	cmdMgr *command.CommandManager
)

func main() {
	fmt.Println("Starting...")

	// Load env :
	godotenv.Load()

	// Create client :
	client = goscord.New(&gateway.Options{
		Token:   os.Getenv("BOT_TOKEN"),
		Intents: gateway.IntentGuilds | gateway.IntentGuildMessages | gateway.IntentGuildMembers | gateway.IntentGuildVoiceStates | gateway.IntentMessageContent,
	})

	client.On(event.EventReady, func() {
		// Login message is handled in event/ready.go
	})

	// Load command manager :
	cmdMgr = command.NewCommandManager(client)

	_ = client.On(event.EventReady, myEvent.OnReady(client, cmdMgr))
	_ = client.On(event.EventInteractionCreate, cmdMgr.Handler(client))
	_ = client.On(event.EventGuildMemberAdd, myEvent.OnGuildMemberAdd(client))

	client.On(event.EventMessageCreate, func(msg *discord.Message) {
		if msg.Content == "ping" {
			_, err := client.Channel.SendMessage(msg.ChannelId, "Pong ! 🏓")
			if err != nil {
				fmt.Printf("Error sending message: %v\n", err)
			}
		}
	})

	client.Login()

	select {}
}
