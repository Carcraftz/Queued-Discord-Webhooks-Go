package queuedwebhooksdiscord

//Embed is a struct representing a Discord embed object
type Embed struct {
	Username  string         `json:"username"`
	AvatarURL string         `json:"avatar_url"`
	Content   string         `json:"content"`
	Embeds    []EmbedElement `json:"embeds"`
}

//EmbedElement is a struct representing an Embed element of the Embed struct
type EmbedElement struct {
	Author      Author  `json:"author"`
	Title       string  `json:"title"`
	URL         string  `json:"url"`
	Description string  `json:"description"`
	Color       int64   `json:"color"`
	Fields      []Field `json:"fields"`
	Thumbnail   Image   `json:"thumbnail,omitempty"`
	Image       Image   `json:"image,omitempty"`
	Footer      Footer  `json:"footer"`
}

//Author represents the author of the embed
type Author struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url"`
}

//Field represents a field in an embed
type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

//Footer represents the footer of an embed
type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url,omitempty"`
}

//Image represents the image of an embed
type Image struct {
	URL string `json:"url"`
}

//Webhook represents a webhook
type Webhook struct {
	URL     string `json:"webhook"`
	IconURL string `json:"icon_url"`
	Text    string `json:"text"`
	Color   string `json:"color"`
}
