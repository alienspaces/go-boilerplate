package harness

import (
	"github.com/brianvoe/gofakeit/v5"
	"github.com/google/uuid"

	"gitlab.com/alienspaces/go-boilerplate/server/service/character/internal/model"
	"gitlab.com/alienspaces/go-boilerplate/server/service/character/internal/record"
)

func (t *Testing) createCharacterRec(entityConfig CharacterConfig) (record.Character, error) {

	entityRec := entityConfig.Record

	if entityRec.CharacterType == "" {
		entityRec.CharacterType = record.CharacterTypePlayerMage
	}

	if entityRec.Name == "" {
		entityRec.Name = gofakeit.Name()
	}
	if entityRec.Avatar == "" {
		entityRec.Avatar = record.MageAvatarRedStripeDruid
	}

	t.Log.Info("Creating entity testing record >%#v<", entityRec)

	err := t.Model.(*model.Model).CreateCharacterRec(&entityRec)
	if err != nil {
		t.Log.Warn("Failed creating testing entity record >%v<", err)
		return entityRec, err
	}

	return entityRec, nil
}

func (t *Testing) createPlayerCharacterRec(entityRec record.Character, accountCharacterConfig PlayerCharacterConfig) (record.PlayerCharacter, error) {

	accountCharacterRec := accountCharacterConfig.Record

	t.Log.Info("Creating account entity testing record >%#v<", accountCharacterRec)

	if accountCharacterRec.PlayerID == "" {
		accountID, _ := uuid.NewRandom()
		accountCharacterRec.PlayerID = accountID.String()
	}
	accountCharacterRec.CharacterID = entityRec.ID

	err := t.Model.(*model.Model).CreatePlayerCharacterRec(&accountCharacterRec)
	if err != nil {
		t.Log.Warn("Failed creating testing account entity record >%v<", err)
		return accountCharacterRec, err
	}

	return accountCharacterRec, nil
}
