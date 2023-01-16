package godiscord

func NewEmbed(Title, Description, URL string) Embed {
	e := Embed{}
	emb := EmbedElement{Title: Title, Description: Description, URL: URL}
	e.Embeds = append(e.Embeds, emb)
	return e
}
