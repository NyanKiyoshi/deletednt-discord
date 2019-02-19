package deletednt

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

var outputChannels = map[string]*discordgo.Channel{}

func _findOutputChannel(session *discordgo.Session, guildID string) *discordgo.Channel {
	if channels, err := session.GuildChannels(guildID); err == nil {
		for _, channel := range channels {
			if strings.ToLower(channel.Name) == outputChannel {
				return channel
			}
		}
	}
	return nil
}

// getOutputChannel finds the output channel ID of the given guild.
func getOutputChannel(session *discordgo.Session, guildID string) *discordgo.Channel {
	channel, found := outputChannels[guildID]

	if !found {
		channel = _findOutputChannel(session, guildID)

		if channel != nil {
			outputChannels[guildID] = channel
		}
	}

	return channel
}

func getOutputChannelMention(session *discordgo.Session, guildID string) string {
	if targetChannel := getOutputChannel(session, guildID); targetChannel != nil {
		return targetChannel.Mention()
	} else {
		return "none"
	}
}
