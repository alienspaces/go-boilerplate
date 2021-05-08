package model

import (
	"fmt"

	"gitlab.com/alienspaces/go-mono-api-boilerplate/server/service/player/internal/record"
)

// ValidatePlayerRec - validates creating and updating an account record
func (m *Model) ValidatePlayerRec(rec *record.Player) error {

	switch rec.Provider {
	case record.PlayerProviderAnonymous:
		// We only require a provider account ID to create an anonymous local account
		if rec.ProviderPlayerID == "" {
			msg := "Missing ProviderPlayerID, cannot create an anonymous account"
			m.Log.Warn(msg)
			return fmt.Errorf(msg)
		}
	case record.PlayerProviderGoogle:
		// We require a provider account ID and email to create a Google local account
		if rec.Email == "" {
			msg := "Missing Email, cannot create a Google account"
			m.Log.Warn(msg)
			return fmt.Errorf(msg)
		}
		if rec.ProviderPlayerID == "" {
			msg := "Missing ProviderPlayerID, cannot create a Google account"
			m.Log.Warn(msg)
			return fmt.Errorf(msg)
		}
	default:
		// no-op
	}
	return nil
}

// ValidateDeletePlayerRec - validates it is okay to delete a account record
func (m *Model) ValidateDeletePlayerRec(recID string) error {

	return nil
}
