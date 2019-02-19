package main

import (
	"deletednt-discord/deletednt"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token    string
	Email    string
	Password string
)

func init() {

	flag.StringVar(&Token, "token", "", "Bot Token")
	flag.StringVar(&Email, "email", "", "Account Email")
	flag.StringVar(&Password, "password", "", "Account Password")

	flag.Parse()

	if Token == "" && (Email == "" || Password == "") {
		flag.Usage()
		os.Exit(1)
	}
}

func createBot() *discordgo.Session {
	var (
		session *discordgo.Session
		err     error
	)

	// Create a new Discord session using the provided bot token.
	if Token != "" {
		session, err = discordgo.New("Bot " + Token)
	} else {
		session, err = discordgo.New(Email, Password)
	}

	if err != nil {
		log.Panic("error creating Discord session: ", err.Error())
	}

	return session
}

func runForever() {
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func main() {
	session := createBot()
	deletednt.InitBot(session)

	// Open a websocket connection to Discord and begin listening.
	if err := session.Open(); err != nil {
		log.Panic("error opening connection: ", err.Error())
	}

	fmt.Printf(
		"Invite: https://discordapp.com/oauth2/authorize?client_id=%s&scope=bot\n", session.State.User.ID)

	runForever()

	// Cleanly close down the Discord session.
	_ = session.Close()
}
