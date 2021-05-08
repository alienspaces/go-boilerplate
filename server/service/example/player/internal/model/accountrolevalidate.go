package model

import (
	"gitlab.com/alienspaces/go-mono-api-boilerplate/server/service/player/internal/record"
)

// ValidatePlayerRoleRec - validates creating and updating a account record
func (m *Model) ValidatePlayerRoleRec(rec *record.PlayerRole) error {

	return nil
}

// ValidateDeletePlayerRoleRec - validates it is okay to delete a account record
func (m *Model) ValidateDeletePlayerRoleRec(recID string) error {

	return nil
}
