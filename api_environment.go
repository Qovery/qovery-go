/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// EnvironmentAPIService EnvironmentAPI service
type EnvironmentAPIService service

type ApiDeployAllApplicationsRequest struct {
	ctx              context.Context
	ApiService       *EnvironmentAPIService
	environmentId    string
	deployAllRequest *DeployAllRequest
}

func (r ApiDeployAllApplicationsRequest) DeployAllRequest(deployAllRequest DeployAllRequest) ApiDeployAllApplicationsRequest {
	r.deployAllRequest = &deployAllRequest
	return r
}

func (r ApiDeployAllApplicationsRequest) Execute() (*Status, *http.Response, error) {
	return r.ApiService.DeployAllApplicationsExecute(r)
}

/*
DeployAllApplications Deploy applications

Start a deployment of the environment. Any of the services within the chosen environment based on the following rule: a service is deployed only if a new version is specified in the payload for that application/container or if there was a change in its configuration that needs to be applied (vCPU,RAM etc..)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentId Environment ID
	@return ApiDeployAllApplicationsRequest
*/
func (a *EnvironmentAPIService) DeployAllApplications(ctx context.Context, environmentId string) ApiDeployAllApplicationsRequest {
	return ApiDeployAllApplicationsRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//
//	@return Status
func (a *EnvironmentAPIService) DeployAllApplicationsExecute(r ApiDeployAllApplicationsRequest) (*Status, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Status
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentAPIService.DeployAllApplications")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/application/deploy"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterValueToString(r.environmentId, "environmentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.deployAllRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
