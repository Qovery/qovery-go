/*
[BETA] Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is in Beta and still in progress. Some endpoints are not available yet.

API version: 1.0.0
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

// EnvironmentDeploymentRuleApiService EnvironmentDeploymentRuleApi service
type EnvironmentDeploymentRuleApiService service

type ApiEditEnvironmentDeploymentRuleRequest struct {
	ctx                                  context.Context
	ApiService                           *EnvironmentDeploymentRuleApiService
	environmentId                        string
	deploymentRuleId                     string
	environmentDeploymentRuleEditRequest *EnvironmentDeploymentRuleEditRequest
}

func (r ApiEditEnvironmentDeploymentRuleRequest) EnvironmentDeploymentRuleEditRequest(environmentDeploymentRuleEditRequest EnvironmentDeploymentRuleEditRequest) ApiEditEnvironmentDeploymentRuleRequest {
	r.environmentDeploymentRuleEditRequest = &environmentDeploymentRuleEditRequest
	return r
}

func (r ApiEditEnvironmentDeploymentRuleRequest) Execute() (*EnvironmentDeploymentRuleResponse, *http.Response, error) {
	return r.ApiService.EditEnvironmentDeploymentRuleExecute(r)
}

/*
EditEnvironmentDeploymentRule Edit an environment deployment rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @param deploymentRuleId Deployment Rule ID
 @return ApiEditEnvironmentDeploymentRuleRequest
*/
func (a *EnvironmentDeploymentRuleApiService) EditEnvironmentDeploymentRule(ctx context.Context, environmentId string, deploymentRuleId string) ApiEditEnvironmentDeploymentRuleRequest {
	return ApiEditEnvironmentDeploymentRuleRequest{
		ApiService:       a,
		ctx:              ctx,
		environmentId:    environmentId,
		deploymentRuleId: deploymentRuleId,
	}
}

// Execute executes the request
//  @return EnvironmentDeploymentRuleResponse
func (a *EnvironmentDeploymentRuleApiService) EditEnvironmentDeploymentRuleExecute(r ApiEditEnvironmentDeploymentRuleRequest) (*EnvironmentDeploymentRuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EnvironmentDeploymentRuleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentDeploymentRuleApiService.EditEnvironmentDeploymentRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/deploymentRule/{deploymentRuleId}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"deploymentRuleId"+"}", url.PathEscape(parameterToString(r.deploymentRuleId, "")), -1)

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
	localVarPostBody = r.environmentDeploymentRuleEditRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetEnvironmentDeploymentRuleRequest struct {
	ctx           context.Context
	ApiService    *EnvironmentDeploymentRuleApiService
	environmentId string
}

func (r ApiGetEnvironmentDeploymentRuleRequest) Execute() (*EnvironmentDeploymentRuleResponse, *http.Response, error) {
	return r.ApiService.GetEnvironmentDeploymentRuleExecute(r)
}

/*
GetEnvironmentDeploymentRule Get environment deployment rule

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @return ApiGetEnvironmentDeploymentRuleRequest
*/
func (a *EnvironmentDeploymentRuleApiService) GetEnvironmentDeploymentRule(ctx context.Context, environmentId string) ApiGetEnvironmentDeploymentRuleRequest {
	return ApiGetEnvironmentDeploymentRuleRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return EnvironmentDeploymentRuleResponse
func (a *EnvironmentDeploymentRuleApiService) GetEnvironmentDeploymentRuleExecute(r ApiGetEnvironmentDeploymentRuleRequest) (*EnvironmentDeploymentRuleResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EnvironmentDeploymentRuleResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvironmentDeploymentRuleApiService.GetEnvironmentDeploymentRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/deploymentRule"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterToString(r.environmentId, "")), -1)

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
