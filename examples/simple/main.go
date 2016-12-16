package main

import (
	"log"
	"os"
	"time"

	"github.com/yanzay/tbot"
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	// Create new telegram bot server using token
	bot, err := tbot.NewServer(token)
	if err != nil {
		log.Fatal(err)
	}

	// Use whitelist for Auth middleware, allow to interact only with user1 and user2
	whitelist := []string{"user1", "user2"}
	bot.AddMiddleware(tbot.NewAuth(whitelist))

	// Yo handler works without slash, simple text response
	bot.Handle("yo", "YO!")

	// Handle with HiHandler function
	bot.HandleFunc("/hi", HiHandler)
	// Handler can accept varialbes
	bot.HandleFunc("/say {text}", SayHandler)
	// Bot can sand stickers, photos, music
	bot.HandleFunc("/sticker", StickerHandler)
	bot.HandleFunc("/photo", PhotoHandler)

	// Set default handler if you want to process unmatched input
	bot.HandleDefault(EchoHandler)

	// Start listening for messages
	err = bot.ListenAndServe()
	log.Fatal(err)
}

func HiHandler(message tbot.Message) {
	// Handler can reply with several messages
	message.Replyf("Hello, %s!", message.From.FirstName)
	time.Sleep(1 * time.Second)
	message.Reply("What's up?")
}

func SayHandler(message tbot.Message) {
	// Message contain it's varialbes from curly brackets
	message.Reply(message.Vars["text"])
}

func EchoHandler(message tbot.Message) {
	message.Reply(message.Text)
}

func StickerHandler(message tbot.Message) {
	message.ReplySticker("sticker.png")
}

func PhotoHandler(message tbot.Message) {
	message.ReplyPhoto("photo.jpg", "it's me")
}
