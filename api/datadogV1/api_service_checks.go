// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"bytes"
	_context "context"
	_io "io"
	_nethttp "net/http"
	_neturl "net/url"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceChecksApi service type
type ServiceChecksApi datadog.Service

type apiSubmitServiceCheckRequest struct {
	ctx  _context.Context
	body *[]ServiceCheck
}

func (a *ServiceChecksApi) buildSubmitServiceCheckRequest(ctx _context.Context, body []ServiceCheck) (apiSubmitServiceCheckRequest, error) {
	req := apiSubmitServiceCheckRequest{
		ctx:  ctx,
		body: &body,
	}
	return req, nil
}

// SubmitServiceCheck Submit a Service Check.
// Submit a list of Service Checks.
//
// **Notes**:
// - A valid API key is required.
// - Service checks can be submitted up to 10 minutes in the past.
func (a *ServiceChecksApi) SubmitServiceCheck(ctx _context.Context, body []ServiceCheck) (IntakePayloadAccepted, *_nethttp.Response, error) {
	req, err := a.buildSubmitServiceCheckRequest(ctx, body)
	if err != nil {
		var localVarReturnValue IntakePayloadAccepted
		return localVarReturnValue, nil, err
	}

	return a.submitServiceCheckExecute(req)
}

// submitServiceCheckExecute executes the request.
func (a *ServiceChecksApi) submitServiceCheckExecute(r apiSubmitServiceCheckRequest) (IntakePayloadAccepted, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		localVarReturnValue IntakePayloadAccepted
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "v1.ServiceChecksApi.SubmitServiceCheck")
	if err != nil {
		return localVarReturnValue, nil, datadog.GenericOpenAPIError{ErrorMessage: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/check_run"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, datadog.ReportError("body is required and must be specified")
	}
	localVarHeaderParams["Content-Type"] = "application/json"
	localVarHeaderParams["Accept"] = "application/json"

	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(datadog.ContextAPIKeys).(map[string]datadog.APIKey); ok {
			if apiKey, ok := auth["apiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["DD-API-KEY"] = key
			}
		}
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, nil)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 || localVarHTTPResponse.StatusCode == 403 || localVarHTTPResponse.StatusCode == 408 || localVarHTTPResponse.StatusCode == 413 || localVarHTTPResponse.StatusCode == 429 {
			var v APIErrorResponse
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.ErrorModel = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := datadog.GenericOpenAPIError{
			ErrorBody:    localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

// NewServiceChecksApi Returns NewServiceChecksApi.
func NewServiceChecksApi(client *datadog.APIClient) *ServiceChecksApi {
	return &ServiceChecksApi{
		Client: client,
	}
}
