package model

import (
	"fmt"

	"gitlab.com/alienspaces/go-mono-api-boilerplate/server/service/character/internal/record"
)

// ValidateCharacterRec - validates creating and updating a mage record
func (m *Model) ValidateCharacterRec(rec *record.Character) error {

	// required fields
	if rec.CharacterType == "" {
		return fmt.Errorf("property CharacterType is required")
	}
	if rec.Name == "" {
		return fmt.Errorf("property Name is required")
	}
	if rec.Avatar == "" {
		return fmt.Errorf("property Avatar is required")
	}

	return nil
}

// ValidateDeleteCharacterRec - validates it is okay to delete a mage record
func (m *Model) ValidateDeleteCharacterRec(recID string) error {

	return nil
}
