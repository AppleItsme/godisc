package godisc

//Overwrite is permissions object
type Overwrite struct {
	ID string `json:"id,omitempty"`	//either role ID or user ID
	Type string `json:"type,omitempty"` //0 for a role and 1 for a user
	Allow int `json:"allow,omitempty"` //What permissions to allow
	Deny int `json:"deny,omitempty"` //What permission to disallow
}

// Constants for the different bit offsets of text channel permissions
const (
	PermissionViewChannels = 1 << (iota + 10)
	PermissionSendMessages
	PermissionSendTTSMessages
	PermissionManageMessages
	PermissionEmbedLinks
	PermissionAttachFiles
	PermissionReadMessageHistory
	PermissionMentionEveryone
	PermissionUseExternalEmojis
)

// Constants for the different bit offsets of voice permissions
const (
	PermissionVoiceConnect = 1 << (iota + 20)
	PermissionVoiceSpeak
	PermissionVoiceMuteMembers
	PermissionVoiceDeafenMembers
	PermissionVoiceMoveMembers
	PermissionVoiceUseVAD
	PermissionVoicePrioritySpeaker = 1 << (iota + 2)
)

// Constants for general management.
const (
	PermissionChangeNickname = 1 << (iota + 26)
	PermissionManageNicknames
	PermissionManageRoles
	PermissionManageWebhooks
	PermissionManageEmojis
)

// Constants for the different bit offsets of general permissions
const (
	PermissionCreateInstantInvite = 1 << iota
	PermissionKickMembers
	PermissionBanMembers
	PermissionAdministrator
	PermissionManageChannels
	PermissionManageServer
	PermissionAddReactions
	PermissionViewAuditLogs
	PermissionViewChannel = 1 << (iota + 2)

	PermissionAllText = PermissionViewChannel |
		PermissionSendMessages |
		PermissionSendTTSMessages |
		PermissionManageMessages |
		PermissionEmbedLinks |
		PermissionAttachFiles |
		PermissionReadMessageHistory |
		PermissionMentionEveryone
	PermissionAllVoice = PermissionViewChannel |
		PermissionVoiceConnect |
		PermissionVoiceSpeak |
		PermissionVoiceMuteMembers |
		PermissionVoiceDeafenMembers |
		PermissionVoiceMoveMembers |
		PermissionVoiceUseVAD |
		PermissionVoicePrioritySpeaker
	PermissionAllChannel = PermissionAllText |
		PermissionAllVoice |
		PermissionCreateInstantInvite |
		PermissionManageRoles |
		PermissionManageChannels |
		PermissionAddReactions |
		PermissionViewAuditLogs
	PermissionAll = PermissionAllChannel |
		PermissionKickMembers |
		PermissionBanMembers |
		PermissionManageServer |
		PermissionAdministrator |
		PermissionManageWebhooks |
		PermissionManageEmojis
)





const (
	TextChannel int = iota
	DM int = iota
	GuildVoice int = iota
	GroupDM int = iota
	Category int = iota
	News int = iota
	Store int = iota
)

type ChannelEdit struct {
	Name                 string                 `json:"name,omitempty"`
	Topic                string                 `json:"topic,omitempty"`
	Nsfw                 bool                   `json:"nsfw,omitempty"`
	Position             int                    `json:"position,omitempty"`
	Bitrate              int                    `json:"bitrate,omitempty"`
	UserLimit            int                    `json:"user_limit,omitempty"`
	PermissionOverwrites []Overwrite `json:"permission_overwrites,omitempty"`
	ParentID             string                 `json:"parent_id,omitempty"`
	SlowDownTime     int                    `json:"rate_limit_per_user,omitempty"`
}

type GetMessagesParams struct {
	Around string `json:"around,omitempty"`
	Before string `json:"before,omitempty"`
	After string `jso:"after,omitempty"`
	Limit int `json:"limit,omitempty"`
}

//Channel object
type Channel struct{
	ID string `json:"id,omitempty"`
	Type int `json:"type,omitempty"`
	GuildID string `json:"guild_id`
	Position int `json:"position,omitempty"`
	PermissionOverwrites []Overwrite `json:"permission_overwrites,omitempty"`
	Name string `json:"name,omitempty"`
	Topic string `json:"topic,omitempty"`
	Nsfw bool `json:"nsfw,omitempty"`
	LastMessageID string `json:"last_message_id,omitempty"`
	Bitrate int `json:"bitrate,omitempty"`
	UserLimit int `json:"user_limit,omitempty"`
	SlowDownTime int `json:"rate_limit_per_user,omitempty"`
	Recipients []*User `json:"recipients,omitempty"`
	Icon string `json:"icon,omitempty"`
	OwnerID string `json:"owner_id,omitempty"`
	ApplicationID string `json:"application_id,omitempty"`
	ParentID string `json:"parent_id,omitempty"`
	LastPinTime string `json:"last_pin_timestamp,omitempty"`
}


