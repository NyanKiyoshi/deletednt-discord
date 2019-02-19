package deletednt

import (
	"bytes"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"strings"
)

var httpClient = getHttpClient()

func processAttachment(attachment *discordgo.MessageAttachment) {
	// Get the attachment using a HEAD request and ensure it succeed: no error and HTTP 200
	if response, err := httpClient.Head(attachment.URL); err != nil {
		log.Fatalf("failed to get %s: %s", attachment.URL, err)
	} else if response.StatusCode != 200 {
		log.Fatalf(
			"failed to get %s, got response code: %d",
			attachment.URL, response.StatusCode)
	} else {
		log.Print("successfully pre-fetched ", attachment.URL)
	}
}

func processAttachments(attachments []*discordgo.MessageAttachment) {
	for _, attachment := range attachments {
		processAttachment(attachment)
	}
}

func retrieveDeletedAttachment(attachment *discordgo.MessageAttachment) *discordgo.File {
	response, err := httpClient.Get(attachment.URL)
	if err == nil {
		defer response.Body.Close()

		if response.StatusCode == 200 {
			var buffer []byte
			buffer, err := ioutil.ReadAll(response.Body)

			if err == nil {
				reader := bytes.NewReader(buffer)
				contentType := strings.Split(response.Header.Get("Content-Type"), ";")[0]

				return &discordgo.File{
					Name:        attachment.Filename,
					ContentType: contentType,
					Reader:      reader,
				}
			}
		}
	}

	if err == nil {
		err = fmt.Errorf("received unexpected status %d", response.StatusCode)
	}
	log.Printf("failed to get %s: %s", attachment.URL, err)
	return nil
}

func processDeletedMessage(
	session *discordgo.Session, deletedMessage *discordgo.Message, outputChannel *discordgo.Channel) {

	filesToSend := make([]*discordgo.File, len(deletedMessage.Attachments))
	messageToSend := discordgo.MessageSend{
		Content: fmt.Sprintf(
			"%s's message was deleted: %s", deletedMessage.Author.String(), deletedMessage.Content),
		Embed: nil,
		Files: filesToSend,
	}

	for index, attachment := range deletedMessage.Attachments {
		filesToSend[index] = retrieveDeletedAttachment(attachment)
	}

	if _, err := session.ChannelMessageSendComplex(
		outputChannel.ID, &messageToSend); err != nil {
		log.Print("Failed to send back deleted message: ", err)
	}
}
