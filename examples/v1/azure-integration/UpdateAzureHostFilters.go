// Update Azure integration host filters returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
)

func main() {
	body := datadogV1.AzureAccount{
		Automute:     datadog.PtrBool(true),
		ClientId:     datadog.PtrString("testc7f6-1234-5678-9101-3fcbf464test"),
		ClientSecret: datadog.PtrString("testingx./Sw*g/Y33t..R1cH+hScMDt"),
		Errors: []string{
			"*",
		},
		HostFilters:   datadog.PtrString("key:value,filter:example"),
		NewClientId:   datadog.PtrString("new1c7f6-1234-5678-9101-3fcbf464test"),
		NewTenantName: datadog.PtrString("new1c44-1234-5678-9101-cc00736ftest"),
		TenantName:    datadog.PtrString("testc44-1234-5678-9101-cc00736ftest"),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewAzureIntegrationApi(apiClient)
	resp, r, err := api.UpdateAzureHostFilters(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureIntegrationApi.UpdateAzureHostFilters`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `AzureIntegrationApi.UpdateAzureHostFilters`:\n%s\n", responseContent)
}
