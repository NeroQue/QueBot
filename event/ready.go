package event

import (
	"log"

	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/gateway"

	"github.com/NeroQue/QueBot/command"
)

func OnReady(client *gateway.Session, cmdMgr *command.CommandManager) func() {
	return func() {
		log.Printf("Logged in as %s\n", client.Me().Tag())

		cmdMgr.Init()

		_ = client.SetActivity(&discord.Activity{Name: "/help", Type: discord.ActivityListening})
	}
}