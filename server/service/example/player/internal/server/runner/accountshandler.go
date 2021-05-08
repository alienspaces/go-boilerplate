package runner

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"gitlab.com/alienspaces/go-mono-api-boilerplate/server/core/type/logger"
	"gitlab.com/alienspaces/go-mono-api-boilerplate/server/core/type/modeller"
	"gitlab.com/alienspaces/go-mono-api-boilerplate/server/schema"
	"gitlab.com/alienspaces/go-mono-api-boilerplate/server/service/player/internal/model"
	"gitlab.com/alienspaces/go-mono-api-boilerplate/server/service/player/internal/record"
)

// GetPlayersHandler -
func (rnr *Runner) GetPlayersHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Get accounts handler ** p >%#v< m >%#v<", pp, m)

	var recs []*record.Player
	var err error

	id := pp.ByName("account_id")

	// single resource
	if id != "" {

		l.Info("Getting account record ID >%s<", id)

		rec, err := m.(*model.Model).GetPlayerRec(id, false)
		if err != nil {
			rnr.WriteModelError(l, w, err)
			return
		}

		// resource not found
		if rec == nil {
			rnr.WriteNotFoundError(l, w, id)
			return
		}

		recs = append(recs, rec)

	} else {

		l.Info("Querying account records")

		params := make(map[string]interface{})
		for paramName, paramValue := range qp {
			params[paramName] = paramValue
		}

		recs, err = m.(*model.Model).GetPlayerRecs(params, nil, false)
		if err != nil {
			rnr.WriteModelError(l, w, err)
			return
		}
	}

	// assign response properties
	data := []schema.PlayerData{}
	for _, rec := range recs {

		// response data
		responseData, err := rnr.RecordToPlayerResponseData(rec)
		if err != nil {
			rnr.WriteSystemError(l, w, err)
			return
		}

		data = append(data, responseData)
	}

	res := schema.PlayerResponse{
		Data: data,
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// PostPlayersHandler -
func (rnr *Runner) PostPlayersHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Post accounts handler ** p >%#v< m >#%v<", pp, m)

	// parameters
	id := pp.ByName("account_id")

	req := schema.PlayerRequest{}

	err := rnr.ReadRequest(l, r, &req)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	rec := record.Player{}

	// assign request properties
	rec.ID = id

	// record data
	err = rnr.PlayerRequestDataToRecord(req.Data, &rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	err = m.(*model.Model).CreatePlayerRec(&rec)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// response data
	responseData, err := rnr.RecordToPlayerResponseData(&rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// assign response properties
	res := schema.PlayerResponse{
		Data: []schema.PlayerData{
			responseData,
		},
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// PutPlayersHandler -
func (rnr *Runner) PutPlayersHandler(w http.ResponseWriter, r *http.Request, pp httprouter.Params, qp map[string]interface{}, l logger.Logger, m modeller.Modeller) {

	l.Info("** Put accounts handler ** p >%#v< m >#%v<", pp, m)

	// parameters
	id := pp.ByName("account_id")

	l.Info("Updating resource ID >%s<", id)

	rec, err := m.(*model.Model).GetPlayerRec(id, false)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// resource not found
	if rec == nil {
		rnr.WriteNotFoundError(l, w, id)
		return
	}

	req := schema.PlayerRequest{}

	err = rnr.ReadRequest(l, r, &req)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// record data
	err = rnr.PlayerRequestDataToRecord(req.Data, rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	err = m.(*model.Model).UpdatePlayerRec(rec)
	if err != nil {
		rnr.WriteModelError(l, w, err)
		return
	}

	// response data
	responseData, err := rnr.RecordToPlayerResponseData(rec)
	if err != nil {
		rnr.WriteSystemError(l, w, err)
		return
	}

	// assign response properties
	res := schema.PlayerResponse{
		Data: []schema.PlayerData{
			responseData,
		},
	}

	err = rnr.WriteResponse(l, w, res)
	if err != nil {
		l.Warn("Failed writing response >%v<", err)
		return
	}
}

// PlayerRequestDataToRecord -
func (rnr *Runner) PlayerRequestDataToRecord(data schema.PlayerData, rec *record.Player) error {

	rec.Name = data.Name
	rec.Email = data.Email
	rec.Provider = data.Provider
	rec.ProviderPlayerID = data.ProviderPlayerID

	return nil
}

// RecordToPlayerResponseData -
func (rnr *Runner) RecordToPlayerResponseData(accountRec *record.Player) (schema.PlayerData, error) {

	data := schema.PlayerData{
		ID:                accountRec.ID,
		Name:              accountRec.Name,
		Email:             accountRec.Email,
		Provider:          accountRec.Provider,
		ProviderPlayerID: accountRec.ProviderPlayerID,
		CreatedAt:         accountRec.CreatedAt,
		UpdatedAt:         accountRec.UpdatedAt.Time,
	}

	return data, nil
}
