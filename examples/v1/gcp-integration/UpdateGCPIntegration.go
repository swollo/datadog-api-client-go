// Update a GCP integration returns "OK" response

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
	body := datadogV1.GCPAccount{
		AuthProviderX509CertUrl: datadog.PtrString("https://www.googleapis.com/oauth2/v1/certs"),
		AuthUri:                 datadog.PtrString("https://accounts.google.com/o/oauth2/auth"),
		ClientEmail:             datadog.PtrString("api-dev@datadog-sandbox.iam.gserviceaccount.com"),
		ClientId:                datadog.PtrString("123456712345671234567"),
		ClientX509CertUrl:       datadog.PtrString("https://www.googleapis.com/robot/v1/metadata/x509/$CLIENT_EMAIL"),
		Errors: []string{
			"*",
		},
		HostFilters:  datadog.PtrString("key:value,filter:example"),
		PrivateKey:   datadog.PtrString("private_key"),
		PrivateKeyId: datadog.PtrString("123456789abcdefghi123456789abcdefghijklm"),
		ProjectId:    datadog.PtrString("datadog-apitest"),
		TokenUri:     datadog.PtrString("https://accounts.google.com/o/oauth2/token"),
		Type:         datadog.PtrString("service_account"),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewGCPIntegrationApi(apiClient)
	resp, r, err := api.UpdateGCPIntegration(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GCPIntegrationApi.UpdateGCPIntegration`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `GCPIntegrationApi.UpdateGCPIntegration`:\n%s\n", responseContent)
}
