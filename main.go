package main

import (
	"flag"
	"os"
	"strconv"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

var (
	token, guildID string
	showStatus     bool
)

func init() {
	flag.StringVar(&token, "token", "", "Bot token (required)")
	flag.StringVar(
		&guildID,
		"guildID",
		"",
		"ID of guild to remove commands from. If no ID is given, it will remove all commands from all guilds",
	)
	flag.BoolVar(&showStatus, "showStatus", false, "Show a status message for the bot during the cleanup process (default: false)")

	flag.Parse()

	if envToken := os.Getenv("DISCORD_CC_TOKEN"); token == "" && envToken == "" {
		log.Fatal("a bot token must be supplied with the -token flag or the DISCORD_CC_TOKEN variable")
	} else if token == "" && envToken != "" {
		token = envToken
	}

	if guildID == "" {
		guildID = os.Getenv("DISCORD_CC_GUILD_ID")
	}

	if !showStatus {
		if s, err := strconv.ParseBool(os.Getenv("DISCORD_CC_SHOW_STATUS")); err == nil {
			showStatus = s
		}
	}
}

func main() {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.WithError(err).Fatal("could not initialize bot")
	}

	if showStatus {
		dg.AddHandler(func(s *discordgo.Session, e *discordgo.Ready) {
			s.UpdateGameStatus(0, "Cleaning some things up")
		})
	}

	err = dg.Open()
	if err != nil {
		log.WithError(err).Fatal("could not start bot")
	}

	commands, err := dg.ApplicationCommands(dg.State.User.ID, guildID)
	if err != nil {
		log.WithError(err).WithField("guildID", guildID).Fatal("could not get application commands")
	}

	if len(commands) == 0 {
		log.Info("no commands were found. exiting...")
		return
	}

	log.WithField("guildID", guildID).Infof("found %d commands to delete", len(commands))

	var numDeleted int
	for _, command := range commands {
		err := dg.ApplicationCommandDelete(dg.State.User.ID, guildID, command.ID)

		if err != nil {
			log.WithError(err).WithFields(log.Fields{
				"id":      command.ID,
				"name":    command.Name,
				"guildID": guildID,
			}).Warn("could not delete command")
		} else {
			numDeleted++
		}
	}

	log.Infof("deleted %d/%d commands. exiting...", numDeleted, len(commands))
	dg.Close()
}
