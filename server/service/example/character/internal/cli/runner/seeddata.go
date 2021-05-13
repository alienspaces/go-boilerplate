package runner

import (
	"github.com/urfave/cli/v2"

	"gitlab.com/alienspaces/go-boilerplate/server/service/character/internal/harness"
	"gitlab.com/alienspaces/go-boilerplate/server/service/character/internal/record"
)

// LoadSeedData -
func (rnr *Runner) LoadSeedData(c *cli.Context) error {

	rnr.Log.Info("** Load Seed Data **")

	// harness
	config := harness.DataConfig{

		CharacterConfig: []harness.CharacterConfig{
			// Mage - Dark Armoured
			{
				Record: record.Character{
					CharacterType:   record.CharacterTypeStarterMage,
					Name:         "Dark Armoured",
					Avatar:       record.MageAvatarDarkArmoured,
					Strength:     16,
					Dexterity:    12,
					Intelligence: 10,
					Coins:        50,
				},
			},
			// Mage - Red Stripe Druid
			{
				Record: record.Character{
					CharacterType:   record.CharacterTypeStarterMage,
					Name:         "Druid",
					Avatar:       record.MageAvatarRedStripeDruid,
					Strength:     14,
					Dexterity:    14,
					Intelligence: 10,
					Coins:        50,
				},
			},
			// Mage - Red Fairy
			{
				Record: record.Character{
					CharacterType:   record.CharacterTypeStarterMage,
					Name:         "Fairy",
					Avatar:       record.MageAvatarRedFairy,
					Strength:     10,
					Dexterity:    14,
					Intelligence: 14,
					Coins:        50,
				},
			},
			// Mage - Red Stripe Necromancer
			{
				Record: record.Character{
					CharacterType:   record.CharacterTypeStarterMage,
					Name:         "Necromancer",
					Avatar:       record.MageAvatarRedStripeNecromancer,
					Strength:     14,
					Dexterity:    10,
					Intelligence: 14,
					Coins:        50,
				},
			},
			// Mage - Green Elven
			{
				Record: record.Character{
					CharacterType:   record.CharacterTypeStarterMage,
					Name:         "Elven",
					Avatar:       record.MageAvatarGreenElven,
					Strength:     12,
					Dexterity:    14,
					Intelligence: 12,
					Coins:        50,
				},
			},
			// Familliar - Brown Cyclops Bat
			{
				Record: record.Character{
					CharacterType:   record.CharacterTypeStarterFamilliar,
					Name:         "Brown Cyclops Bat",
					Avatar:       record.FamilliarAvatarBrownCyclopsBat,
					Strength:     10,
					Dexterity:    10,
					Intelligence: 10,
				},
			},
			// Familliar - Brown Yeti
			{
				Record: record.Character{
					CharacterType:   record.CharacterTypeStarterFamilliar,
					Name:         "Brown Yeti",
					Avatar:       record.FamilliarAvatarBrownYeti,
					Strength:     10,
					Dexterity:    10,
					Intelligence: 10,
				},
			},
			// Familliar - Green Tribble
			{
				Record: record.Character{
					CharacterType:   record.CharacterTypeStarterFamilliar,
					Name:         "Green Tribble",
					Avatar:       record.FamilliarAvatarGreenTribble,
					Strength:     10,
					Dexterity:    10,
					Intelligence: 10,
				},
			},
			// Familliar - Grey Cyclops
			{
				Record: record.Character{
					CharacterType:   record.CharacterTypeStarterFamilliar,
					Name:         "Grey Cyclops",
					Avatar:       record.FamilliarAvatarGreyCyclops,
					Strength:     10,
					Dexterity:    10,
					Intelligence: 10,
				},
			},
			// Familliar - Orange Spotted Tribble
			{
				Record: record.Character{
					CharacterType:   record.CharacterTypeStarterFamilliar,
					Name:         "Orange Spotten Tribble",
					Avatar:       record.FamilliarAvatarOrangeSpottedTribble,
					Strength:     10,
					Dexterity:    10,
					Intelligence: 10,
				},
			},
			// Familliar - Purple Bat
			{
				Record: record.Character{
					CharacterType:   record.CharacterTypeStarterFamilliar,
					Name:         "Purple Bat",
					Avatar:       record.FamilliarAvatarPurpleBat,
					Strength:     10,
					Dexterity:    10,
					Intelligence: 10,
				},
			},
			// Familliar - Purple Minotaur
			{
				Record: record.Character{
					CharacterType:   record.CharacterTypeStarterFamilliar,
					Name:         "Purple Minotaur",
					Avatar:       record.FamilliarAvatarPurpleMinotaur,
					Strength:     10,
					Dexterity:    10,
					Intelligence: 10,
				},
			},
		},
	}

	h, err := harness.NewTesting(config)
	if err != nil {
		rnr.Log.Warn("Failed new testing harness >%v<", err)
		return err
	}

	// harness commit data
	h.CommitData = true

	err = h.Setup()
	if err != nil {
		rnr.Log.Warn("Failed testing harness setup >%v<", err)
		return err
	}

	rnr.Log.Info("All done")

	return nil
}
