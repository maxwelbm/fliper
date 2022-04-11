package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	Token string = ""
)

const KuteGoAPIURL = "https://kutego-api-xxxxx-ew.a.run.app"

func main() {
	println("[ WELCOME ]: i am bot Fliper!")

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

type Gopher struct {
	Name string `json: "name"`
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "healthy" {
		s.ChannelMessageSend(m.ChannelID, "i am healthy!")
	}

	if m.Content == "gopher" {
		response, err := http.Get("https://camo.githubusercontent.com/d79500e3799b96ff64ee049acfec5e001e9e8b20d4ceaf6811dace0bf0f7ecc4/68747470733a2f2f6d69726f2e6d656469756d2e636f6d2f6d61782f3338342f302a413645425f596b6b73356250705f724d2e676966")
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			_, err = s.ChannelFileSend(m.ChannelID, "golpher.png", response.Body)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: Can't get dr-who Gopher! :-(")
		}
	}

	if m.Content == "fliper" {
		response, err := http.Get("https://raw.githubusercontent.com/MaxwelMazur/fliper/main/fliper.png")
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			_, err = s.ChannelFileSend(m.ChannelID, "random-gopher.png", response.Body)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: Can't get random Gopher! :-(")
		}
	}
}
