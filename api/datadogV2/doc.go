// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

// List of APIs:
//  - [AuditApi.ListAuditLogs]
//  - [AuditApi.SearchAuditLogs]
//  - [AuthNMappingsApi.CreateAuthNMapping]
//  - [AuthNMappingsApi.DeleteAuthNMapping]
//  - [AuthNMappingsApi.GetAuthNMapping]
//  - [AuthNMappingsApi.ListAuthNMappings]
//  - [AuthNMappingsApi.UpdateAuthNMapping]
//  - [CIVisibilityPipelinesApi.AggregateCIAppPipelineEvents]
//  - [CIVisibilityPipelinesApi.ListCIAppPipelineEvents]
//  - [CIVisibilityPipelinesApi.SearchCIAppPipelineEvents]
//  - [CIVisibilityTestsApi.AggregateCIAppTestEvents]
//  - [CIVisibilityTestsApi.ListCIAppTestEvents]
//  - [CIVisibilityTestsApi.SearchCIAppTestEvents]
//  - [CloudWorkloadSecurityApi.CreateCloudWorkloadSecurityAgentRule]
//  - [CloudWorkloadSecurityApi.DeleteCloudWorkloadSecurityAgentRule]
//  - [CloudWorkloadSecurityApi.DownloadCloudWorkloadPolicyFile]
//  - [CloudWorkloadSecurityApi.GetCloudWorkloadSecurityAgentRule]
//  - [CloudWorkloadSecurityApi.ListCloudWorkloadSecurityAgentRules]
//  - [CloudWorkloadSecurityApi.UpdateCloudWorkloadSecurityAgentRule]
//  - [ConfluentCloudApi.CreateConfluentAccount]
//  - [ConfluentCloudApi.CreateConfluentResource]
//  - [ConfluentCloudApi.DeleteConfluentAccount]
//  - [ConfluentCloudApi.DeleteConfluentResource]
//  - [ConfluentCloudApi.GetConfluentAccount]
//  - [ConfluentCloudApi.GetConfluentResource]
//  - [ConfluentCloudApi.ListConfluentAccount]
//  - [ConfluentCloudApi.ListConfluentResource]
//  - [ConfluentCloudApi.UpdateConfluentAccount]
//  - [ConfluentCloudApi.UpdateConfluentResource]
//  - [DashboardListsApi.CreateDashboardListItems]
//  - [DashboardListsApi.DeleteDashboardListItems]
//  - [DashboardListsApi.GetDashboardListItems]
//  - [DashboardListsApi.UpdateDashboardListItems]
//  - [EventsApi.ListEvents]
//  - [EventsApi.SearchEvents]
//  - [IncidentServicesApi.CreateIncidentService]
//  - [IncidentServicesApi.DeleteIncidentService]
//  - [IncidentServicesApi.GetIncidentService]
//  - [IncidentServicesApi.ListIncidentServices]
//  - [IncidentServicesApi.UpdateIncidentService]
//  - [IncidentTeamsApi.CreateIncidentTeam]
//  - [IncidentTeamsApi.DeleteIncidentTeam]
//  - [IncidentTeamsApi.GetIncidentTeam]
//  - [IncidentTeamsApi.ListIncidentTeams]
//  - [IncidentTeamsApi.UpdateIncidentTeam]
//  - [IncidentsApi.CreateIncident]
//  - [IncidentsApi.DeleteIncident]
//  - [IncidentsApi.GetIncident]
//  - [IncidentsApi.ListIncidentAttachments]
//  - [IncidentsApi.ListIncidents]
//  - [IncidentsApi.SearchIncidents]
//  - [IncidentsApi.UpdateIncident]
//  - [IncidentsApi.UpdateIncidentAttachments]
//  - [KeyManagementApi.CreateAPIKey]
//  - [KeyManagementApi.CreateCurrentUserApplicationKey]
//  - [KeyManagementApi.DeleteAPIKey]
//  - [KeyManagementApi.DeleteApplicationKey]
//  - [KeyManagementApi.DeleteCurrentUserApplicationKey]
//  - [KeyManagementApi.GetAPIKey]
//  - [KeyManagementApi.GetApplicationKey]
//  - [KeyManagementApi.GetCurrentUserApplicationKey]
//  - [KeyManagementApi.ListAPIKeys]
//  - [KeyManagementApi.ListApplicationKeys]
//  - [KeyManagementApi.ListCurrentUserApplicationKeys]
//  - [KeyManagementApi.UpdateAPIKey]
//  - [KeyManagementApi.UpdateApplicationKey]
//  - [KeyManagementApi.UpdateCurrentUserApplicationKey]
//  - [LogsApi.AggregateLogs]
//  - [LogsApi.ListLogs]
//  - [LogsApi.ListLogsGet]
//  - [LogsApi.SubmitLog]
//  - [LogsArchivesApi.AddReadRoleToArchive]
//  - [LogsArchivesApi.CreateLogsArchive]
//  - [LogsArchivesApi.DeleteLogsArchive]
//  - [LogsArchivesApi.GetLogsArchive]
//  - [LogsArchivesApi.GetLogsArchiveOrder]
//  - [LogsArchivesApi.ListArchiveReadRoles]
//  - [LogsArchivesApi.ListLogsArchives]
//  - [LogsArchivesApi.RemoveRoleFromArchive]
//  - [LogsArchivesApi.UpdateLogsArchive]
//  - [LogsArchivesApi.UpdateLogsArchiveOrder]
//  - [LogsMetricsApi.CreateLogsMetric]
//  - [LogsMetricsApi.DeleteLogsMetric]
//  - [LogsMetricsApi.GetLogsMetric]
//  - [LogsMetricsApi.ListLogsMetrics]
//  - [LogsMetricsApi.UpdateLogsMetric]
//  - [MetricsApi.CreateBulkTagsMetricsConfiguration]
//  - [MetricsApi.CreateTagConfiguration]
//  - [MetricsApi.DeleteBulkTagsMetricsConfiguration]
//  - [MetricsApi.DeleteTagConfiguration]
//  - [MetricsApi.EstimateMetricsOutputSeries]
//  - [MetricsApi.ListActiveMetricConfigurations]
//  - [MetricsApi.ListTagConfigurationByName]
//  - [MetricsApi.ListTagConfigurations]
//  - [MetricsApi.ListTagsByMetricName]
//  - [MetricsApi.ListVolumesByMetricName]
//  - [MetricsApi.SubmitMetrics]
//  - [MetricsApi.UpdateTagConfiguration]
//  - [OpsgenieIntegrationApi.CreateOpsgenieService]
//  - [OpsgenieIntegrationApi.DeleteOpsgenieService]
//  - [OpsgenieIntegrationApi.GetOpsgenieService]
//  - [OpsgenieIntegrationApi.ListOpsgenieServices]
//  - [OpsgenieIntegrationApi.UpdateOpsgenieService]
//  - [OrganizationsApi.UploadIdPMetadata]
//  - [ProcessesApi.ListProcesses]
//  - [RUMApi.AggregateRUMEvents]
//  - [RUMApi.CreateRUMApplication]
//  - [RUMApi.DeleteRUMApplication]
//  - [RUMApi.GetRUMApplication]
//  - [RUMApi.GetRUMApplications]
//  - [RUMApi.ListRUMEvents]
//  - [RUMApi.SearchRUMEvents]
//  - [RUMApi.UpdateRUMApplication]
//  - [RolesApi.AddPermissionToRole]
//  - [RolesApi.AddUserToRole]
//  - [RolesApi.CloneRole]
//  - [RolesApi.CreateRole]
//  - [RolesApi.DeleteRole]
//  - [RolesApi.GetRole]
//  - [RolesApi.ListPermissions]
//  - [RolesApi.ListRolePermissions]
//  - [RolesApi.ListRoleUsers]
//  - [RolesApi.ListRoles]
//  - [RolesApi.RemovePermissionFromRole]
//  - [RolesApi.RemoveUserFromRole]
//  - [RolesApi.UpdateRole]
//  - [SecurityMonitoringApi.CreateSecurityFilter]
//  - [SecurityMonitoringApi.CreateSecurityMonitoringRule]
//  - [SecurityMonitoringApi.DeleteSecurityFilter]
//  - [SecurityMonitoringApi.DeleteSecurityMonitoringRule]
//  - [SecurityMonitoringApi.EditSecurityMonitoringSignalAssignee]
//  - [SecurityMonitoringApi.EditSecurityMonitoringSignalIncidents]
//  - [SecurityMonitoringApi.EditSecurityMonitoringSignalState]
//  - [SecurityMonitoringApi.GetSecurityFilter]
//  - [SecurityMonitoringApi.GetSecurityMonitoringRule]
//  - [SecurityMonitoringApi.GetSecurityMonitoringSignal]
//  - [SecurityMonitoringApi.ListSecurityFilters]
//  - [SecurityMonitoringApi.ListSecurityMonitoringRules]
//  - [SecurityMonitoringApi.ListSecurityMonitoringSignals]
//  - [SecurityMonitoringApi.SearchSecurityMonitoringSignals]
//  - [SecurityMonitoringApi.UpdateSecurityFilter]
//  - [SecurityMonitoringApi.UpdateSecurityMonitoringRule]
//  - [ServiceAccountsApi.CreateServiceAccountApplicationKey]
//  - [ServiceAccountsApi.DeleteServiceAccountApplicationKey]
//  - [ServiceAccountsApi.GetServiceAccountApplicationKey]
//  - [ServiceAccountsApi.ListServiceAccountApplicationKeys]
//  - [ServiceAccountsApi.UpdateServiceAccountApplicationKey]
//  - [ServiceDefinitionApi.CreateOrUpdateServiceDefinitions]
//  - [ServiceDefinitionApi.DeleteServiceDefinition]
//  - [ServiceDefinitionApi.GetServiceDefinition]
//  - [ServiceDefinitionApi.ListServiceDefinitions]
//  - [UsageMeteringApi.GetCostByOrg]
//  - [UsageMeteringApi.GetEstimatedCostByOrg]
//  - [UsageMeteringApi.GetHistoricalCostByOrg]
//  - [UsageMeteringApi.GetHourlyUsage]
//  - [UsageMeteringApi.GetUsageApplicationSecurityMonitoring]
//  - [UsageMeteringApi.GetUsageLambdaTracedInvocations]
//  - [UsageMeteringApi.GetUsageObservabilityPipelines]
//  - [UsersApi.CreateServiceAccount]
//  - [UsersApi.CreateUser]
//  - [UsersApi.DisableUser]
//  - [UsersApi.GetInvitation]
//  - [UsersApi.GetUser]
//  - [UsersApi.ListUserOrganizations]
//  - [UsersApi.ListUserPermissions]
//  - [UsersApi.ListUsers]
//  - [UsersApi.SendInvitations]
//  - [UsersApi.UpdateUser]
package datadogV2
