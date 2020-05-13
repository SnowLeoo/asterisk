package commands

import (
	"github.com/Lukaesebrot/asterisk/static"
	"github.com/Lukaesebrot/asterisk/utils"
	"github.com/Lukaesebrot/dgc"
)

// Blacklist handles the blacklist command
func Blacklist(ctx *dgc.Ctx) {
	// Check if the executor is a bot admin
	if !utils.IsBotAdmin(ctx.Event.Author.ID) {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateInsufficientPermissionsEmbed("You need to be a bot admin to use this command."))
		return
	}

	// Validate the arguments
	userID := ctx.Arguments.Get(0).AsUserMentionID()
	if userID == "" {
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateInvalidUsageEmbed(ctx.Command.Usage))
		return
	}

	// Add/Remove the user to/from the backlist
	if static.Blacklist[userID] {
		delete(static.Blacklist, userID)
		ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateSuccessEmbed("The user has been removed from the command blacklist."))
		return
	}
	static.Blacklist[userID] = true
	ctx.Session.ChannelMessageSendEmbed(ctx.Event.ChannelID, utils.GenerateSuccessEmbed("The user has been added to the command blacklist."))
	return
}
