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
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// JobDeploymentHistoryApiService JobDeploymentHistoryApi service
type JobDeploymentHistoryApiService service

type ApiListJobDeploymentHistoryRequest struct {
	ctx        context.Context
	ApiService *JobDeploymentHistoryApiService
	jobId      string
}

func (r ApiListJobDeploymentHistoryRequest) Execute() (*ListJobDeploymentHistory200Response, *http.Response, error) {
	return r.ApiService.ListJobDeploymentHistoryExecute(r)
}

/*
ListJobDeploymentHistory List job deployments

Returns the 20 last job deployments

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param jobId Job ID
 @return ApiListJobDeploymentHistoryRequest
*/
func (a *JobDeploymentHistoryApiService) ListJobDeploymentHistory(ctx context.Context, jobId string) ApiListJobDeploymentHistoryRequest {
	return ApiListJobDeploymentHistoryRequest{
		ApiService: a,
		ctx:        ctx,
		jobId:      jobId,
	}
}

// Execute executes the request
//  @return ListJobDeploymentHistory200Response
func (a *JobDeploymentHistoryApiService) ListJobDeploymentHistoryExecute(r ApiListJobDeploymentHistoryRequest) (*ListJobDeploymentHistory200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ListJobDeploymentHistory200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "JobDeploymentHistoryApiService.ListJobDeploymentHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/job/{jobId}/deploymentHistory"
	localVarPath = strings.Replace(localVarPath, "{"+"jobId"+"}", url.PathEscape(parameterToString(r.jobId, "")), -1)

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
