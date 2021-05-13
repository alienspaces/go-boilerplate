package harness

import (
	"gitlab.com/alienspaces/go-boilerplate/server/core/harness"
	"gitlab.com/alienspaces/go-boilerplate/server/core/type/modeller"
	"gitlab.com/alienspaces/go-boilerplate/server/service/character/internal/model"
	"gitlab.com/alienspaces/go-boilerplate/server/service/character/internal/record"
)

// Testing -
type Testing struct {
	harness.Testing
	Data       *Data
	DataConfig DataConfig
}

// DataConfig -
type DataConfig struct {
	CharacterConfig        []CharacterConfig
	PlayerCharacterConfig []PlayerCharacterConfig
}

// PlayerCharacterConfig -
type PlayerCharacterConfig struct {
	Record       record.PlayerCharacter
	CharacterConfig []CharacterConfig
}

// CharacterConfig -
type CharacterConfig struct {
	Record record.Character
}

// Data -
type Data struct {
	CharacterRecs        []record.Character
	PlayerCharacterRecs []record.PlayerCharacter
}

// NewTesting -
func NewTesting(config DataConfig) (t *Testing, err error) {

	// harness
	t = &Testing{}

	// modeller
	t.ModellerFunc = t.Modeller

	// data
	t.CreateDataFunc = t.CreateData
	t.RemoveDataFunc = t.RemoveData

	t.DataConfig = config
	t.Data = &Data{}

	err = t.Init()
	if err != nil {
		return nil, err
	}

	return t, nil
}

// Modeller -
func (t *Testing) Modeller() (modeller.Modeller, error) {

	m, err := model.NewModel(t.Config, t.Log, t.Store)
	if err != nil {
		t.Log.Warn("Failed new model >%v<", err)
		return nil, err
	}

	return m, nil
}

// CreateData - Custom data
func (t *Testing) CreateData() error {

	// Player entity with entity records
	for _, accountCharacterConfig := range t.DataConfig.PlayerCharacterConfig {

		// NOTE: For an account entity record to be created there must
		// be at least one entity record created.
		if len(accountCharacterConfig.CharacterConfig) == 0 {
			accountCharacterConfig.CharacterConfig = []CharacterConfig{
				{
					Record: record.Character{},
				},
			}
		}

		for _, entityConfig := range accountCharacterConfig.CharacterConfig {

			entityRec, err := t.createCharacterRec(entityConfig)
			if err != nil {
				t.Log.Warn("Failed creating entity record >%v<", err)
				return err
			}
			t.Data.CharacterRecs = append(t.Data.CharacterRecs, entityRec)

			accountCharacterRec, err := t.createPlayerCharacterRec(entityRec, accountCharacterConfig)
			if err != nil {
				t.Log.Warn("Failed creating account entity record >%v<", err)
				return err
			}
			t.Data.PlayerCharacterRecs = append(t.Data.PlayerCharacterRecs, accountCharacterRec)
		}
	}

	// Stand alone entity records
	for _, entityConfig := range t.DataConfig.CharacterConfig {

		entityRec, err := t.createCharacterRec(entityConfig)
		if err != nil {
			t.Log.Warn("Failed creating entity record >%v<", err)
			return err
		}
		t.Data.CharacterRecs = append(t.Data.CharacterRecs, entityRec)
	}

	return nil
}

// RemoveData -
func (t *Testing) RemoveData() error {

ACCOUNT_RECS:
	for {
		if len(t.Data.PlayerCharacterRecs) == 0 {
			break ACCOUNT_RECS
		}
		rec := record.PlayerCharacter{}
		rec, t.Data.PlayerCharacterRecs = t.Data.PlayerCharacterRecs[0], t.Data.PlayerCharacterRecs[1:]

		err := t.Model.(*model.Model).RemovePlayerCharacterRec(rec.ID)
		if err != nil {
			t.Log.Warn("Failed removing account entity record >%v<", err)
			return err
		}
	}

ENTITY_RECS:
	for {
		if len(t.Data.CharacterRecs) == 0 {
			break ENTITY_RECS
		}
		rec := record.Character{}
		rec, t.Data.CharacterRecs = t.Data.CharacterRecs[0], t.Data.CharacterRecs[1:]

		err := t.Model.(*model.Model).RemoveCharacterRec(rec.ID)
		if err != nil {
			t.Log.Warn("Failed removing entity record >%v<", err)
			return err
		}
	}

	return nil
}
