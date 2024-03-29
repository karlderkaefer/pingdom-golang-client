/*
Pingdom Public API

Testing TeamsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/karlderkaefer/pingdom-golang-client/pkg/pingdom/openapi"
)

func Test_openapi_TeamsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TeamsAPIService AlertingTeamsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TeamsAPI.AlertingTeamsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TeamsAPIService AlertingTeamsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TeamsAPI.AlertingTeamsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TeamsAPIService AlertingTeamsTeamidDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var teamid int32

		resp, httpRes, err := apiClient.TeamsAPI.AlertingTeamsTeamidDelete(context.Background(), teamid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TeamsAPIService AlertingTeamsTeamidGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var teamid int32

		resp, httpRes, err := apiClient.TeamsAPI.AlertingTeamsTeamidGet(context.Background(), teamid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TeamsAPIService AlertingTeamsTeamidPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var teamid int32

		resp, httpRes, err := apiClient.TeamsAPI.AlertingTeamsTeamidPut(context.Background(), teamid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
