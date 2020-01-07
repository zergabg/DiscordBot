package main

import (
	"fmt"
	"log"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
	"github.com/go-ini/ini"
)

var (
	dg 						*discordgo.Session
	GuildID 				string
)

func main() {
	var err error
	cfg, err := ini.Load("login.ini")
	dg, err = discordgo.New(cfg.Section("Login").Key("Token").String())
	dg.AddHandler(func(_ *discordgo.Session, m *discordgo.MessageCreate) {

		if m.GuildID != "226368006115033091" {
			return 			
		} else {
			log.Print("OK")
		}
	})

	if err = dg.Open(); err != nil {
		log.Print("Error Opening Connection", err)
	}


    subEv := map[string]interface{}{
        "guild_id":   GuildID,
        "typing":     false,
        "activities": true, // this is the shit
    }

    if err := dg.SendWSEvent(14, subEv); err != nil {
        log.Print("SendWSevent failed badly", err)
    }


	defer dg.Close()

	fmt.Println("The bot is running")

// Keep The Bot Running Until You Press CTRL + C
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	return
}