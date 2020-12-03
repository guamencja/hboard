package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"time"
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

	if os.Getenv("STARSELF") == "FALSE" {
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
	}


	channels, err := s.GuildChannels(r.GuildID)
	if err != nil {
		fmt.Println("lol coś się popsuło")
		return
	}

	var hchannelID = ""

    for _, c := range channels {
        if c.Type != discordgo.ChannelTypeGuildText {
            continue
        }

		if c.Name == "h-board" {
			hchannelID = c.ID
		}
	}

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name: msg.Author.Username+"#"+msg.Author.Discriminator,
			IconURL: msg.Author.AvatarURL(""),
		},
		Title: "1 🇭",
		Color: 0x00ADD8,
		Description: msg.Content,
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name: "Link",
				Value: "[Skocz ~~z mostu~~ do wiadomości](https://discord.com/channels/"+r.GuildID+"/"+r.ChannelID+"/"+r.MessageID+")",
			},
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	_, err = s.ChannelMessageSendEmbed(hchannelID, embed)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
