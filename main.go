package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/go-ini/ini"
)

var dg *discordgo.Session
var cfg struct {
	GuildID string
	Token   string
}

func main() {
	f, err := ini.Load("login.ini")
	if err != nil {
		log.Fatalln("Failed to load login.ini:", err)
	}

	if err := f.MapTo(&cfg); err != nil {
		log.Fatalln("Failed to parse login.ini:", err)
	}

	dg, err = discordgo.New(cfg.Token)
	if err != nil {
		log.Fatalln("Failed to create new Discord:", err)
	}

	dg.AddHandler(func(_ *discordgo.Session, m *discordgo.MessageCreate) {
		if m.GuildID != cfg.Token {
			return
		}

		log.Print("OK")
	})

	if err = dg.Open(); err != nil {
		log.Fatalln("Error Opening Connection:", err)
	}

	defer dg.Close()

	if err := dg.SubscribeGuild(discordgo.SubscribeGuildData{
		GuildID:    cfg.GuildID,
		Typing:     false,
		Activities: true,
	}); err != nil {
		log.Fatalln("Failed to subscribe to guild ID:", err)
	}

	fmt.Println("The bot is running")

	// Keep The Bot Running Until You Press CTRL + C
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	return
}