type User struct {
	ID string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	UserTag string `json:"discriminator,omitempty"`
	Image string `json:"avatar,omitempty"`
	Bot bool `json:"bot,omitempty"`
	System bool `json:"system,omitempty"`
	MFA bool `json:"mfa_enabled,omitempty"`
	Language string `json:"locale,omitempty"`
	EmailVerified bool `json:"verified,omitempty"`
	Email string `json:"email,omitempty"`
	Flags int `json:"flags,omitempty"`
	PremiumType int `json:"premium_type,omitempty"` //0 - none; 1 - nitro classic; 2 - nitro
	PublicFlags int `json:"public_flags,omitempty"` //https://discord.com/developers/docs/resources/user#user-object-user-flags
}

//PermissionFlags represents bitwise permission flags(discord api documentation)
type PermissionFlags struct {
	InstantInviteCreation string `json:"CREATE_INSTANT_INVITE,omitempty"`
	KickMembers string `json:"KICK_MEMBERS,omitempty"`
	banMembers string `json:"BAN_MEMBERS,omitempty"`
	Admin string `json:"ADMINISTRATOR,omitempty"`
	ManageChannels string `json:"MANAGE_CHANNELS,omitempty"`
	ManageGuild string `json:"MANAGE_GUILD,omitempty"`
	AddReactions string `json:"ADD_REACTIONS,omitempty"`
	ViewAuditLog string `json:"VIEW_AUDIT_LOG,omitempty"`
	PrioritySpeaker string `json:"PRIORITY_SPEAKER,omitempty"`
	Stream string `json:"STREAM,omitempty"`
	ViewChannel string `json:"VIEW_CHANNEL,omitempty"`
	SendMessages string `json:"SEND_MESSAGES,omitempty"`
	SendTTSMessages string `json:"SEND_TTS_MESSAGES,omitempty"`
	ManageMessages string `json:"MANAGE_MESSAGES,omitempty"`
	EmbedLinks string `json:"EMBED_LINKS,omitempty"`
	AttachFiles string `json:"ATTACH_FILES,omitempty"`
	ReadMessageHistory string `json:"READ_MESSAGE_HISTORY,omitempty"`
	MentionEveryone string `json:"MENTION_EVERYONE,omitempty"`
	UseExternalEmojis string `json:"USE_EXTERNAL_EMOJIS,omitempty"`
	ViewGuildInsights string `json:"VIEW_GUILD_INSIGHTS,omitempty"`
	Connect string `json:"CONNECT,omitempty"`
	Speak string `json:"SPEAK,omitempty"`
	MuteMembers string `json:"MUTE_MEMBETS,omitempty"`
	DeafenMembers string `json:"DEAFEN_MEMBERS,omitempty"`
	MoveMembers string `json:"MOVE_MEMBERS,omitempty"`
	UseVAD string `json:"USE_VAD,omitempty"`
	ChangeNickname string `json:"CHANGE_NICKNAME,omitempty"`
	ManageNicknames string `json:"MANAGE_NICKNAMES,omitempty"`
	ManageRoles string `json:"MANAGE_ROLES,omitempty"`
	ManageWebhooks string `json:"MANGE_WEBHOOKS,omitempty"`
	ManageEmojis string `json:"MANAGE_EMOJIS,omitempty"`
}


