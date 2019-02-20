package deletednt

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"runtime"
	"time"
	"unsafe"
)

const kibiByte = float64(1 << 10)
const mibiByte = float64(1 << 20)

var startTime = time.Now()

func getUptime() time.Duration {
	return time.Since(startTime)
}

func InitBot(session *discordgo.Session) {
	// Register the onMessageCreate func as a callback for MessageCreate events.
	session.AddHandler(onMessageCreate)

	// Register the onMessageDelete func as a callback for MessageDelete events.
	session.AddHandler(onMessageDelete)
}

func sendBotState(session *discordgo.Session, message *discordgo.Message) {
	var memstats runtime.MemStats
	runtime.ReadMemStats(&memstats)

	_, _ = session.ChannelMessageSend(message.ChannelID, fmt.Sprintf(
		"Bot Uptime: %s\nMemory Usage: %.2f MiB\nOutput Channel: %s\nAllocated Cache: %.3f KiB",
		getUptime(),
		float64(memstats.Sys)/mibiByte,
		getOutputChannelMention(session, message.GuildID),
		float64(unsafe.Sizeof(messageHistory))/kibiByte,
	))
}
