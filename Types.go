package godisc

//Overwrite object
type Overwrite struct {
	ID string
	Type int
	allow string
	deny string
}

//Channel object
type Channel struct{
	ID string
	Type int
	GuildID string
	Position int
	PermissionOverwrites Overwrite
}
//PermissionFlags represents bitwise permission flags(discord api documentation)
type PermissionFlags struct {
	InstantInviteCreation string
	KickMembers string
	banMembers string
	Admin string
	ManageChannels string
	ManageGuild string
	AddReactions string
	ViewAuditLog string
	PrioritySpeaker string
	Stream string
	ViewChannel string
	SendMessages string
	SendTTSMessages string
	ManageMessages string
	EmbedLinks string
	AttachFiles string
	ReadMessageHistory string
	MentionEveryone string
	UseExternalEmojis string
	ViewGuildInsights string
	Connect string
	Speak string
	MuteMembers string
	DeafenMembers string
	MoveMembers string
	UseVAD string
	ChangeNickname string
	ManageNicknames string
	ManageRoles string
	ManageWebhooks string
	ManageEmojis string
}


//Embed object
type Embed struct {
	Title string
	Type string
	Description string
	URL string
	Timestamp string
	Color int
	Footer EmbedFooter
	Image EmbedImage
	Thumbnail EmbedThumbnail
	Video EmbedVideo
	Provider EmbedProvider
	Author EmbedAuthor
	Fields []EmbedField
}
//EmbedFooter is that little text below all the things in Embed object
type EmbedFooter struct {
	Text string
	IconURL string
	ProxyIconURL string
}
//EmbedImage is the image in embed texts
type EmbedImage struct {
	URL string
	ProxyURL string
	Height int
	Width int
}
//EmbedThumbnail is the thumbnail of the embed message
type EmbedThumbnail struct {
	URL string
	ProxyURL string
	Height int
	Width int
}
//EmbedVideo is the video attached to embed message
type EmbedVideo struct {
	URL string
	Height int
	Width int
}
//EmbedAuthor is the author of the embed message
type EmbedAuthor struct {
	name string
	URL string
	IconURL string
	ProxyIconURL string
}
//EmbedProvider is the provider of the embed message
type EmbedProvider struct {
	name string
	URL string
}
//EmbedField is literally a text that'll be in the embed message
type EmbedField struct {
	name string
	value string
	inline bool
}
//Attachment is just external things u add to message, such as pictures.
type Attachment struct {
	ID string
	Filename string
	Size int
	URL string
	ProxyURL string
	Height int
	Width int
}
//Mention is... you know, PING
type Mention struct {
	ID string
	GuildID string
	Type int
	Name string
}
//AllowedMentions object is techincally a mentions option
type AllowedMentions struct {
	Parse []string
	Roles []string
	Users []string
	RepliedUser bool
}