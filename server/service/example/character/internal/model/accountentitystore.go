package model

import (
	"database/sql"
	"fmt"

	"gitlab.com/alienspaces/go-boilerplate/server/service/character/internal/record"
)

// GetPlayerCharacterRecs -
func (m *Model) GetPlayerCharacterRecs(params map[string]interface{}, operators map[string]string, forUpdate bool) ([]*record.PlayerCharacter, error) {

	m.Log.Info("Getting account character records params >%s<", params)

	r := m.PlayerCharacterRepository()

	return r.GetMany(params, operators, forUpdate)
}

// GetPlayerCharacterRec -
func (m *Model) GetPlayerCharacterRec(recID string, forUpdate bool) (*record.PlayerCharacter, error) {

	m.Log.Info("Getting account character rec ID >%s<", recID)

	r := m.PlayerCharacterRepository()

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

// CreatePlayerCharacterRec -
func (m *Model) CreatePlayerCharacterRec(rec *record.PlayerCharacter) error {

	m.Log.Info("Creating account character rec >%#v<", rec)

	r := m.PlayerCharacterRepository()

	err := m.ValidatePlayerCharacterRec(rec)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.CreateOne(rec)
}

// UpdatePlayerCharacterRec -
func (m *Model) UpdatePlayerCharacterRec(rec *record.PlayerCharacter) error {

	m.Log.Info("Updating account character rec >%#v<", rec)

	r := m.PlayerCharacterRepository()

	err := m.ValidatePlayerCharacterRec(rec)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.UpdateOne(rec)
}

// DeletePlayerCharacterRec -
func (m *Model) DeletePlayerCharacterRec(recID string) error {

	m.Log.Info("Deleting account character rec ID >%s<", recID)

	r := m.PlayerCharacterRepository()

	// validate UUID
	if !m.IsUUID(recID) {
		return fmt.Errorf("ID >%s< is not a valid UUID", recID)
	}

	err := m.ValidateDeletePlayerCharacterRec(recID)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.DeleteOne(recID)
}

// RemovePlayerCharacterRec -
func (m *Model) RemovePlayerCharacterRec(recID string) error {

	m.Log.Info("Removing account character rec ID >%s<", recID)

	r := m.PlayerCharacterRepository()

	// validate UUID
	if !m.IsUUID(recID) {
		return fmt.Errorf("ID >%s< is not a valid UUID", recID)
	}

	err := m.ValidateDeletePlayerCharacterRec(recID)
	if err != nil {
		m.Log.Info("Failed model validation >%v<", err)
		return err
	}

	return r.RemoveOne(recID)
}
