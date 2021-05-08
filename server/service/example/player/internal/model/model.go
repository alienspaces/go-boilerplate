package model

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/alienspaces/go-mono-api-boilerplate/server/core/model"
	"gitlab.com/alienspaces/go-mono-api-boilerplate/server/core/type/configurer"
	"gitlab.com/alienspaces/go-mono-api-boilerplate/server/core/type/logger"
	"gitlab.com/alienspaces/go-mono-api-boilerplate/server/core/type/preparer"
	"gitlab.com/alienspaces/go-mono-api-boilerplate/server/core/type/repositor"
	"gitlab.com/alienspaces/go-mono-api-boilerplate/server/core/type/storer"

	"gitlab.com/alienspaces/go-mono-api-boilerplate/server/service/player/internal/repository/player"
	"gitlab.com/alienspaces/go-mono-api-boilerplate/server/service/player/internal/repository/playerrole"
)

// Model -
type Model struct {
	model.Model
	// Allows auth token verification to be mocked for testing
	AuthVerifyTokenFunc func(provider, token string) (*VerifiedData, error)
}

// NewModel -
func NewModel(c configurer.Configurer, l logger.Logger, s storer.Storer) (*Model, error) {

	m := &Model{
		Model: model.Model{
			Config: c,
			Log:    l,
			Store:  s,
		},
		AuthVerifyTokenFunc: nil,
	}

	m.AuthVerifyTokenFunc = m.authVerifyToken
	m.RepositoriesFunc = m.NewRepositories

	return m, nil
}

// NewRepositories - Custom repositories for this model
func (m *Model) NewRepositories(p preparer.Preparer, tx *sqlx.Tx) ([]repositor.Repositor, error) {

	repositoryList := []repositor.Repositor{}

	accountRepo, err := account.NewRepository(m.Log, p, tx)
	if err != nil {
		m.Log.Warn("Failed new account repository >%v<", err)
		return nil, err
	}

	repositoryList = append(repositoryList, accountRepo)

	accountRoleRepo, err := accountrole.NewRepository(m.Log, p, tx)
	if err != nil {
		m.Log.Warn("Failed new account role repository >%v<", err)
		return nil, err
	}

	repositoryList = append(repositoryList, accountRoleRepo)

	return repositoryList, nil
}

// PlayerRepository -
func (m *Model) PlayerRepository() *account.Repository {

	r := m.Repositories[account.TableName]
	if r == nil {
		m.Log.Warn("Repository >%s< is nil", account.TableName)
		return nil
	}

	return r.(*account.Repository)
}

// PlayerRoleRepository -
func (m *Model) PlayerRoleRepository() *accountrole.Repository {

	r := m.Repositories[accountrole.TableName]
	if r == nil {
		m.Log.Warn("Repository >%s< is nil", accountrole.TableName)
		return nil
	}

	return r.(*accountrole.Repository)
}
