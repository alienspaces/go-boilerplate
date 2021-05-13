package record

import (
	"gitlab.com/alienspaces/go-boilerplate/server/core/repository"
)

// CharacterType -
const (
	CharacterTypeMage             string = "mage"
	CharacterTypeFamilliar        string = "familliar"
	CharacterTypePlayerMage       string = "player-mage"
	CharacterTypePlayerFamilliar  string = "player-familliar"
	CharacterTypeStarterMage      string = "starter-mage"
	CharacterTypeStarterFamilliar string = "starter-familliar"
)

// MageAvatar -
const (
	MageAvatarDarkArmoured         string = "dark-armoured"
	MageAvatarRedStripeDruid       string = "red-stripe-druid"
	MageAvatarRedFairy             string = "red-fairy"
	MageAvatarRedStripeNecromancer string = "red-stripe-necromancer"
	MageAvatarGreenElven           string = "green-elven"
)

// FamilliarAvatar -
const (
	FamilliarAvatarBrownCyclopsBat      string = "brown-cyclops-bat"
	FamilliarAvatarBrownYeti            string = "brown-yeti"
	FamilliarAvatarGreenTribble         string = "green-tribble"
	FamilliarAvatarGreyCyclops          string = "grey-cyclops"
	FamilliarAvatarOrangeSpottedTribble string = "orange-spotted-tribble"
	FamilliarAvatarPurpleBat            string = "purple-bat"
	FamilliarAvatarPurpleMinotaur       string = "purple-minotaur"
)

// Character -
type Character struct {
	repository.Record
	CharacterType       string `db:"entity_type"`
	Name             string `db:"name"`
	Avatar           string `db:"avatar"`
	Strength         int    `db:"strength"`
	Dexterity        int    `db:"dexterity"`
	Intelligence     int    `db:"intelligence"`
	AttributePoints  int64  `db:"attribute_points"`
	ExperiencePoints int64  `db:"experience_points"`
	Coins            int64  `db:"coins"`
}

// PlayerCharacter -
type PlayerCharacter struct {
	repository.Record
	PlayerID string `db:"account_id"`
	CharacterID  string `db:"entity_id"`
}

// CharacterItem -
type CharacterItem struct {
	repository.Record
	ItemID   string `db:"item_id"`
	CharacterID string `db:"entity_id"`
}

// CharacterSpell -
type CharacterSpell struct {
	repository.Record
	SpellID  string `db:"spell_id"`
	CharacterID string `db:"entity_id"`
}
