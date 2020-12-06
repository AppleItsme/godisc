package godisc

//Overwrite object
type Overwrite struct {
	ID string `json:"id"`
	Type int `json:"type"`
	Allow string `json:"allow"`
	Deny string `json:"deny"`
}

//Channel object
type Channel struct{
	ID string `json:"id"`
	Type int `json:"type"`
	GuildID string `json:"guild_id`
	Position int `json:"position"`
	PermissionOverwrites Overwrite `json:"permission_overwrites"`
	Name string `json:"name"`
	Topic string `json:"topic"`
	Nsfw bool `json:"nsfw"`
	LastMessageID string `json:"last_message_id"`
	Bitrate int `json:"bitrate"`
	UserLimit int `json:"user_limit"`
	SlowDownTime int `json:"rate_limit_per_user"`
	Recipients []*User `json:"recipients"`
	Icon string `json:"icon"`
	OwnerID string `json:"owner_id"`
	ApplicationID string `json:"application_id"`
	ParentID string `json:"parent_id"`
	LastPinTime string `json:"last_pin_timestamp"`
}

type ChannelEdit struct {
	Name string `json:"name"`
	Type int `json:"type"`
	Position int `json:"position"`
	Topic string `json:"topic"`
	Nsfw bool `json:"nsfw"`
	SlowDownTime int `json:"rate_limit_per_user"`
	Bitrate int `json:"bitrate"`
	UserLimit int `json:"user_limit"`
	PermissionOverwrites Overwrite `json:"permission_overwrites"`
	ParentID string `json:"parent_id"`
}

type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
	UserTag string `json:"discriminator"`
	Image string `json:"avatar"`
	Bot bool `json:"bot"`
	System bool `json:"system"`
	MFA bool `json:"mfa_enabled"`
	Language string `json:"locale"`
	EmailVerified bool `json:"verified"`
	Email string `json:"email"`
	Flags int `json:"flags"`
	PremiumType int `json:"premium_type"` //0 - none; 1 - nitro classic; 2 - nitro
	PublicFlags int `json:"public_flags"` //https://discord.com/developers/docs/resources/user#user-object-user-flags
}

//PermissionFlags represents bitwise permission flags(discord api documentation)
type PermissionFlags struct {
	InstantInviteCreation string `json:"CREATE_INSTANT_INVITE"`
	KickMembers string `json:"KICK_MEMBERS"`
	banMembers string `json:"BAN_MEMBERS"`
	Admin string `json:"ADMINISTRATOR"`
	ManageChannels string `json:"MANAGE_CHANNELS"`
	ManageGuild string `json:"MANAGE_GUILD"`
	AddReactions string `json:"ADD_REACTIONS"`
	ViewAuditLog string `json:"VIEW_AUDIT_LOG"`
	PrioritySpeaker string `json:"PRIORITY_SPEAKER"`
	Stream string `json:"STREAM"`
	ViewChannel string `json:"VIEW_CHANNEL"`
	SendMessages string `json:"SEND_MESSAGES"`
	SendTTSMessages string `json:"SEND_TTS_MESSAGES"`
	ManageMessages string `json:"MANAGE_MESSAGES"`
	EmbedLinks string `json:"EMBED_LINKS"`
	AttachFiles string `json:"ATTACH_FILES"`
	ReadMessageHistory string `json:"READ_MESSAGE_HISTORY"`
	MentionEveryone string `json:"MENTION_EVERYONE"`
	UseExternalEmojis string `json:"USE_EXTERNAL_EMOJIS"`
	ViewGuildInsights string `json:"VIEW_GUILD_INSIGHTS"`
	Connect string `json:"CONNECT"`
	Speak string `json:"SPEAK"`
	MuteMembers string `json:"MUTE_MEMBETS"`
	DeafenMembers string `json:"DEAFEN_MEMBERS"`
	MoveMembers string `json:"MOVE_MEMBERS"`
	UseVAD string `json:"USE_VAD"`
	ChangeNickname string `json:"CHANGE_NICKNAME"`
	ManageNicknames string `json:"MANAGE_NICKNAMES"`
	ManageRoles string `json:"MANAGE_ROLES"`
	ManageWebhooks string `json:"MANGE_WEBHOOKS"`
	ManageEmojis string `json:"MANAGE_EMOJIS"`
}


