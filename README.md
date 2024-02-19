# Pingdom Golang Client

This client is generated with [openapitools/openapi-generator](https://github.com/OpenAPITools/openapi-generator). It is based on the [Pingdom OpenAPI specification](https://docs.pingdom.com/api/).

## Usage

Check out the examples in the [test directory](./pkg/pingdom/test/openapi_test.go).
Further documentation and examples for each endpoint can found at [generated documentation](./pkg/pingdom/openapi/README.md).
However that is not 100% accurate due to wrong openapi specification is not matching the actual AP in some casesI.

```go
package main

import (
	"context"
	"fmt"
	"os"
	pingdom "github.com/karlderkaefer/pingdom-golang-client/pkg/pingdom/openapi"
)

func main() {
	extendedTags := true // bool | Specifies whether to return an extended tags representation in the response (with type and count). (optional)
	tags := []string{"Inner_example"} // []string | Tag list separated by commas. As an example \"nginx,apache\" would filter out all responses except those tagged nginx or apache (optional)
	type_ := "type__example" // string | Filter to return only checks of a given type (a TMS `script` or a WPM `recording`). If not provided, all checks are returned. (optional)
	limit := "limit_example" // string | Limit of returned checks (optional) (default to "1000")
	offset := "offset_example" // string | Offset of returned checks (optional) (default to "0")

	configuration := pingdom.NewConfiguration()
	configuration.SetApiToken(os.Getenv("PINGDOM_API_TOKEN"))
	apiClient := pingdom.NewAPIClient(configuration)
	resp, r, err := apiClient.TMSChecksAPI.GetAllChecks(context.Background()).ExtendedTags(extendedTags).Tags(tags).Type_(type_).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TMSChecksAPI.GetAllChecks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllChecks`: ChecksAll
	fmt.Fprintf(os.Stdout, "Response from `TMSChecksAPI.GetAllChecks`: %v\n", resp)
	for _, check := range resp.Checks {
		fmt.Fprintf(os.Stdout, "Check name: %v\n", *check.Name)
	}
}
```

This is a scetion that contain a lot of typo. we gonna use crewai agent to correct spelling in this one. 
