package example

import (
	"context"
	"fmt"
	"os"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/karlderkaefer/pingdom-golang-client/pkg/pingdom/client"
	"github.com/karlderkaefer/pingdom-golang-client/pkg/pingdom/client/checks"
)

func main() {
	// Set the authorization token in the HTTP client headers
	token := os.Getenv("PINGDOM_API_TOKEN")

	bearerTokenProvider, err := securityprovider.NewSecurityProviderBearerToken(token)
	if err != nil {
		panic(err)
	}
	client, err := checks.NewClientWithResponses(
		client.DefaultBaseURL, 
		checks.WithRequestEditorFn(bearerTokenProvider.Intercept),
	)
	if err != nil {
		fmt.Println("Error creating Pingdom client:", err)
		return
	}
	// Send the request to create the check
	checks, err := client.GetChecksWithResponse(context.Background(), &checks.GetChecksParams{})
	if err != nil {
		fmt.Println("Error creating check:", err)
		return
	}

	// Print the created check details
	for _, check := range *checks.JSON200.Checks {
		fmt.Printf("Check created: %+v\n", *check.Name)
	}
}
