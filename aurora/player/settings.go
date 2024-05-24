package player

type ClientSettings struct {
	Locale               string // e.g. en_US
	ViewDistance         int8   // Client-side render distance, in chunks.
	ChatMode             int32  // 0: enabled, 1: commands only, 2: hidden.
	ChatColors           bool   // Can chat be colored?
	DisplayedSkinParts   uint8  // Flags in a comment below
	MainHand             int32  // 0: Left, 1: Right.
	TextFilteringEnabled bool   // Enables filtering of text on signs and written book titles.
	AllowServerListing   bool   // Should the player show in the status player sample
}
