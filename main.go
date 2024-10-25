package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

const telegramAPI = "https://api.telegram.org/bot%s/sendMessage"

func sendTelegramMessage(token string, chatID string, message string) error {
	apiUrl := fmt.Sprintf(telegramAPI, token)
	data := url.Values{
		"chat_id": {chatID},
		"text":    {message},
	}

	resp, err := http.PostForm(apiUrl, data)
	if err != nil {
		return err
	}
	_, _ = io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message, status: %d", resp.StatusCode)
	}
	return nil
}

func main() {
	article := `
Once upon a time, in a cozy forest full of tall trees and colorful flowers, there lived a little squirrel named Nibbles. Nibbles was a curious squirrel with a fluffy tail, and he loved to explore every nook and cranny of the forest. His best friend was a small bluebird named Blue, who loved to fly high and tell stories about the places he had seen.

One chilly autumn morning, Nibbles was busy gathering acorns to store for the winter. Blue fluttered down beside him, chirping with excitement.

“Nibbles!” Blue said, “I heard from the owls that there’s a secret treasure hidden in the forest! They say it’s something sparkly and magical, and no one has found it in years!”

Nibbles’ eyes sparkled with curiosity. “A treasure? Let’s go find it, Blue!” he said, his fluffy tail twitching with excitement.

Blue led the way, flying above the treetops while Nibbles scampered along the forest floor. They crossed over babbling brooks, through patches of tall grass, and under branches covered in bright red and yellow leaves. As they searched, they met some friends along the way.

First, they met Hoot, the wise old owl, who was napping on a low branch. Hoot opened one eye and asked, “What brings you two so far from home?”

“We’re looking for the secret treasure!” Nibbles squeaked. “Do you know where it is?”

Hoot smiled, his feathers fluffing up. “I’ve heard it lies beyond the Old Oak, where the flowers grow even in winter.”

With Hoot’s advice, Nibbles and Blue continued their journey until they reached the Old Oak tree. It was a huge tree with a trunk so wide that it would take ten squirrels to wrap around it! Beyond it, they saw a meadow, and there, blooming among the fallen leaves, were tiny white flowers. And right in the middle of the flowers was a little shiny stone, sparkling under the sunlight.

Nibbles gasped. “This must be the treasure!”

He picked up the stone, and as he held it, the stone began to glow softly. Suddenly, the entire meadow felt warmer, as if the stone was giving off a gentle, comforting light. Blue chirped in delight and flew in circles around Nibbles.

“This treasure isn’t just sparkly—it’s magical!” Blue chirped. “The forest will be warm through winter with this!”

Nibbles and Blue decided to share the stone’s warmth with all their friends. They took it to the center of the forest, where everyone could come and feel its cozy glow on chilly days. All the animals in the forest gathered around, from the tiny ants to the big, sleepy bear, feeling the warmth of the magical stone.

As winter came, the forest was still warm, and everyone was happy, all thanks to Nibbles and Blue’s adventure. And whenever Nibbles looked at his friend, he remembered the joy of exploring and the magic of sharing.

And so, every autumn, Nibbles and Blue would tell the story of their adventure, reminding everyone that sometimes, the greatest treasures are the ones that bring warmth and happiness to everyone around.`
	cmd := exec.Command("python3", "summarize.py", article)

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	summary := strings.TrimSpace(string(out))
	fmt.Println("Summary:", summary)

	botToken := os.Getenv("BOT_TOKEN")
	chatID := os.Getenv("CHAT_ID")
	err = sendTelegramMessage(botToken, chatID, strings.ToValidUTF8(summary, ""))
	if err != nil {
		fmt.Println("Error sending summary to Telegram:", err)
	} else {
		fmt.Println("Summary sent successfully!")
	}
}
