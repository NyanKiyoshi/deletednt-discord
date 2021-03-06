package deletednt

import "github.com/bwmarrin/discordgo"

var messageHistory = map[string]map[string]*discordgo.Message{}

func appendMessageToHistory(session *discordgo.Session, message *discordgo.Message) {
	session.Lock()
	defer session.Unlock()

	history, found := messageHistory[message.ChannelID]

	if !found {
		history = make(map[string]*discordgo.Message, 1)
		messageHistory[message.ChannelID] = history
	}

	messageHistory[message.ChannelID][message.ID] = message
}

func getMessageFromHistory(message *discordgo.Message) *discordgo.Message {
	if history, found := messageHistory[message.ChannelID]; found {
		return history[message.ID]
	}
	return nil
}

func popMessageFromHistory(
	session *discordgo.Session, message *discordgo.Message) *discordgo.Message {
	session.Lock()
	defer session.Unlock()

	cachedMessage := getMessageFromHistory(message)
	if cachedMessage != nil {
		delete(messageHistory[message.ChannelID], message.ID)
	}

	return cachedMessage
}