//Embed object
type Embed struct {
	Title string `json:"title,omitempty"`
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	URL string `json:"url,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Color int `json:"color,omitempty"`
	Footer EmbedFooter `json:"footer,omitempty"`
	Image EmbedImage `json:"image,omitempty"`
	Thumbnail EmbedThumbnail `json:"thumbnail,omitempty"`
	Video EmbedVideo `json:"video,omitempty"`
	Provider EmbedProvider `json:"provider,omitempty"`
	Author EmbedAuthor `json:"author,omitempty"`
	Fields []EmbedField `json:"fields,omitempty"`
}

const (
	Default int = 0
	Aqua  int = 1752220
	Green  int = 3066993
	Blue  int = 3447003
	Purple int = 10181046
	Gold  int = 15844367
	Orange  int = 15105570
	Grey  int = 9807270
	Red  int = 15158332
	Navy  int = 3426654
	DarkerGrey  int = 8359053
	DarkGreen  int = 2067276
	DarkBlue  int = 2123412
	DarkPurple int = 7419530
	DarkGold  int = 12745742
	DarkOrange  int = 11027200
	DarkRed  int = 10038562
	DarkGrey  int = 9936031
	LIGHTGrey  int = 12370112
	DarkNavy  int = 2899536
	LuminousVividPink  int = 16580705
	DarkVividPink  int = 12320855
	DarkAqua  int = 1146986
)
//EmbedFooter is that little text below all the things in Embed object
type EmbedFooter struct {
	Text string `json:"text,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}
//EmbedImage is the image in embed texts
type EmbedImage struct {
	URL string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height int `json:"height,omitempty"`
	Width int `json:"width,omitempty"`
}
//EmbedThumbnail is the thumbnail of the embed message
type EmbedThumbnail struct {
	URL string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height int `json:"height,omitempty"`
	Width int `json:"width,omitempty"`
}
//EmbedVideo is the video attached to embed message
type EmbedVideo struct {
	URL string `json:"url,omitempty"`
	Height int `json:"height,omitempty"`
	Width int `json:"width,omitempty"`
}
//EmbedAuthor is the author of the embed message
type EmbedAuthor struct {
	name string `json:"author,omitempty"`
	URL string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}
//EmbedProvider is the provider of the embed message
type EmbedProvider struct {
	name string `json:"name,omitempty"`
	URL string `json:"url,omitempty"`
}
//EmbedField is literally a text that'll be in the embed message
type EmbedField struct {
	name string `json:"name,omitempty"`
	value string `json:"value,omitempty"`
	inline bool `json:"inline,omitempty"`
}
//Attachment is just external things you add to message, such as pictures.
type Attachment struct {
	ID string `json:"id,omitempty"`
	Filename string `json:"filename,omitempty"`
	Size int `json:"size,omitempty"`
	URL string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height int `json:"height,omitempty"`
	Width int `json:"width,omitempty"`
}
//Mention is... you know, PING
type Mention struct {
	ID string `json:"id,omitempty"`
	GuildID string `json:"guild_id,omitempty"`
	Type int `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
}
//AllowedMentions object is techincally an option for mentions
type AllowedMentions struct {
	Parse []string `json:"parse,omitempty"`
	Roles []string `json:"roles,omitempty"`
	Users []string `json:"users,omitempty"`
	RepliedUser bool `json:"replied_user,omitempty"`
}

type Message struct {
	ID string `json:"id,omitempty"`
	ChannelID string `json:"channel_id,omitempty"`
	GuildID string `json:"guild_id,omitempty"`
	Author User `json:"author,omitempty"`
	Member GuildMember `json:"member,omitempty"`
	Content string `json:"content,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	EditTimestamp string `json:"edited_timestamp,omitempty"`
	Tts bool `json:"tts,omitempty"`
	MentionEveryone bool `json:"mention_everyone,omitempty"`
	MentionUsrs []User `json:"mentions,omitempty"`
	MentionRoles []Role `json:"mentions,omitempty"`
	MentionChannels []ChannelMention `json:"mention_channels,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	Embeds []Embed `json:"embeds,omitempty"`
	Reactions []Reaction `json:"reactions,omitempty"`
	Nonce string `json:"nonce,omitempty"`
	Pinned bool `json:"pinned,omitempty"`
	WebhookID string `json:"webhook_id,omitempty"`
	Type int `json:"type,omitempty"`
	Activity MessageActivity `json:"activity,omitempty"`
	Application MessageApplication `json:"application,omitempty"`
	Reference MessageReference `json:"message_reference,omitempty"`
	Flags int `json:"flags,omitempty"`
	Stickers []MessageSticker `json:"stickers,omitempty"`
	ReferencedMessage MessageReference `json:"referenced_message,omitempty"`
}

type MessageEdit struct {
	Text string `json:"content,omitempty"`
	Embed *Embed `json:"embed,omitempty"`
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"` 
}

type MessageSend struct {
	Text string `json:"content,omitempty"`
	Nonce string  `json:"nonce,omitempty"`
	TTS bool  `json:"tts,omitempty"`
	Embed *Embed  `json:"embed,omitempty"`
	MentionsAllowed *AllowedMentions `json:"allowed_mentions,omitempty"` 
	MessageReply *MessageReference `json:"message_reference,omitzero,omitempty"`
}

type GuildMember struct {
	Usr User `json:"user,omitempty"`
	Nickname string `json:"nick,omitempty"`
	Roles []string `json:"roles,omitempty"`
	JoinedAt string `json:"joined_at,omitempty"`
	PremiumSince string `json:"premium_since,omitempty"`
	Deaf bool `json:"deaf,omitempty"`
	Mute bool `json:"mute,omitempty"`
}

type ChannelMention struct {
	ID string `json:"id,omitempty"`
	GuildID string `json:"guild_id,omitempty"`
	Type int `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
}

type Role struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Color int `json:"color,omitempty"`
	Highlighted bool `json:"hoist,omitempty"`
	Position int `json:"position,omitempty"`
	Permissions string `json:"permissions,omitempty"`
	Managed bool `json:"managed,omitempty"`
	Mentionable bool `json:"mentionable,omitempty"`
}

type Reaction struct {
	Count int `json:"count,omitempty"`
	Me bool `json:"me,omitempty"`
	Emoji Emoji `json:"emoji,omitempty"`
}
type MessageActivity struct {
	Type int `json:"type,omitempty"`
	PartyID string `json:"party_id,omitempty"`
}

type MessageApplication struct {
	ID string `json:"id,omitempty"`
	CoverImage string `json:"cover_image,omitempty"`
	Description string `json:"description,omitempty"`
	Icon string `json:"icon,omitempty"`
	Name string `json:"name,omitempty"`
}
type MessageReference struct {
	MessageID string `json:"message_id,omitempty"`
	ChannelID string `json:"channel_id,omitempty"`
	GuildID string `json:"guild_id,omitempty"`
}

type MessageSticker struct {
	ID string `json:"id,omitempty"`
	PackID string `json:"pack_id,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Asset string `json:"asset,omitempty"`
	PreviewAsset string `json:"asset,omitempty"`
	FormatType int `json:"format_type,omitempty"`
}

type Emoji struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Roles []string `json:"roles,omitempty"`
	User User `json:"user,omitempty"`
	RequireColons bool `json:"require_colons,omitempty"`
	Managed bool `json:"managed,omitempty"`
	Animated bool `json:"animated,omitempty"`
	Available bool `json:"available,omitempty"`
}

