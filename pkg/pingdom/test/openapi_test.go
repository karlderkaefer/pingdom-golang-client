package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	pingdom "github.com/karlderkaefer/pingdom-golang-client/pkg/pingdom/openapi"
	"github.com/stretchr/testify/suite"
)

type APITestSuite struct {
	suite.Suite
	apiClient *pingdom.APIClient
}

func (suite *APITestSuite) SetupSuite() {
	cfg := pingdom.NewConfiguration()
	cfg.AddApiToken(os.Getenv("PINGDOM_API_TOKEN"))
	suite.apiClient = pingdom.NewAPIClient(cfg)
}

func (suite *APITestSuite) TestGetCheck() {
	// Send the request to get the check
	checks, _, err := suite.apiClient.ChecksAPI.ChecksGet(context.Background()).Execute()
	suite.NoError(err)

	// Print the created check details
	for _, check := range checks.Checks {
		suite.T().Log(*check.Name)
	}
}

func (suite *APITestSuite) TestGetTransactionCheck() {
	// Send the request to get the check
	checks, _, err := suite.apiClient.TMSChecksAPI.GetAllChecks(context.Background()).Execute()
	suite.NoError(err)

	// Print the created check details
	for _, check := range checks.Checks {
		suite.T().Log(*check.Name)
	}

}

func (suite *APITestSuite) TestCreateTransactionCheck() {
	url := "https://www.google.com"
	step := pingdom.NewStep()
	step.SetArgs(pingdom.StepArgs{Url: &url})
	step.SetFn("go_to")
	checkWithoutID := *pingdom.NewCheckWithoutID("manual-create-karl", []pingdom.Step{*step})
	id := getCheckIdByName("manual-create-karl", suite)
	if id == 0 {
		created, resp, err := suite.apiClient.TMSChecksAPI.AddCheck(context.Background()).CheckWithoutID(checkWithoutID).Execute()
		suite.NoError(err)
		suite.T().Logf("Response from `TMSChecksAPI.AddCheck` Create: %v\n", resp)
		suite.Assert().NotNil(created.Name)
		suite.T().Logf("Created check with name `manual-create-karl` and id %v\n", created)
	} else {
		suite.T().Logf("Check with name `manual-create-karl` and id %d already exists. Updating...", id)
		_, r, err := suite.apiClient.TMSChecksAPI.ModifyCheck(context.Background(), id).CheckWithoutIDPUT(*checkWithoutID.AsPut()).Execute()
		suite.NoError(err)
		suite.T().Logf("Response from `TMSChecksAPI.ModifyCheck` Updated: %v\n", r)
		//suite.Assert().Equal(updated.Name, "manual-create-karl")
	}
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}

func getCheckIdByName(name string, suite *APITestSuite) int64 {
	checks, _, err := suite.apiClient.TMSChecksAPI.GetAllChecks(context.Background()).Execute()
	if err != nil {
		fmt.Printf("Error listing checks: %v", err)
	}
	for _, check := range checks.Checks {
		if *check.Name == name {
			return check.GetId()
		}
	}
	return 0
}
