package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"flag"
	"os"
	"bufio"
)

const (
	prefix = "miku"
)

var (
	token	string
	r	*router
)

func messageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// if user is a bot, ignore his message
	if m.Author.Bot {
		return
	}

	// filter out any messages less than 4 characters otherwise we will get IndexOutOfRange
	if len(m.Content) < len(prefix) {
		return
	}

	// process any message starting with "miku"
	if m.Content[:4] != prefix {
		return
	}

	// call the corresponding method
	r.getRoute(s, m.Message)
}

func main() {
	// parse token flag
	tokenPtr := flag.String("token", "", "Discord Bot Token")
	flag.Parse()
	token = *tokenPtr

	log.Printf("----- Logging in with token : %s", token)
	discord, err := discordgo.New(token)
	if err != nil {
		log.Println("----- Error logging in -----")
		log.Println(err)
		os.Exit(0)
	}
	log.Println("----- Login successful -----")

	err = discord.Open()
	if err != nil {
		log.Println("----- Error opening Discord -----")
		log.Println(err)
		os.Exit(0)
	}
	log.Println("----- Discord session started -----")
	discord.AddHandler(messageCreateHandler)

	r = newRouter()
	go r.addRoute("thisguyisafaggot", thisguyisafaggot)
	go r.addRoute("whoisafaggot", whoisafaggot)
	go r.addRoute("join", join)
	go r.addRoute("sing", sing)
	go r.addRoute("skip", skip)
	go r.addRoute("stop", stop)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		discord.ChannelMessageSend(hupsoonheng, scanner.Text())
	}

	lock := make(chan int)
	<-lock
}
