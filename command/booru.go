package command

import (
	"github.com/Goscord/goscord/goscord/discord"

	"github.com/NeroQue/QueBot/booru"
)

type BooruCommand struct{}

func (c *BooruCommand) Name() string {
	return "booru"
}

func (c *BooruCommand) Description() string {
	return "Search for images on booru sites"
}

func (c *BooruCommand) Category() string {
	return "general"
}

func (c *BooruCommand) Options() []*discord.ApplicationCommandOption {
	return []*discord.ApplicationCommandOption{
		{
			Type:        discord.ApplicationCommandOptionString,
			Name:        "provider",
			Description: "Choose the booru provider",
			Required:    true,
			Choices: []*discord.ApplicationCommandOptionChoice{
				{Name: "Safebooru", Value: "safebooru"},
				{Name: "Danbooru", Value: "danbooru"},
				{Name: "Gelbooru", Value: "gelbooru"},
			},
		},
		{
			Type:        discord.ApplicationCommandOptionString,
			Name:        "tag",
			Description: "Tag to search for",
			Required:    true,
		},
	}
}

func (c *BooruCommand) Execute(ctx *Context) bool {
	// Get the provider option
	provider := ctx.Interaction.ApplicationCommandData().Options[0].Value.(string)

	// Get the tag option
	tag := ctx.Interaction.ApplicationCommandData().Options[1].Value.(string)

	image, err := booru.GetRandomImage(booru.Provider(provider), tag)
	if err != nil {
		ctx.Client.Interaction.CreateFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, err.Error())
		return false
	}

	ctx.Client.Interaction.CreateFollowupMessage(ctx.Client.Me().Id, ctx.Interaction.Token, image.ImageURL)

	return true
}
