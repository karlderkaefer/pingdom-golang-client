package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/karlderkaefer/pingdom-golang-client/pkg/pingdom/client"
	"github.com/karlderkaefer/pingdom-golang-client/pkg/pingdom/client/checks"
	"github.com/karlderkaefer/pingdom-golang-client/pkg/pingdom/client/ptr"
	"github.com/karlderkaefer/pingdom-golang-client/pkg/pingdom/client/tmschecks"
)

func TestGetChecks(t *testing.T) {
	// Set the authorization token in the HTTP client headers
	token := os.Getenv("PINGDOM_API_TOKEN")

	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken(token)
	if err != nil {
		t.Fatalf("Error creating bearer token provider: %v", err)
	}
	client, err := checks.NewClientWithResponses(
		client.DefaultBaseURL,
		checks.WithRequestEditorFn(bearerTokenProvider.Intercept),
	)
	if err != nil {
		t.Fatalf("Error creating Pingdom client: %v", err)
	}
	// Send the request to create the check
	checks, err := client.GetChecksWithResponse(context.Background(), &checks.GetChecksParams{})
	if err != nil {
		t.Fatalf("Error creating check: %v", err)
	}

	// Print the created check details
	for _, check := range *checks.JSON200.Checks {
		t.Logf("Check created: %+v\n", *check.Name)
	}
}

func TestUpsertTransactionCheck(t *testing.T) {
	// Set the authorization token in the HTTP client headers
	token := os.Getenv("PINGDOM_API_TOKEN")

	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken(token)
	if err != nil {
		t.Fatalf("Error creating bearer token provider: %v", err)
	}
	client, err := tmschecks.NewClientWithResponses(
		client.DefaultBaseURL,
		tmschecks.WithRequestEditorFn(bearerTokenProvider.Intercept),
	)
	if err != nil {
		t.Fatalf("Error creating Pingdom client: %v", err)
	}

	// Send the request to create the check
	check := tmschecks.CheckWithoutID{
		Name: "manual-create-karl",
		Steps: []tmschecks.Step{
			{
				Function: ptr.String("go_to"),
				StepArg: &tmschecks.StepArg{
					URL: ptr.String("https://www.google.com"),
				},
			},
		},
	}
	created, err := client.AddCheckWithResponse(context.Background(), check)
	if err != nil {
		t.Fatalf("Error creating check: %v", err)
	}
	fmt.Printf("Created Check: %+v\n", created.JSON200)
	//assert.Equal(t, *created.JSON200.Name, "manual-create-karl")
}
