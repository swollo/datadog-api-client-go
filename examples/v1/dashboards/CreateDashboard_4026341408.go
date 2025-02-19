// Create a new dashboard with apm_issue_stream list_stream widget

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
	body := datadogV1.Dashboard{
		LayoutType: datadogV1.DASHBOARDLAYOUTTYPE_ORDERED,
		Title:      "Example-Create_a_new_dashboard_with_apm_issue_stream_list_stream_widget with list_stream widget",
		Widgets: []datadogV1.Widget{
			{
				Definition: datadogV1.WidgetDefinition{
					ListStreamWidgetDefinition: &datadogV1.ListStreamWidgetDefinition{
						Type: datadogV1.LISTSTREAMWIDGETDEFINITIONTYPE_LIST_STREAM,
						Requests: []datadogV1.ListStreamWidgetRequest{
							{
								Columns: []datadogV1.ListStreamColumn{
									{
										Width: datadogV1.LISTSTREAMCOLUMNWIDTH_AUTO,
										Field: "timestamp",
									},
								},
								Query: datadogV1.ListStreamQuery{
									DataSource:  datadogV1.LISTSTREAMSOURCE_APM_ISSUE_STREAM,
									QueryString: "",
								},
								ResponseFormat: datadogV1.LISTSTREAMRESPONSEFORMAT_EVENT_LIST,
							},
						},
					}},
			},
		},
	}
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV1.NewDashboardsApi(apiClient)
	resp, r, err := api.CreateDashboard(ctx, body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsApi.CreateDashboard`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from `DashboardsApi.CreateDashboard`:\n%s\n", responseContent)
}
