// Create a new dashboard with apm dependency stats widget

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
		Title: "Example-Create_a_new_dashboard_with_apm_dependency_stats_widget",
		Widgets: []datadogV1.Widget{
			{
				Definition: datadogV1.WidgetDefinition{
					TableWidgetDefinition: &datadogV1.TableWidgetDefinition{
						Title:      datadog.PtrString(""),
						TitleSize:  datadog.PtrString("16"),
						TitleAlign: datadogV1.WIDGETTEXTALIGN_LEFT.Ptr(),
						Type:       datadogV1.TABLEWIDGETDEFINITIONTYPE_QUERY_TABLE,
						Requests: []datadogV1.TableWidgetRequest{
							{
								ResponseFormat: datadogV1.FORMULAANDFUNCTIONRESPONSEFORMAT_SCALAR.Ptr(),
								Queries: []datadogV1.FormulaAndFunctionQueryDefinition{
									datadogV1.FormulaAndFunctionQueryDefinition{
										FormulaAndFunctionApmDependencyStatsQueryDefinition: &datadogV1.FormulaAndFunctionApmDependencyStatsQueryDefinition{
											PrimaryTagValue: datadog.PtrString("edge-eu1.prod.dog"),
											Stat:            datadogV1.FORMULAANDFUNCTIONAPMDEPENDENCYSTATNAME_AVG_DURATION,
											ResourceName:    "DELETE FROM monitor_history.monitor_state_change_history WHERE org_id = ? AND monitor_id IN ? AND group = ?",
											Name:            "query1",
											Service:         "cassandra",
											DataSource:      datadogV1.FORMULAANDFUNCTIONAPMDEPENDENCYSTATSDATASOURCE_APM_DEPENDENCY_STATS,
											Env:             "ci",
											PrimaryTagName:  datadog.PtrString("datacenter"),
											OperationName:   "cassandra.query",
										}},
								},
							},
						},
					}},
				Layout: &datadogV1.WidgetLayout{
					X:      0,
					Y:      0,
					Width:  4,
					Height: 4,
				},
			},
		},
		LayoutType: datadogV1.DASHBOARDLAYOUTTYPE_ORDERED,
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
