package record

import (
	"gitlab.com/alienspaces/go-boilerplate/server/core/repository"
)

// PlayerProvider - Valid account providers
const (
	PlayerProviderAnonymous string = "anonymous"
	PlayerProviderApple     string = "apple"
	PlayerProviderFacebook  string = "facebook"
	PlayerProviderGithub    string = "github"
	PlayerProviderGoogle    string = "google"
	PlayerProviderTwitter   string = "twitter"
)

// Player -
type Player struct {
	repository.Record
	Name              string `db:"name"`
	Email             string `db:"email"`
	Provider          string `db:"provider"`
	ProviderPlayerID string `db:"provider_account_id"`
}

// PlayerRole - Valid roles
const (
	PlayerRoleDefault       string = "default"
	PlayerRoleAdministrator string = "administrator"
)

// PlayerRole -
type PlayerRole struct {
	repository.Record
	PlayerID string `db:"account_id"`
	Role      string `db:"role"`
}
