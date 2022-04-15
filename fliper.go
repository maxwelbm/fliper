package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

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

// This regex will find the fliper emoji in chat
var rgxFindFliper = regexp.MustCompile(`(<:fliper:(\d)*>)`)

func getRoleByName(roleName string, s *discordgo.Session, guildID string) (string, error) {
	roles, err := s.GuildRoles(guildID)

	if err != nil {
		return "", nil
	}

	var fliperRoleID string

	for _, v := range roles {
		if *&v.Name == roleName {
			fliperRoleID = *&v.ID
		}
	}

	if len(fliperRoleID) == 0 {
		return "", fmt.Errorf("role %s not found", roleName)
	}

	return fliperRoleID, nil
}

func getChannelIDByName(chanName string, s *discordgo.Session, guildID string) (string, error) {
	channels, err := s.GuildChannels(guildID)

	if err != nil {
		return "", err
	}

	var ChID string

	for _, v := range channels {
		if v.Name == chanName {
			ChID = v.ID
		}
	}

	if len(ChID) == 0 {
		return "", fmt.Errorf("channel %s not found", chanName)
	}

	return ChID, nil

}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	targetChanId, err := getChannelIDByName("test_chan", s, m.GuildID) // We'll need to change this to name of our target channel

	if err != nil {
		panic(err)
	}

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	fliperRoleID, err := getRoleByName("supremo", s, m.GuildID) // We'll need to change this to name of our role

	if err != nil {
		panic(err)
	}

	findEmoji := rgxFindFliper.FindStringSubmatch(m.Content)

	if len(findEmoji) != 0 {
		if len(findEmoji[1]) != 0 {
			s.ChannelMessageSend(targetChanId, genLines(m.Author.ID, fliperRoleID))
		}
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

func genLines(callerID, roleID string) string {

	rand.Seed(time.Now().UnixNano())

	randNum := rand.Intn(10)

	switch randNum {
	case 0:
		return "<@&" + roleID + "> " + "The user <@" + callerID + "> needs your help\n"
	case 1:
		return "Please Dungeon Master <@&" + roleID + ">" + ", thy knight Sir <@" + callerID + "> needs your help to win the battle against the TS!\n"
	case 2:
		return "<@&" + roleID + ">" + " HELP MEEE!!!" + " \n by: <@" + callerID + "> :)\n"
	case 3:
		return "Hey <@&" + roleID + ">" + " you are the guy!" + " \nPlease help <@" + callerID + "> to fly!\n"
	case 4:
		return "Hey <@&" + roleID + ">" + " !" + " \n2 minds think better than one <@" + callerID + "> ;)\n"
	case 5:
		return "<@&" + roleID + ">" + " .... " + " \nThe user <@" + callerID + "> has a question that only you know the answer\n"
	case 6:
		return "<@&" + roleID + ">" + " !" + "\nI ( <@" + callerID + "> ) choose you! \n"
	default:
		return "Please <@&" + roleID + ">" + " enter at chat" + " \nThe user <@" + callerID + "> is calling you :)\n"
	}
}