//Embed object
type Embed struct {
	Title string `json:"title"`
	Type string `json:"type"`
	Description string `json:"description"`
	URL string `json:"url"`
	Timestamp string `json:"timestamp"`
	Color int `json:"color"`
	Footer EmbedFooter `json:"footer"`
	Image EmbedImage `json:"image"`
	Thumbnail EmbedThumbnail `json:"thumbnail"`
	Video EmbedVideo `json:"video"`
	Provider EmbedProvider `json:"provider"`
	Author EmbedAuthor `json:"author"`
	Fields []EmbedField `json:"fields"`
}
//EmbedFooter is that little text below all the things in Embed object
type EmbedFooter struct {
	Text string `json:"text"`
	IconURL string `json:"icon_url"`
	ProxyIconURL string `json:"proxy_icon_url"`
}
//EmbedImage is the image in embed texts
type EmbedImage struct {
	URL string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height int `json:"height"`
	Width int `json:"width"`
}
//EmbedThumbnail is the thumbnail of the embed message
type EmbedThumbnail struct {
	URL string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height int `json:"height"`
	Width int `json:"width"`
}
//EmbedVideo is the video attached to embed message
type EmbedVideo struct {
	URL string `json:"url"`
	Height int `json:"height"`
	Width int `json:"width"`
}
//EmbedAuthor is the author of the embed message
type EmbedAuthor struct {
	name string `json:"author"`
	URL string `json:"url"`
	IconURL string `json:"icon_url"`
	ProxyIconURL string `json:"proxy_icon_url"`
}
//EmbedProvider is the provider of the embed message
type EmbedProvider struct {
	name string `json:"name"`
	URL string `json:"url"`
}
//EmbedField is literally a text that'll be in the embed message
type EmbedField struct {
	name string `json:"name"`
	value string `json:"value"`
	inline bool `json:"inline"`
}
//Attachment is just external things you add to message, such as pictures.
type Attachment struct {
	ID string `json:"id"`
	Filename string `json:"filename"`
	Size int `json:"size"`
	URL string `json:"url"`
	ProxyURL string `json:"proxy_url"`
	Height int `json:"height"`
	Width int `json:"width"`
}
//Mention is... you know, PING
type Mention struct {
	ID string `json:"id"`
	GuildID string `json:"guild_id"`
	Type int `json:"type"`
	Name string `json:"name"`
}
//AllowedMentions object is techincally an option for mentions
type AllowedMentions struct {
	Parse []string `json:"parse"`
	Roles []string `json:"roles"`
	Users []string `json:"users"`
	RepliedUser bool `json:"replied_user"`
}

type Message struct {
	ID string `json:"id"`
	ChannelID string `json:"channel_id"`
	GuildID string `json:"guild_id"`
	Author User `json:"author"`
	Member GuildMember `json:"member"`
	Content string `json:"content"`
	Timestamp string `json:"timestamp"`
	EditTimestamp string `json:"edited_timestamp"`
	Tts bool `json:"tts"`
	MentionEveryone bool `json:"mention_everyone"`
	MentionUsrs []User `json:"mentions"`
	MentionRoles []Role `json:"mentions"`
	MentionChannels []ChannelMention `json:"mention_channels"`
	Attachments []Attachment `json:"attachments"`
	Embeds []Embed `json:"embeds"`
	Reactions []Reaction `json:"reactions"`
	Nonce string `json:"nonce"`
	Pinned bool `json:"pinned"`
	WebhookID string `json:"webhook_id"`
	Type int `json:"type"`
	Activity MessageActivity `json:"activity"`
	Application MessageApplication `json:"application"`
	Reference MessageReference `json:"message_reference"`
	Flags int `json:"flags"`
	Stickers []MessageSticker `json:"stickers"`
	ReferencedMessage MessageReference `json:"referenced_message"`
}

type GuildMember struct {
	Usr User `json:"user"`
	Nickname string `json:"nick"`
	Roles []string `json:"roles"`
	JoinedAt string `json:"joined_at"`
	PremiumSince string `json:"premium_since"`
	Deaf bool `json:"deaf"`
	Mute bool `json:"mute"`
}

type ChannelMention struct {
	ID string `json:"id"`
	GuildID string `json:"guild_id"`
	Type int `json:"type"`
	Name string `json:"name"`
}

type Role struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Color int `json:"color"`
	Highlighted bool `json:"hoist"`
	Position int `json:"position"`
	Permissions string `json:"permissions"`
	Managed bool `json:"managed"`
	Mentionable bool `json:"mentionable"`
}