type FollowedChannel struct {
	ChannelID string `json:"channel_id,omitempty"`
	WebhookID string `json:"webhook_id,omitempty"`
}

type Guild struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Icon string `json:"icon,omitempty"`
	IconHash string `json:"icon_hash,omitempty"`
	Splash string `json:"splash,omitempty"`
	DiscoverySplash string `json:"discovery_splash,omitempty"`
	Owner bool `json:"owner,omitempty"`
	OwnerID string `json:"owner_id,omitempty"`
	Permissions string `json:"permissions,omitempty"`
	Region string `json:"region,omitempty"`
	AfkChannelID string `json:"afk_channel_id,omitempty"`
	AfkTimeout int `json:"afk_timeout,omitempty"`
	WidgetEnabled bool `json:"widget_enabled,omitempty"`
	WidgetChannelID string `json:"widget_channel_id,omitempty"`
	VerificationLvl int `json:"verification_level,omitempty"`
	DefaultMessageNotifications int `json:"default_message_notifications,omitempty"`
	ExplicitContentFilter int `json:"explicit_content_filter,omitempty"`
	Roles []Role `json:"roles,omitempty"`
	Emojis []Emoji `json:"emojis,omitempty"`
	Features []string `json:"features,omitempty"`
	MfaLvl int `json:"mfa_level,omitempty"`
	ApplicationID string `json:"application_id,omitempty"`
	SystemChannelID string `json:"system_channel_id,omitempty"`
	RulesChannelID string `json:"rules_channel_id,omitempty"`
	JoinedAt string `json:"joined_at,omitempty"`
	Large bool `json:"large,omitempty"`
	Unavailable bool `json:"unavailable,omitempty"`
	MemberCount int `json:"member_count,omitempty"`
	VoiceStates []VoiceState `json:"voice_states,omitempty"`
	Members []GuildMember `json:"members,omitempty"`
	Channels []Channel `json:"channels,omitempty"`
	Presences []PresenceUpdate `json:"presences,omitempty"`
	MaxPresences int `json:"max_presences,omitempty"`
	MaxMembers int `json:"max_members,omitempty"`
	VanityURLCode string `json:"vanity_url_code,omitempty"`
	Descriprion string `json:"description,omitempty"`
	Banner string `json:"banner,omitempty"`
	PremiumTier int `json:"premium_tier,omitempty"`
	PremiumBoostCount int `json:"premium_subscription_count,omitempty"`
	PreferredLocale string `json:"preferred_locale,omitempty"`
	PublicUpdatesChannelID string `json:"public_updates_channel_id,omitempty"`
	MaxVideoChannelUsers int `json:"max_video_channel_users,omitempty"`
	ApproximateMemberCount int `json:"approximate_member_count,omitempty"`
	ApproximatePresenceCount int `json:"approximate_presence_count,omitempty"`
}
//VoiceState is used to represent a user's voice connection status.
type VoiceState struct {
	GuildID string `json:"guild_id,omitempty"`
	ChannelID string `json:"channel_id,omitempty"`
	UserID string `json:"user_id,omitempty"`
	Member GuildMember `json:"member,omitempty"`
	SessionID string `json:"sessionID,omitempty"`
	Deaf bool `json:"deaf,omitempty"`
	Mute bool `json:"mute,omitempty"`
	SelfDeaf bool `json:"self_deaf,omitempty"`
	SelfMute bool `json:"self_mute,omitempty"`
	SelfStream bool `json:"self_stream,omitempty"`
	CameraOn bool `json:"self_video,omitempty"`
	Suppress bool `json:"suppress,omitempty"`
}

