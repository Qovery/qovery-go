/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

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

// DatabasesApiService DatabasesApi service
type DatabasesApiService service

type ApiCreateDatabaseRequest struct {
	ctx             context.Context
	ApiService      *DatabasesApiService
	environmentId   string
	databaseRequest *DatabaseRequest
}

func (r ApiCreateDatabaseRequest) DatabaseRequest(databaseRequest DatabaseRequest) ApiCreateDatabaseRequest {
	r.databaseRequest = &databaseRequest
	return r
}

func (r ApiCreateDatabaseRequest) Execute() (*DatabaseResponse, *http.Response, error) {
	return r.ApiService.CreateDatabaseExecute(r)
}

/*
CreateDatabase Create a database

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @return ApiCreateDatabaseRequest
*/
func (a *DatabasesApiService) CreateDatabase(ctx context.Context, environmentId string) ApiCreateDatabaseRequest {
	return ApiCreateDatabaseRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return DatabaseResponse
func (a *DatabasesApiService) CreateDatabaseExecute(r ApiCreateDatabaseRequest) (*DatabaseResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DatabaseResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabasesApiService.CreateDatabase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/database"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentId"+"}", url.PathEscape(parameterToString(r.environmentId, "")), -1)

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
	localVarPostBody = r.databaseRequest
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

type ApiGetEnvironmentDatabaseStatusRequest struct {
	ctx           context.Context
	ApiService    *DatabasesApiService
	environmentId string
}

func (r ApiGetEnvironmentDatabaseStatusRequest) Execute() (*ReferenceObjectStatusResponseList, *http.Response, error) {
	return r.ApiService.GetEnvironmentDatabaseStatusExecute(r)
}

/*
GetEnvironmentDatabaseStatus List all environment databases statuses

Returns a list of databases with only their id and status.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @return ApiGetEnvironmentDatabaseStatusRequest
*/
func (a *DatabasesApiService) GetEnvironmentDatabaseStatus(ctx context.Context, environmentId string) ApiGetEnvironmentDatabaseStatusRequest {
	return ApiGetEnvironmentDatabaseStatusRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return ReferenceObjectStatusResponseList
func (a *DatabasesApiService) GetEnvironmentDatabaseStatusExecute(r ApiGetEnvironmentDatabaseStatusRequest) (*ReferenceObjectStatusResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ReferenceObjectStatusResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabasesApiService.GetEnvironmentDatabaseStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/database/status"
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

type ApiListDatabaseRequest struct {
	ctx           context.Context
	ApiService    *DatabasesApiService
	environmentId string
}

func (r ApiListDatabaseRequest) Execute() (*DatabaseResponseList, *http.Response, error) {
	return r.ApiService.ListDatabaseExecute(r)
}

/*
ListDatabase List environment databases

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @return ApiListDatabaseRequest
*/
func (a *DatabasesApiService) ListDatabase(ctx context.Context, environmentId string) ApiListDatabaseRequest {
	return ApiListDatabaseRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return DatabaseResponseList
func (a *DatabasesApiService) ListDatabaseExecute(r ApiListDatabaseRequest) (*DatabaseResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DatabaseResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabasesApiService.ListDatabase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/database"
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

type ApiListEnvironmentDatabaseConfigRequest struct {
	ctx           context.Context
	ApiService    *DatabasesApiService
	environmentId string
}

func (r ApiListEnvironmentDatabaseConfigRequest) Execute() (*DatabaseConfigurationResponseList, *http.Response, error) {
	return r.ApiService.ListEnvironmentDatabaseConfigExecute(r)
}

/*
ListEnvironmentDatabaseConfig List eligible database types, versions and modes for the environment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @return ApiListEnvironmentDatabaseConfigRequest
*/
func (a *DatabasesApiService) ListEnvironmentDatabaseConfig(ctx context.Context, environmentId string) ApiListEnvironmentDatabaseConfigRequest {
	return ApiListEnvironmentDatabaseConfigRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return DatabaseConfigurationResponseList
func (a *DatabasesApiService) ListEnvironmentDatabaseConfigExecute(r ApiListEnvironmentDatabaseConfigRequest) (*DatabaseConfigurationResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DatabaseConfigurationResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabasesApiService.ListEnvironmentDatabaseConfig")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/databaseConfiguration"
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

type ApiListEnvironmentDatabaseCurrentMetricRequest struct {
	ctx           context.Context
	ApiService    *DatabasesApiService
	environmentId string
}

func (r ApiListEnvironmentDatabaseCurrentMetricRequest) Execute() (*EnvironmentDatabasesCurrentMetricResponseList, *http.Response, error) {
	return r.ApiService.ListEnvironmentDatabaseCurrentMetricExecute(r)
}

/*
ListEnvironmentDatabaseCurrentMetric List current metric consumption for each database

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentId Environment ID
 @return ApiListEnvironmentDatabaseCurrentMetricRequest
*/
func (a *DatabasesApiService) ListEnvironmentDatabaseCurrentMetric(ctx context.Context, environmentId string) ApiListEnvironmentDatabaseCurrentMetricRequest {
	return ApiListEnvironmentDatabaseCurrentMetricRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentId: environmentId,
	}
}

// Execute executes the request
//  @return EnvironmentDatabasesCurrentMetricResponseList
func (a *DatabasesApiService) ListEnvironmentDatabaseCurrentMetricExecute(r ApiListEnvironmentDatabaseCurrentMetricRequest) (*EnvironmentDatabasesCurrentMetricResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EnvironmentDatabasesCurrentMetricResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabasesApiService.ListEnvironmentDatabaseCurrentMetric")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environment/{environmentId}/database/currentMetric"
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
