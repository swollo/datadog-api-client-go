// Create a detection rule with type 'workload_security' returns "OK" response

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	body := datadog.SecurityMonitoringRuleCreatePayload{
		Name: "Example-Create_a_detection_rule_with_type_workload_security_returns_OK_response",
		Queries: []datadog.SecurityMonitoringRuleQueryCreate{
			{
				Query:          "@test:true",
				Aggregation:    datadog.SECURITYMONITORINGRULEQUERYAGGREGATION_COUNT.Ptr(),
				GroupByFields:  &[]string{},
				DistinctFields: &[]string{},
				Metric:         datadog.PtrString(""),
			},
		},
		Filters: &[]datadog.SecurityMonitoringFilter{},
		Cases: []datadog.SecurityMonitoringRuleCaseCreate{
			{
				Name:          datadog.PtrString(""),
				Status:        datadog.SECURITYMONITORINGRULESEVERITY_INFO,
				Condition:     datadog.PtrString("a > 0"),
				Notifications: &[]string{},
			},
		},
		Options: datadog.SecurityMonitoringRuleOptions{
			EvaluationWindow:  datadog.SECURITYMONITORINGRULEEVALUATIONWINDOW_FIFTEEN_MINUTES.Ptr(),
			KeepAlive:         datadog.SECURITYMONITORINGRULEKEEPALIVE_ONE_HOUR.Ptr(),
			MaxSignalDuration: datadog.SECURITYMONITORINGRULEMAXSIGNALDURATION_ONE_DAY.Ptr(),
		},
		Message:   "Test rule",
		Tags:      &[]string{},
		IsEnabled: true,
		Type:      datadog.SECURITYMONITORINGRULETYPECREATE_WORKLOAD_SECURITY.Ptr(),
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.CloudSIEMApi.CreateSecurityMonitoringRule(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CloudSIEMApi.CreateSecurityMonitoringRule`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `CloudSIEMApi.CreateSecurityMonitoringRule`:\n%s\n", responseContent)
}
