/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

package test

import (
	"context"
	"testing"

	"github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/DataDog/datadog-api-client-go/tests"
)

func TestIncidentSelectionsLifecycle(t *testing.T) {
	ctx, finish := WithRecorder(WithTestAuth(context.Background()), t)
	defer finish()
	assert := tests.Assert(ctx, t)
	client := Client(ctx)

	setIncidentsUnstableOperations(client, true)
	defer setIncidentsUnstableOperations(client, false)

	incident := createTestIncident(t, client, ctx)
	defer deleteTestIncident(t, client, ctx, incident.GetId())
	field := createIncidentConfigField(t, client, ctx)
	defer deleteIncidentConfigField(t, client, ctx, field.GetId())
	choice1 := createIncidentConfigFieldChoice(t, client, ctx, field.GetId())
	defer deleteIncidentConfigFieldChoice(t, client, ctx, field.GetId(), choice1.GetId())
	choice2 := createIncidentConfigFieldChoice(t, client, ctx, field.GetId())
	defer deleteIncidentConfigFieldChoice(t, client, ctx, field.GetId(), choice2.GetId())

	testIncidentSelectionData := datadog.NewIncidentSelectionWithDefaults()
	testIncidentSelectionData.SetType("selections")
	testIncidentSelectionData.SetAttributes(*datadog.NewIncidentSelectionAttributesWithDefaults())
	testIncidentSelectionData.SetRelationships(*datadog.NewIncidentSelectionRelationshipsWithDefaults())
	testIncidentSelectionData.Attributes.SetObjectId(incident.GetId())
	testIncidentSelectionData.Relationships.SetField(*datadog.NewIncidentSelectionRelationshipsFieldWithDefaults())
	testIncidentSelectionData.Relationships.Field.SetData(field)
	testIncidentSelectionData.Relationships.SetChoice(*datadog.NewIncidentSelectionRelationshipsChoiceWithDefaults())
	testIncidentSelectionData.Relationships.Choice.SetData(choice1)

	// Create IncidentSelection
	incidentSelectionRsp, httpresp, err := client.IncidentsApi.CreateIncidentSelection(ctx, incident.GetId()).
		Body(*datadog.NewIncidentSelectionPayload(*testIncidentSelectionData)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error creating IncidentSelection %v: Response %s: %v", testIncidentSelectionData, err.Error(), bStr)
	}
	incidentSelection := incidentSelectionRsp.GetData()
	incidentSelectionAttrs := incidentSelection.GetAttributes()
	incidentSelectionRelationships := incidentSelection.GetRelationships()
	assert.Equal(httpresp.StatusCode, 201)

	// Lets ensure that we always delete the incidentSelection if we successfully created it, even if the update and
	// get requests fail
	defer func(test_client *datadog.APIClient, incidentSelection datadog.IncidentSelection) {
		//Delete IncidentSelection
		httpresp, err := test_client.IncidentsApi.DeleteIncidentSelection(ctx, incident.GetId(), incidentSelection.GetId()).Execute()
		if err != nil {
			bStr := err.(datadog.GenericOpenAPIError).Model()
			t.Fatalf("Error deleting IncidentSelection %v: Response %s: %v", incidentSelection, err.Error(), bStr)
		}
		assert.Equal(httpresp.StatusCode, 204)
	}(client, incidentSelection)

	assert.Equal(incidentSelection.GetType(), testIncidentSelectionData.GetType())
	assert.Equal(incidentSelectionAttrs.GetObjectId(), incident.GetId())
	assert.Equal(incidentSelectionRelationships.Choice.Data.GetId(), choice1.GetId())
	assert.Equal(incidentSelectionRelationships.Field.Data.GetId(), field.GetId())

	// Edit IncidentSelection
	incidentSelection.Relationships.Choice.SetData(choice2)
	incidentSelectionUpdatedRsp, httpresp, err := client.IncidentsApi.
		PatchIncidentSelection(ctx, incident.GetId(), incidentSelection.GetId()).
		Body(*datadog.NewIncidentSelectionPayload(incidentSelection)).Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error updating IncidentSelection %v: Response %s: %v", incidentSelection, err.Error(), bStr)
	}
	incidentSelectionUpdated := incidentSelectionUpdatedRsp.GetData()
	incidentSelectionAttrsUpdated := incidentSelectionUpdated.GetAttributes()
	incidentSelectionRelationshipsUpdated := incidentSelectionUpdated.GetRelationships()
	assert.Equal(httpresp.StatusCode, 200)
	assert.Equal(incidentSelectionAttrsUpdated.GetObjectId(), incident.GetId())
	assert.Equal(incidentSelectionRelationshipsUpdated.Choice.Data.GetId(), choice2.GetId())
	assert.Equal(incidentSelectionRelationshipsUpdated.Field.Data.GetId(), field.GetId())

	// Get IncidentSelections
	incidentSelectionsGetResp, httpresp, err := client.IncidentsApi.GetIncidentSelections(ctx, incident.GetId()).
		Execute()
	if err != nil {
		bStr := err.(datadog.GenericOpenAPIError).Model()
		t.Fatalf("Error getting IncidentSelections %v: Response %s: %v", incidentSelection, err.Error(), bStr)
	}
	assert.Equal(httpresp.StatusCode, 200)
	assert.True(len(incidentSelectionsGetResp.GetData()) >= 1)

}
