package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/diamondburned/discordgo"
	"github.com/go-ini/ini"
)

var dg *discordgo.Session
var cfg struct {
	GuildID int64
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

	dg, err = discordgo.New(discordgo.Login{
		Token: cfg.Token,
	})
	if err != nil {
		log.Fatalln("Failed to create new Discord:", err)
	}

	g, err := dg.Guild(cfg.GuildID)
	if err != nil {
		log.Fatalln("Failed to get guild:", err)
	}

	log.Println("Listening to", g.Name)

	dg.AddHandler(func(_ *discordgo.Session, m *discordgo.MessageCreate) {
		if m.GuildID != cfg.GuildID {
			return
		}

		log.Println("OK", m.Content)
	})

	if err = dg.Open(); err != nil {
		log.Fatalln("Error Opening Connection:", err)
	}

	defer dg.Close()

	dg.GatewayManager.SubscribeGuild(cfg.GuildID, true, true)

	fmt.Println("The bot is running")

	// Keep The Bot Running Until You Press CTRL + C
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	return
}
