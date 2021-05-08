package model

import (
	"database/sql"
	"fmt"

	"gitlab.com/alienspaces/go-mono-api-boilerplate/server/service/player/internal/record"
)

// GetPlayerRoleRecs -
func (m *Model) GetPlayerRoleRecs(params map[string]interface{}, operators map[string]string, forUpdate bool) ([]*record.PlayerRole, error) {

	m.Log.Info("Getting account records params >%s<", params)

	r := m.PlayerRoleRepository()

	return r.GetMany(params, operators, forUpdate)
}

// GetPlayerRoleRec -
func (m *Model) GetPlayerRoleRec(recID string, forUpdate bool) (*record.PlayerRole, error) {

	m.Log.Info("Getting account rec ID >%s<", recID)

	r := m.PlayerRoleRepository()

	// validate UUID
	if !m.IsUUID(recID) {
		return nil, fmt.Errorf("ID >%s< is not a valid UUID", recID)
	}

	rec, err := r.GetOne(recID, forUpdate)
	if err == sql.ErrNoRows {
		m.Log.Warn("No record found ID >%s<", recID)
		return nil, nil
	}

	return rec, err
}

// CreatePlayerRoleRec -
func (m *Model) CreatePlayerRoleRec(rec *record.PlayerRole) error {

	m.Log.Info("Creating account rec >%v<", rec)

	r := m.PlayerRoleRepository()

	err := m.ValidatePlayerRoleRec(rec)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.CreateOne(rec)
}

// UpdatePlayerRoleRec -
func (m *Model) UpdatePlayerRoleRec(rec *record.PlayerRole) error {

	m.Log.Info("Updating account rec >%v<", rec)

	r := m.PlayerRoleRepository()

	err := m.ValidatePlayerRoleRec(rec)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.UpdateOne(rec)
}

// DeletePlayerRoleRec -
func (m *Model) DeletePlayerRoleRec(recID string) error {

	m.Log.Info("Deleting account rec ID >%s<", recID)

	r := m.PlayerRoleRepository()

	// validate UUID
	if !m.IsUUID(recID) {
		return fmt.Errorf("ID >%s< is not a valid UUID", recID)
	}

	err := m.ValidateDeletePlayerRoleRec(recID)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.DeleteOne(recID)
}

// RemovePlayerRoleRec -
func (m *Model) RemovePlayerRoleRec(recID string) error {

	m.Log.Info("Removing account rec ID >%s<", recID)

	r := m.PlayerRoleRepository()

	// validate UUID
	if !m.IsUUID(recID) {
		return fmt.Errorf("ID >%s< is not a valid UUID", recID)
	}

	err := m.ValidateDeletePlayerRoleRec(recID)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.RemoveOne(recID)
}
