/*
Pingdom Public API

Testing ResultsAPIService

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

func Test_openapi_ResultsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ResultsAPIService ResultsCheckidGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var checkid int32

		resp, httpRes, err := apiClient.ResultsAPI.ResultsCheckidGet(context.Background(), checkid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
