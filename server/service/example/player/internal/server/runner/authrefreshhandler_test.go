package runner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/require"

	"gitlab.com/alienspaces/go-boilerplate/server/constant"
	"gitlab.com/alienspaces/go-boilerplate/server/core/auth"
	"gitlab.com/alienspaces/go-boilerplate/server/core/server"
	"gitlab.com/alienspaces/go-boilerplate/server/schema"
	"gitlab.com/alienspaces/go-boilerplate/server/service/player/internal/harness"
)

func TestAuthRefreshHandler(t *testing.T) {

	// Test harness
	th, err := NewTestHarness()
	require.NoError(t, err, "New test data returns without error")

	//  Test dependencies
	c, l, s, err := th.NewDefaultDependencies()
	require.NoError(t, err, "NewDefaultDependencies returns without error")

	type TestCase struct {
		name          string
		config        func(rnr *Runner) server.HandlerConfig
		requestParams func(data *harness.Data) map[string]string
		queryParams   func(data *harness.Data) map[string]string
		requestData   func(data *harness.Data) *schema.AuthRefreshRequest
		responseCode  int
		responseData  func(data *harness.Data) *schema.AuthRefreshResponse
	}

	getToken := func(data *harness.Data) (string, error) {

		a, err := auth.NewAuth(c, l)
		if err != nil {
			return "", err
		}

		// TODO: Expand on account roles
		roles := []string{
			constant.AuthRoleDefault,
		}

		identity := map[string]interface{}{
			constant.AuthIdentityPlayerID: data.PlayerRecs[0].ID,
		}

		claims := auth.Claims{
			Roles:    roles,
			Identity: identity,
		}

		tokenString, err := a.EncodeJWT(&claims)
		if err != nil {
			return "", err
		}

		return tokenString, nil
	}

	tests := []TestCase{
		// Refresh Auth
		{
			name: "POST - Refresh succeeds",
			config: func(rnr *Runner) server.HandlerConfig {
				return rnr.HandlerConfig[1]
			},
			requestData: func(data *harness.Data) *schema.AuthRefreshRequest {

				token, _ := getToken(data)
				res := schema.AuthRefreshRequest{
					Data: schema.AuthRefreshData{
						Token: token,
					},
				}
				return &res
			},
			responseCode: http.StatusOK,
			responseData: func(data *harness.Data) *schema.AuthRefreshResponse {
				res := schema.AuthRefreshResponse{
					Data: []schema.AuthRefreshData{
						{
							PlayerID:    data.PlayerRecs[0].ID,
							PlayerEmail: data.PlayerRecs[0].Email,
							PlayerName:  data.PlayerRecs[0].Name,
						},
					},
				}
				return &res
			},
		},
	}

	for _, tc := range tests {

		t.Logf("Running test >%s<", tc.name)

		func() {
			rnr := NewRunner()

			err = rnr.Init(c, l, s)
			require.NoError(t, err, "Runner init returns without error")

			err = th.Setup()
			require.NoError(t, err, "Test data setup returns without error")
			defer func() {
				err = th.Teardown()
				require.NoError(t, err, "Test data teardown returns without error")
			}()

			// config
			cfg := tc.config(rnr)

			// handler
			h, _ := rnr.DefaultMiddleware(cfg, cfg.HandlerFunc)

			// router
			rtr := httprouter.New()

			switch cfg.Method {
			case http.MethodGet:
				rtr.GET(cfg.Path, h)
			case http.MethodPost:
				rtr.POST(cfg.Path, h)
			case http.MethodPut:
				rtr.PUT(cfg.Path, h)
			case http.MethodDelete:
				rtr.DELETE(cfg.Path, h)
			default:
				//
			}

			// request params
			requestParams := map[string]string{}
			if tc.requestParams != nil {
				requestParams = tc.requestParams(th.Data)
			}

			requestPath := cfg.Path
			for paramKey, paramValue := range requestParams {
				requestPath = strings.Replace(requestPath, paramKey, paramValue, 1)
			}

			// query params
			queryParams := map[string]string{}
			if tc.queryParams != nil {
				queryParams = tc.queryParams(th.Data)
			}

			if len(queryParams) > 0 {
				count := 0
				for paramKey, paramValue := range queryParams {
					if count == 0 {
						requestPath = requestPath + `?`
					} else {
						requestPath = requestPath + `&`
					}
					t.Logf("Adding parameter key >%s< param >%s<", paramKey, paramValue)
					requestPath = fmt.Sprintf("%s%s=%s", requestPath, paramKey, url.QueryEscape(paramValue))
				}
				t.Logf("Resulting requestPath >%s<", requestPath)
			}

			// request data
			data := tc.requestData(th.Data)

			var req *http.Request

			if data != nil {
				jsonData, err := json.Marshal(data)
				require.NoError(t, err, "Marshal returns without error")

				req, err = http.NewRequest(cfg.Method, requestPath, bytes.NewBuffer(jsonData))
				require.NoError(t, err, "NewRequest returns without error")
			} else {
				req, err = http.NewRequest(cfg.Method, requestPath, nil)
				require.NoError(t, err, "NewRequest returns without error")
			}

			// recorder
			rec := httptest.NewRecorder()

			// serve
			rtr.ServeHTTP(rec, req)

			// test status
			require.Equalf(t, tc.responseCode, rec.Code, "%s - Response code equals expected", tc.name)

			res := schema.AuthRefreshResponse{}
			err = json.NewDecoder(rec.Body).Decode(&res)
			require.NoError(t, err, "Decode returns without error")

			// response data
			var resData *schema.AuthRefreshResponse
			if tc.responseData != nil {
				resData = tc.responseData(th.Data)
			}

			// test data
			if tc.responseCode == http.StatusOK {

				// response data
				if resData != nil {
					require.NotEmpty(t, res.Data[0].Token, "Token is not empty")
					require.Equal(t, resData.Data[0].PlayerID, res.Data[0].PlayerID, "PlayerID equals expected")
					require.Equal(t, resData.Data[0].PlayerEmail, res.Data[0].PlayerEmail, "PlayerEmail equals expected")
					require.Equal(t, resData.Data[0].PlayerName, res.Data[0].PlayerName, "PlayerName equals expected")
				}

				// record timestamps
				require.False(t, res.Data[0].CreatedAt.IsZero(), "CreatedAt is not zero")
				if cfg.Method == http.MethodPost {
					require.True(t, res.Data[0].UpdatedAt.IsZero(), "UpdatedAt is zero")
				}
				if cfg.Method == http.MethodPut {
					require.False(t, res.Data[0].UpdatedAt.IsZero(), "UpdatedAt is not zero")
				}
			}
		}()
	}
}
