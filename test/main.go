package main

import "fmt"

func main() {
	fmt.Print("Hello, world.")
}

func SendBasicWebhook(webhook string, title string, content string) {

	embed := godiscord.NewEmbed(title, content, "https://tesla.com")
	embed.SetColor("#B2010A")
	embed.SendToWebhook(webhook)
}
