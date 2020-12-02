package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"fmt"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("plik .env dostał depresji, albo go nie stworzyłeś, upewnij się, czy wszystko jest dobrze")
		return
	}

	dg, err := discordgo.New("Bot "+os.Getenv("TOKEN"));
	if err != nil {
		fmt.Println("sesja discordowa dostała raka raka, ", err)
		return
	}
	err = dg.Open()
	if err != nil {
		fmt.Println("połączenie dostało raka raka, ", err)
		return
	}

	dg.AddHandler(reactionAddEvent)

	user, err := dg.User("@me")

	dg.UpdateStatusComplex(discordgo.UpdateStatusData{
		Game: &discordgo.Game{
			Name: "🇭",
			Type: discordgo.GameTypeWatching,
		},
	})

	fmt.Println("Zalogowano jako "+user.Username+"#"+user.Discriminator)

	<-make(chan struct{})
	return
}

func reactionAddEvent (s *discordgo.Session, r *discordgo.MessageReactionAdd){
	if r.Emoji.Name != "🇭" {
		return
	}

	msg, err := s.ChannelMessage(r.ChannelID, r.MessageID)
	if err != nil {
		fmt.Println("lol coś się popsuło")
		return
	}

	var contains = false
	if r.UserID == msg.Author.ID {
		contains = true
	}

	if contains {
		err := s.MessageReactionRemove(r.ChannelID, r.MessageID, r.Emoji.Name, r.UserID)
		if err != nil {
			fmt.Println("lol coś się popsuło")
			return
		}
		return
	}

	// a tutaj bedziemy kasować czy coś
}