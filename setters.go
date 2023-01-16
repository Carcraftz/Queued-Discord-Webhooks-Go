package godiscord

import (
	"errors"
	"strconv"
	"strings"
)

//SetAuthor sets the author of the Embed
func (e *Embed) SetAuthor(Name, URL, IconURL string) {
	if len(e.Embeds) == 0 {
		emb := EmbedElement{Author: Author{Name, URL, IconURL}}
		e.Embeds = append(e.Embeds, emb)
	} else {
		e.Embeds[0].Author = Author{Name, URL, IconURL}
	}
}

//SetColor takes in a hex code and sets the color of the Embed.
//Returns an error if the hex is invalid
func (e *Embed) SetColor(color string) error {
	color = strings.Replace(color, "0x", "", -1)
	color = strings.Replace(color, "0X", "", -1)
	color = strings.Replace(color, "#", "", -1)
	colorInt, err := strconv.ParseInt(color, 16, 64)
	if err != nil {
		return errors.New("Invalid hex code passed")
	}
	e.Embeds[0].Color = colorInt
	return nil
}

//SetThumbnail sets the thumbnail of the embed.
//Returns an error if the embed was not initialized properly
func (e *Embed) SetThumbnail(URL string) error {
	if len(e.Embeds) < 1 {
		return errors.New("Invalid Embed passed in, Embed.Embeds must have at least one EmbedElement")
	}
	e.Embeds[0].Thumbnail = Image{URL}
	return nil
}

//SetImage sets the image of the embed
//Returns an error if the embed was not initialized properly
func (e *Embed) SetImage(URL string) error {
	if len(e.Embeds) < 1 {
		return errors.New("Invalid Embed passed in, Embed.Embeds must have at least one EmbedElement")
	}
	e.Embeds[0].Image = Image{URL}
	return nil
}

//SetFooter sets the footer of the embed.
//Returns an error if the embed was not initialized properly
func (e *Embed) SetFooter(Text, IconURL string) error {
	if len(e.Embeds) < 1 {
		return errors.New("Invalid Embed passed in, Embed.Embeds must have at least one EmbedElement")
	}
	e.Embeds[0].Footer = Footer{Text, IconURL}
	return nil
}

//AddField adds a frield to the Embed.
//Returns an error if the embed was not initialized properly
func (e *Embed) AddField(Name, Value string, Inline bool) error {
	if len(e.Embeds) < 1 {
		return errors.New("Invalid Embed passed in, Embed.Embeds must have at least one EmbedElement")
	}
	e.Embeds[0].Fields = append(e.Embeds[0].Fields, Field{Name, Value, Inline})
	return nil
}

//SetUsername sets the username of the Embed sender
func (e *Embed) SetUsername(Name string) {
	e.Username = Name
}

//SetAvatarURL sets the Avatar URL of the Embed sender
func (e *Embed) SetAvatarURL(avatarURL string) {
	e.AvatarURL = avatarURL
}

//SetContent gives the message a Content value
func (e *Embed) SetContent(content string) {
	e.Content = content
}
