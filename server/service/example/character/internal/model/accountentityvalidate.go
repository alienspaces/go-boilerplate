package model

import (
	"fmt"

	"gitlab.com/alienspaces/go-mono-api-boilerplate/server/service/character/internal/record"
)

// ValidatePlayerCharacterRec - validates creating and updating a mage record
func (m *Model) ValidatePlayerCharacterRec(rec *record.PlayerCharacter) error {

	// required fields
	if rec.PlayerID == "" {
		return fmt.Errorf("PlayerID is required")
	}
	if rec.CharacterID == "" {
		return fmt.Errorf("CharacterID is required")
	}

	return nil
}

// ValidateDeletePlayerCharacterRec - validates it is okay to delete a mage record
func (m *Model) ValidateDeletePlayerCharacterRec(recID string) error {

	return nil
}