type VoiceRegion struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	VIP bool `json:"vip,omitempty"`
	Optimal bool `json:"optimal,omitempty"`
	Deprecated bool `json:"deprecated,omitempty"`
	Custom bool `json:"custom,omitempty"`
}

type PresenceUpdate struct {
	User User `json:"user,omitempty"`
	GuildID string `json:"guild_id,omitempty"`
	Status string `json:"status,omitempty"`
	Activities []Activity `json:"activities,omitempty"`
	ClientStatus ClientStatus `json:"client_status,omitempty"`
}

type ClientStatus struct {
	Desktop string `json:"desktop,omitempty"`
	Mobile string `json:"mobile,omitempty"`
	Web string `json:"web,omitempty"`
}

type Activity struct {
	Name string `json:"name,omitempty"`
	Type int `json:"type,omitempty"`
	URL string `json:"url,omitempty"`
	CreatedAt int `json:"created_at,omitempty"`
	Timestamps ActivityTimestamps `json:"timestamps,omitempty"`
	ApplicationID string `json:"application_id,omitempty"`
	Details string `json:"details,omitempty"`
	State string `json:"state,omitempty"`
	Emoji ActivityEmoji `json:"emoji,omitempty"`
	Party ActivityParty `json:"party,omitempty"`
	Assets ActivityAssets `json:"assets,omitempty"`
	Secrets ActivitySecrets `json:"secrets,omitempty"`
	Instance bool `json:"instance,omitempty"`
	Flags int `json:"flags,omitempty"`
}

type ActivityTimestamps struct {
	Start int `json:"start,omitempty"`
	End int `json:"end,omitempty"`
}

type ActivityParty struct {
	ID string `json:"id,omitempty"`
	Size [2]int `json:"size,omitempty"`
}

type ActivityEmoji struct {
	Name string `json:"name,omitempty"`
	ID string `json:"id,omitempty"`
	Animated bool `json:"animated,omitempty"`
}

type ActivityAssets struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText string `json:"small_text,omitempty"`
}

type ActivitySecrets struct {
	Join string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
	Match string `json:"match,omitempty"`
}

type GuildPreview struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Icon string `json:"icon,omitempty"`
	Splash string `json:"splash,omitempty"`
	DiscoverySplash string `json:"discovery_splash,omitempty"`
	Emojis []Emoji `json:"emojis,omitempty"`
	Features []string `json:"features,omitempty"`
	ApproxMemberCount int `json:"approximate_member_count,omitempty"`
	ApproxPresenceCount int `json:"approximate_presence_count,omitempty"`
	Description string `json:"description,omitempty"`
}

