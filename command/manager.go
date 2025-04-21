package command

import (
	"log"

	"github.com/Goscord/goscord/goscord/discord"
	"github.com/Goscord/goscord/goscord/gateway"
)

type CommandManager struct {
	client   *gateway.Session
	commands map[string]Command
}

func NewCommandManager(client *gateway.Session) *CommandManager {
	mgr := &CommandManager{
		client:   client,
		commands: make(map[string]Command),
	}

	return mgr
}

func (mgr *CommandManager) Init() {
	// // Get all existing commands
	// commands, err := mgr.client.Application.GetCommands(mgr.client.Me().Id, "")
	// if err != nil {
	// 	log.Printf("Error getting commands: %v\n", err)
	// 	return
	// }

	// // Delete all existing commands
	// for _, cmd := range commands {
	// 	_ = mgr.client.Application.DeleteCommand(mgr.client.Me().Id, "", cmd.Id)
	// }

	// Register new commands
	mgr.Register(new(PingCommand))
	mgr.Register(new(BooruCommand))
}

func (mgr *CommandManager) Handler(client *gateway.Session) func(*discord.Interaction) {
	return func(interaction *discord.Interaction) {
		if interaction.Type != discord.InteractionTypeApplicationCommand {
			return
		}

		if interaction.Member == nil {
			return
		}

		if interaction.Member.User.Bot {
			return
		}

		cmd := mgr.Get(interaction.ApplicationCommandData().Name)

		if cmd != nil {
			_ = client.Interaction.DeferResponse(interaction.Id, interaction.Token, true)

			_ = cmd.Execute(&Context{Client: client, Interaction: interaction, CmdMgr: mgr})
		}
	}
}

func (mgr *CommandManager) Get(name string) Command {
	if cmd, ok := mgr.commands[name]; ok {
		return cmd
	}

	return nil
}

func (mgr *CommandManager) Register(cmd Command) {
	appCmd := &discord.ApplicationCommand{
		Name:        cmd.Name(),
		Type:        discord.ApplicationCommandChat,
		Description: cmd.Description(),
		Options:     cmd.Options(),
	}

	if _, err := mgr.client.Application.RegisterCommand(mgr.client.Me().Id, "", appCmd); err != nil {
		log.Printf("Error registering command %s: %v\n", cmd.Name(), err)
	}

	mgr.commands[cmd.Name()] = cmd
}