type Reaction struct {
	Count int `json:"count"`
	Me bool `json:"me"`
	Emoji Emoji `json:"emoji"`
}
type MessageActivity struct {
	Type int `json:"type"`
	PartyID string `json:"party_id"`
}

type MessageApplication struct {
	ID string `json:"id"`
	CoverImage string `json:"cover_image"`
	Description string `json:"description"`
	Icon string `json:"icon"`
	name string `json:"name"`
}
type MessageReference struct {
	MessageID string `json:"message_id"`
	ChannelID string `json:"channel_id"`
	GuildID string `json:"guild_id"`
}

type MessageSticker struct {
	ID string `json:"id"`
	PackID string `json:"pack_id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Asset string `json:"asset"`
	PreviewAsset string `json:"asset"`
	FormatType int `json:"format_type"`
}

type Emoji struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Roles []string `json:"roles"`
	User User `json:"user"`
	RequireColons bool `json:"require_colons"`
	Managed bool `json:"managed"`
	Animated bool `json:"animated"`
	Available bool `json:"available"`
}

type FollowedChannel struct {
	ChannelID string `json:"channel_id"`
	WebhookID string `json:"webhook_id"`
}

type Guild struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	IconHash string `json:"icon_hash"`
	Splash string `json:"splash"`
	DiscoverySplash string `json:"discovery_splash"`
	Owner bool `json:"owner"`
	OwnerID string `json:"owner_id"`
	Permissions string `json:"permissions"`
	Region string `json:"region"`
	AfkChannelID string `json:"afk_channel_id"`
	AfkTimeout int `json:"afk_timeout"`
	WidgetEnabled bool `json:"widget_enabled"`
	WidgetChannelID string `json:"widget_channel_id"`
	VerificationLvl int `json:"verification_level"`
	DefaultMessageNotifications int `json:"default_message_notifications"`
	ExplicitContentFilter int `json:"explicit_content_filter"`
	Roles []Role `json:"roles"`
	Emojis []Emoji `json:"emojis"`
	Features []string `json:"features"`
	MfaLvl int `json:"mfa_level"`
	ApplicationID string `json:"application_id"`
	SystemChannelID string `json:"system_channel_id"`
	RulesChannelID string `json:"rules_channel_id"`
	JoinedAt string `json:"joined_at"`
	Large bool `json:"large"`
	Unavailable bool `json:"unavailable"`
	MemberCount int `json:"member_count"`
	VoiceStates []VoiceState `json:"voice_states"`
	Members []GuildMember `json:"members"`
	Channels []Channel `json:"channels"`
	Presences []PresenceUpdate `json:"presences"`
	MaxPresences int `json:"max_presences"`
	MaxMembers int `json:"max_members"`
	VanityURLCode string `json:"vanity_url_code"`
	Descriprion string `json:"description"`
	Banner string `json:"banner"`
	PremiumTier int `json:"premium_tier"`
	PremiumBoostCount int `json:"premium_subscription_count"`
	PreferredLocale string `json:"preferred_locale"`
	PublicUpdatesChannelID string `json:"public_updates_channel_id"`
	MaxVideoChannelUsers int `json:"max_video_channel_users"`
	ApproximateMemberCount int `json:"approximate_member_count"`
	ApproximatePresenceCount int `json:"approximate_presence_count"`
}
//VoiceState is used to represent a user's voice connection status.
type VoiceState struct {
	GuildID string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
	UserID string `json:"user_id"`
	Member GuildMember `json:"member"`
	SessionID string `json:"sessionID"`
	Deaf bool `json:"deaf"`
	Mute bool `json:"mute"`
	SelfDeaf bool `json:"self_deaf"`
	SelfMute bool `json:"self_mute"`
	SelfStream bool `json:"self_stream"`
	CameraOn bool `json:"self_video"`
	Suppress bool `json:"suppress"`
}

type VoiceRegion struct {
	ID string `json:"id"`
	Name string `json:"name"`
	VIP bool `json:"vip"`
	Optimal bool `json:"optimal"`
	Deprecated bool `json:"deprecated"`
	Custom bool `json:"custom"`
}

type PresenceUpdate struct {
	User User `json:"user"`
	GuildID string `json:"guild_id"`
	Status string `json:"status"`
	Activities []Activity `json:"activities"`
	ClientStatus ClientStatus `json:"client_status"`
}

type ClientStatus struct {
	Desktop string `json:"desktop"`
	Mobile string `json:"mobile"`
	Web string `json:"web"`
}

type Activity struct {
	Name string `json:"name"`
	Type int `json:"type"`
	URL string `json:"url"`
	CreatedAt int `json:"created_at"`
	Timestamps ActivityTimestamps `json:"timestamps"`
	ApplicationID string `json:"application_id"`
	Details string `json:"details"`
	State string `json:"state"`
	Emoji ActivityEmoji `json:"emoji"`
	Party ActivityParty `json:"party"`
	Assets ActivityAssets `json:"assets"`
	Secrets ActivitySecrets `json:"secrets"`
	Instance bool `json:"instance"`
	Flags int `json:"flags"`
}

type ActivityTimestamps struct {
	Start int `json:"start"`
	End int `json:"end"`
}

type ActivityParty struct {
	ID string `json:"id"`
	Size [2]int `json:"size"`
}

type ActivityEmoji struct {
	Name string `json:"name"`
	ID string `json:"id"`
	Animated bool `json:"animated"`
}

type ActivityAssets struct {
	LargeImage string `json:"large_image"`
	LargeText string `json:"large_text"`
	SmallImage string `json:"small_image"`
	SmallText string `json:"small_text"`
}

type ActivitySecrets struct {
	Join string `json:"join"`
	Spectate string `json:"spectate"`
	Match string `json:"match"`
}

type GuildPreview struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Splash string `json:"splash"`
	DiscoverySplash string `json:"discovery_splash"`
	Emojis []Emoji `json:"emojis"`
	Features []string `json:"features"`
	ApproxMemberCount int `json:"approximate_member_count"`
	ApproxPresenceCount int `json:"approximate_presence_count"`
	Description string `json:"description"`
}

type GuildWidget struct {
	Enabled bool `json:"enabled"`
	ChannelID string `json:"channel_id"`
}


type Integration struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Enabled bool `json:"enabled"`
	Syncing bool `json:"syncing"`
	RoleID string `json:"role_id"`
	EnableEmoticons bool `json:"enable_emoticons"`
	ExpireBehavior int `json:"expire_behavior"`
	ExpireGracePeriod int `json:"expire_grace_periond"`
	User User `json:"user"`
	Account IntegrationAccount `json:"account"`
	SyncedAt string `json:"synced_at"`
	SubscriberCount int `json:"subscriber_count"`
	Revoked bool `json:"revoked"`
	Application IntegrationApplication `json:"application"`
}

type IntegrationAccount struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

type IntegrationApplication struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Description string `json:"description"`
	Summary string `json:"summary"`
	Bot User `json:"bot"`
}

type Ban struct {
	Reason string `json:"reason"`
	User User `json:"user"`
}

type Invite struct {
	Code string `json:"code"`
	Guild Guild `json:"guild"`
	Channel Channel `json:"channel"`
	Inviter User `json:"inviter"`
	TargetUser User `json:"target_user"`
	TargetUserType int `json:"target_user_type"`
	ApproximatePresenceCount int `json:"approximate_presence_count"`
	ApproximateMemberCount int `json:"approximate_member_count"`
}

type InviteMetadata struct {
	Uses int `json:"uses"`
	MaxUses int `json:"max_uses"`
	MaxAge int `json:"max_age"`
	Temporary bool `json:"temporary"`
	CreatedAt string `json:"created_at"`
}

type Template struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Description string `json:"description"`
	UsageCount int `json:"usage_count"`
	CreatorID string `json:"creator_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	SourceGuildID string `json:"source_guild_id"`
	SerializedSourceGuild Guild `json:"serialized_source_guild"`
	IsDirty bool `json:"is_dirty"`
}

type Connection struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Revoked bool `json:"revoked"`
	Integrations []Integration `json:"integrations"`
	Verified bool `json:"verified"`
	FriendSync bool `json:"friend_sync"`
	ShowActivity bool `json:"show_activity"`
	Visibility int `json:"visibility"`
}

type Webhook struct {
	ID string `json:"id"`
	Type int `json:"type"`
	GuildID string `json:"guild_id"`
	ChannelID string `json:"channe_id"`
	User User `json:"user"`
	Name string `json:"name"`
	Avatar string `json:"avatar"`
	Token string `json:"token"`
	ApplicationID string `json:"application_id"`
}