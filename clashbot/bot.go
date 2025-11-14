package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Congiguration failed: %s", err)
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatalf("Decoding failed: %s", err)
	}

	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			playerTag := update.Message.Text
			url := "https://api.clashroyale.com/v1/players/%23" + playerTag

			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				log.Fatalf("Bro fuck you: %s", err)
			}

			req.Header.Set("Accept", "application/json")
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.ClashApiKey))

			fmt.Printf("TOKEN RAW: %q\n", cfg.ClashApiKey)
			fmt.Println("LEN:", len(cfg.ClashApiKey))

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Fatalf("Lol: %s", err)
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("hog rider: %s", err)
			}

			resp2, _ := http.Get("https://api.ipify.org")
			myIP, _ := io.ReadAll(resp2.Body)
			fmt.Println("My public IP:", string(myIP))

			fmt.Println(string(body))

			var player Player
			lolerr := json.Unmarshal(body, &player)
			if lolerr != nil {
				log.Fatalf("Marshal: %s", lolerr)
			}

			reply := "Имя: " + player.Name + "\nКубки: " + strconv.Itoa(player.Trophies)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
			bot.Send(msg)
		}
	}
}
