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

// HelmActionsAPIService HelmActionsAPI service
type HelmActionsAPIService service

type ApiDeployHelmRequest struct {
	ctx               context.Context
	ApiService        *HelmActionsAPIService
	helmId            string
	forceEvent        *HelmForceEvent
	helmDeployRequest *HelmDeployRequest
}

// When filled, it indicates the target event to be deployed.   If the concerned helm hasn&#39;t the target event provided, the helm won&#39;t be deployed.
func (r ApiDeployHelmRequest) ForceEvent(forceEvent HelmForceEvent) ApiDeployHelmRequest {
	r.forceEvent = &forceEvent
	return r
}

func (r ApiDeployHelmRequest) HelmDeployRequest(helmDeployRequest HelmDeployRequest) ApiDeployHelmRequest {
	r.helmDeployRequest = &helmDeployRequest
	return r
}

func (r ApiDeployHelmRequest) Execute() (*Status, *http.Response, error) {
	return r.ApiService.DeployHelmExecute(r)
}

/*
DeployHelm Deploy helm

You must provide a git commit id or an image tag depending on the source location of your code (git vs image repository).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param helmId Helm ID
	@return ApiDeployHelmRequest
*/
func (a *HelmActionsAPIService) DeployHelm(ctx context.Context, helmId string) ApiDeployHelmRequest {
	return ApiDeployHelmRequest{
		ApiService: a,
		ctx:        ctx,
		helmId:     helmId,
	}
}

// Execute executes the request
//
//	@return Status
func (a *HelmActionsAPIService) DeployHelmExecute(r ApiDeployHelmRequest) (*Status, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Status
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HelmActionsAPIService.DeployHelm")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/helm/{helmId}/deploy"
	localVarPath = strings.Replace(localVarPath, "{"+"helmId"+"}", url.PathEscape(parameterValueToString(r.helmId, "helmId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.forceEvent != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "forceEvent", r.forceEvent, "")
	}
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
	localVarPostBody = r.helmDeployRequest
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

type ApiRestartHelmRequest struct {
	ctx        context.Context
	ApiService *HelmActionsAPIService
	helmId     string
	forceEvent *HelmForceEvent
}

// When filled, it indicates the target event to be deployed.   If the concerned helm hasn&#39;t the target event provided, the helm won&#39;t be deployed.
func (r ApiRestartHelmRequest) ForceEvent(forceEvent HelmForceEvent) ApiRestartHelmRequest {
	r.forceEvent = &forceEvent
	return r
}

func (r ApiRestartHelmRequest) Execute() (*Status, *http.Response, error) {
	return r.ApiService.RestartHelmExecute(r)
}

/*
RestartHelm Deprecated - Restart helm

**Deprecated** - Please use the "Redeploy helm" endpoint now

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param helmId Helm ID
	@return ApiRestartHelmRequest
*/
func (a *HelmActionsAPIService) RestartHelm(ctx context.Context, helmId string) ApiRestartHelmRequest {
	return ApiRestartHelmRequest{
		ApiService: a,
		ctx:        ctx,
		helmId:     helmId,
	}
}

// Execute executes the request
//
//	@return Status
func (a *HelmActionsAPIService) RestartHelmExecute(r ApiRestartHelmRequest) (*Status, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Status
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HelmActionsAPIService.RestartHelm")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/helm/{helmId}/restart"
	localVarPath = strings.Replace(localVarPath, "{"+"helmId"+"}", url.PathEscape(parameterValueToString(r.helmId, "helmId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.forceEvent != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "forceEvent", r.forceEvent, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiStopHelmRequest struct {
	ctx        context.Context
	ApiService *HelmActionsAPIService
	helmId     string
}

func (r ApiStopHelmRequest) Execute() (*Status, *http.Response, error) {
	return r.ApiService.StopHelmExecute(r)
}

/*
StopHelm Stop helm

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param helmId Helm ID
	@return ApiStopHelmRequest
*/
func (a *HelmActionsAPIService) StopHelm(ctx context.Context, helmId string) ApiStopHelmRequest {
	return ApiStopHelmRequest{
		ApiService: a,
		ctx:        ctx,
		helmId:     helmId,
	}
}

// Execute executes the request
//
//	@return Status
func (a *HelmActionsAPIService) StopHelmExecute(r ApiStopHelmRequest) (*Status, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Status
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HelmActionsAPIService.StopHelm")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/helm/{helmId}/stop"
	localVarPath = strings.Replace(localVarPath, "{"+"helmId"+"}", url.PathEscape(parameterValueToString(r.helmId, "helmId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