type GuildWidget struct {
	Enabled bool `json:"enabled,omitempty"`
	ChannelID string `json:"channel_id,omitempty"`
}


type Integration struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Syncing bool `json:"syncing,omitempty"`
	RoleID string `json:"role_id,omitempty"`
	EnableEmoticons bool `json:"enable_emoticons,omitempty"`
	ExpireBehavior int `json:"expire_behavior,omitempty"`
	ExpireGracePeriod int `json:"expire_grace_periond,omitempty"`
	User User `json:"user,omitempty"`
	Account IntegrationAccount `json:"account,omitempty"`
	SyncedAt string `json:"synced_at,omitempty"`
	SubscriberCount int `json:"subscriber_count,omitempty"`
	Revoked bool `json:"revoked,omitempty"`
	Application IntegrationApplication `json:"application,omitempty"`
}

type IntegrationAccount struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type IntegrationApplication struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Icon string `json:"icon,omitempty"`
	Description string `json:"description,omitempty"`
	Summary string `json:"summary,omitempty"`
	Bot User `json:"bot,omitempty"`
}

type Ban struct {
	Reason string `json:"reason,omitempty"`
	User User `json:"user,omitempty"`
}

type Invite struct {
	Code string `json:"code,omitempty"`
	Guild Guild `json:"guild,omitempty"`
	Channel Channel `json:"channel,omitempty"`
	Inviter User `json:"inviter,omitempty"`
	TargetUser User `json:"target_user,omitempty"`
	TargetUserType int `json:"target_user_type,omitempty"`
	ApproximatePresenceCount int `json:"approximate_presence_count,omitempty"`
	ApproximateMemberCount int `json:"approximate_member_count,omitempty"`
	Uses int `json:"uses,omitempty"`
	MaxUses int `json:"max_uses,omitempty"`
	MaxAge int `json:"max_age,omitempty"`
	Temporary bool `json:"temporary,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type CreateInviteObject struct {
	MaxAge int `json:"max_age,omitempty"`
	MaxUses int `json:"max_uses,omitempty"`
	TempMembership bool `json:"temporary`
	Unique bool `json:"unique,omitempty"`
	TargetUser string `json:"target_user,omitempty"`
	TargetUserType int `json:"target_user_type,omitempty"`
}

type Connection struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Revoked bool `json:"revoked,omitempty"`
	Integrations []Integration `json:"integrations,omitempty"`
	Verified bool `json:"verified,omitempty"`
	FriendSync bool `json:"friend_sync,omitempty"`
	ShowActivity bool `json:"show_activity,omitempty"`
	Visibility int `json:"visibility,omitempty"`
}

type Webhook struct {
	ID string `json:"id,omitempty"`
	Type int `json:"type,omitempty"`
	GuildID string `json:"guild_id,omitempty"`
	ChannelID string `json:"channe_id,omitempty"`
	User User `json:"user,omitempty"`
	Name string `json:"name,omitempty"`
	Avatar string `json:"avatar,omitempty"`
	Token string `json:"token,omitempty"`
	ApplicationID string `json:"application_id,omitempty"`
}

type AuditLog struct {
	Webhooks []Webhook `json:"webhooks,omitempty"`
	Users []User `json:"users,omitempty"`
	AuditLogEntries []AuditLogEntry `json:"audit_log_entries,omitempty"`
	Integrations []Integration `json:"integrations,omitempty"`
}

type AuditLogEntry struct {
	TargetID string `json:"target_id,omitempty"`
	Changes []AuditLogChange `json:"changes,omitempty"`
	UserID string `json:"user_id,omitempty"`
	ID string `json:"id,omitempty"`
	ActionType int `json:"action_type,omitempty"`
	Options OptionalAuditEntry `json:"options,omitempty"`
	Reason string `json:"reason,omitempty"`
}

type OptionalAuditEntry struct {
	DeleteMemberDays string `json:"delete_member_days,omitempty"`
	MembersRemoved string `json:"memebers_removed,omitempty"`
	ChannelID string `json:"channel_id,omitempty"`
	MessageID string `json:"message_id,omitempty"`
	Count string `json:"count,omitempty"`
	ID string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	RoleName string `json:"role_name,omitempty"`
}

type AuditLogChange struct {
	NewValue interface{}        `json:"new_value,omitempty"`
	OldValue interface{}        `json:"old_value,omitempty"`
	Key      *AuditLogChangeKey `json:"key,omitempty"`
}


type AuditLogChangeKey string

