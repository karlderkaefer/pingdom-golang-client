/*
Pingdom Public API

Testing SummaryAverageAPIService

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

func Test_openapi_SummaryAverageAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SummaryAverageAPIService SummaryAverageCheckidGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var checkid int32

		resp, httpRes, err := apiClient.SummaryAverageAPI.SummaryAverageCheckidGet(context.Background(), checkid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
