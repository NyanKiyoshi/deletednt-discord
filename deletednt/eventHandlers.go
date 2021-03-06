package deletednt

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

const statusCommand string = "?status"
const outputChannel string = "logs"

// onMessageCreate will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func onMessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if message.Author.ID == session.State.User.ID {
		return
	}

	// If the message is the status command
	if strings.TrimSpace(message.Content) == statusCommand {
		sendBotState(session, message.Message)
	}

	// Ignore messages if there are no attachments
	if len(message.Attachments) < 1 {
		return
	}

	// Process message attachments
	processAttachments(message.Message.Attachments)

	// Cache the message
	appendMessageToHistory(session, message.Message)
}

func onMessageDelete(session *discordgo.Session, message *discordgo.MessageDelete) {
	if cachedMessage := popMessageFromHistory(session, message.Message); cachedMessage != nil && len(cachedMessage.Attachments) > 0 {

		if targetChannel := getOutputChannel(session, message.GuildID); targetChannel != nil {
			go processDeletedMessage(session, cachedMessage, targetChannel)
		}
	}
}