// Block of valid AuditLogChangeKey
const (
	AuditLogChangeKeyName                       AuditLogChangeKey = "name"
	AuditLogChangeKeyIconHash                   AuditLogChangeKey = "icon_hash"
	AuditLogChangeKeySplashHash                 AuditLogChangeKey = "splash_hash"
	AuditLogChangeKeyOwnerID                    AuditLogChangeKey = "owner_id"
	AuditLogChangeKeyRegion                     AuditLogChangeKey = "region"
	AuditLogChangeKeyAfkChannelID               AuditLogChangeKey = "afk_channel_id"
	AuditLogChangeKeyAfkTimeout                 AuditLogChangeKey = "afk_timeout"
	AuditLogChangeKeyMfaLevel                   AuditLogChangeKey = "mfa_level"
	AuditLogChangeKeyVerificationLevel          AuditLogChangeKey = "verification_level"
	AuditLogChangeKeyExplicitContentFilter      AuditLogChangeKey = "explicit_content_filter"
	AuditLogChangeKeyDefaultMessageNotification AuditLogChangeKey = "default_message_notifications"
	AuditLogChangeKeyVanityURLCode              AuditLogChangeKey = "vanity_url_code"
	AuditLogChangeKeyRoleAdd                    AuditLogChangeKey = "$add"
	AuditLogChangeKeyRoleRemove                 AuditLogChangeKey = "$remove"
	AuditLogChangeKeyPruneDeleteDays            AuditLogChangeKey = "prune_delete_days"
	AuditLogChangeKeyWidgetEnabled              AuditLogChangeKey = "widget_enabled"
	AuditLogChangeKeyWidgetChannelID            AuditLogChangeKey = "widget_channel_id"
	AuditLogChangeKeySystemChannelID            AuditLogChangeKey = "system_channel_id"
	AuditLogChangeKeyPosition                   AuditLogChangeKey = "position"
	AuditLogChangeKeyTopic                      AuditLogChangeKey = "topic"
	AuditLogChangeKeyBitrate                    AuditLogChangeKey = "bitrate"
	AuditLogChangeKeyPermissionOverwrite        AuditLogChangeKey = "permission_overwrites"
	AuditLogChangeKeyNSFW                       AuditLogChangeKey = "nsfw"
	AuditLogChangeKeyApplicationID              AuditLogChangeKey = "application_id"
	AuditLogChangeKeyRateLimitPerUser           AuditLogChangeKey = "rate_limit_per_user"
	AuditLogChangeKeyPermissions                AuditLogChangeKey = "permissions"
	AuditLogChangeKeyColor                      AuditLogChangeKey = "color"
	AuditLogChangeKeyHoist                      AuditLogChangeKey = "hoist"
	AuditLogChangeKeyMentionable                AuditLogChangeKey = "mentionable"
	AuditLogChangeKeyAllow                      AuditLogChangeKey = "allow"
	AuditLogChangeKeyDeny                       AuditLogChangeKey = "deny"
	AuditLogChangeKeyCode                       AuditLogChangeKey = "code"
	AuditLogChangeKeyChannelID                  AuditLogChangeKey = "channel_id"
	AuditLogChangeKeyInviterID                  AuditLogChangeKey = "inviter_id"
	AuditLogChangeKeyMaxUses                    AuditLogChangeKey = "max_uses"
	AuditLogChangeKeyUses                       AuditLogChangeKey = "uses"
	AuditLogChangeKeyMaxAge                     AuditLogChangeKey = "max_age"
	AuditLogChangeKeyTempoary                   AuditLogChangeKey = "temporary"
	AuditLogChangeKeyDeaf                       AuditLogChangeKey = "deaf"
	AuditLogChangeKeyMute                       AuditLogChangeKey = "mute"
	AuditLogChangeKeyNick                       AuditLogChangeKey = "nick"
	AuditLogChangeKeyAvatarHash                 AuditLogChangeKey = "avatar_hash"
	AuditLogChangeKeyID                         AuditLogChangeKey = "id"
	AuditLogChangeKeyType                       AuditLogChangeKey = "type"
	AuditLogChangeKeyEnableEmoticons            AuditLogChangeKey = "enable_emoticons"
	AuditLogChangeKeyExpireBehavior             AuditLogChangeKey = "expire_behavior"
	AuditLogChangeKeyExpireGracePeriod          AuditLogChangeKey = "expire_grace_period"
)