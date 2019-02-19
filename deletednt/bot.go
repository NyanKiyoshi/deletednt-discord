package deletednt

import "github.com/bwmarrin/discordgo"

func InitBot(session *discordgo.Session) {
	// Register the onMessageCreate func as a callback for MessageCreate events.
	session.AddHandler(onMessageCreate)

	// Register the onMessageDelete func as a callback for MessageDelete events.
	session.AddHandler(onMessageDelete)
}
