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

// DatabaseMetricsApiService DatabaseMetricsApi service
type DatabaseMetricsApiService service

type ApiGetDatabaseCurrentMetricRequest struct {
	ctx        context.Context
	ApiService *DatabaseMetricsApiService
	databaseId string
}

func (r ApiGetDatabaseCurrentMetricRequest) Execute() (*DatabaseCurrentMetricResponse, *http.Response, error) {
	return r.ApiService.GetDatabaseCurrentMetricExecute(r)
}

/*
GetDatabaseCurrentMetric Get current metric consumption of the database

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseId Database ID
 @return ApiGetDatabaseCurrentMetricRequest
*/
func (a *DatabaseMetricsApiService) GetDatabaseCurrentMetric(ctx context.Context, databaseId string) ApiGetDatabaseCurrentMetricRequest {
	return ApiGetDatabaseCurrentMetricRequest{
		ApiService: a,
		ctx:        ctx,
		databaseId: databaseId,
	}
}

// Execute executes the request
//  @return DatabaseCurrentMetricResponse
func (a *DatabaseMetricsApiService) GetDatabaseCurrentMetricExecute(r ApiGetDatabaseCurrentMetricRequest) (*DatabaseCurrentMetricResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DatabaseCurrentMetricResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseMetricsApiService.GetDatabaseCurrentMetric")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/database/{databaseId}/currentMetric"
	localVarPath = strings.Replace(localVarPath, "{"+"databaseId"+"}", url.PathEscape(parameterToString(r.databaseId, "")), -1)

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

type ApiGetDatabaseMetricCpuRequest struct {
	ctx         context.Context
	ApiService  *DatabaseMetricsApiService
	databaseId  string
	lastSeconds *float32
}

// Up to how many seconds in the past to ask analytics results
func (r ApiGetDatabaseMetricCpuRequest) LastSeconds(lastSeconds float32) ApiGetDatabaseMetricCpuRequest {
	r.lastSeconds = &lastSeconds
	return r
}

func (r ApiGetDatabaseMetricCpuRequest) Execute() (*MetricCPUDatapointResponseList, *http.Response, error) {
	return r.ApiService.GetDatabaseMetricCpuExecute(r)
}

/*
GetDatabaseMetricCpu Get CPU consumption metric over time for the database

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseId Database ID
 @return ApiGetDatabaseMetricCpuRequest
*/
func (a *DatabaseMetricsApiService) GetDatabaseMetricCpu(ctx context.Context, databaseId string) ApiGetDatabaseMetricCpuRequest {
	return ApiGetDatabaseMetricCpuRequest{
		ApiService: a,
		ctx:        ctx,
		databaseId: databaseId,
	}
}

// Execute executes the request
//  @return MetricCPUDatapointResponseList
func (a *DatabaseMetricsApiService) GetDatabaseMetricCpuExecute(r ApiGetDatabaseMetricCpuRequest) (*MetricCPUDatapointResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MetricCPUDatapointResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseMetricsApiService.GetDatabaseMetricCpu")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/database/{databaseId}/metric/cpu"
	localVarPath = strings.Replace(localVarPath, "{"+"databaseId"+"}", url.PathEscape(parameterToString(r.databaseId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lastSeconds == nil {
		return localVarReturnValue, nil, reportError("lastSeconds is required and must be specified")
	}

	localVarQueryParams.Add("lastSeconds", parameterToString(*r.lastSeconds, ""))
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

type ApiGetDatabaseMetricHealthCheckRequest struct {
	ctx         context.Context
	ApiService  *DatabaseMetricsApiService
	databaseId  string
	lastSeconds *float32
}

// Up to how many seconds in the past to ask analytics results
func (r ApiGetDatabaseMetricHealthCheckRequest) LastSeconds(lastSeconds float32) ApiGetDatabaseMetricHealthCheckRequest {
	r.lastSeconds = &lastSeconds
	return r
}

func (r ApiGetDatabaseMetricHealthCheckRequest) Execute() (*MetricGenericResponseList, *http.Response, error) {
	return r.ApiService.GetDatabaseMetricHealthCheckExecute(r)
}

/*
GetDatabaseMetricHealthCheck Get Health Check latency  metric over time for the database

The value returned corresponds to the 95th centile

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseId Database ID
 @return ApiGetDatabaseMetricHealthCheckRequest
*/
func (a *DatabaseMetricsApiService) GetDatabaseMetricHealthCheck(ctx context.Context, databaseId string) ApiGetDatabaseMetricHealthCheckRequest {
	return ApiGetDatabaseMetricHealthCheckRequest{
		ApiService: a,
		ctx:        ctx,
		databaseId: databaseId,
	}
}

// Execute executes the request
//  @return MetricGenericResponseList
func (a *DatabaseMetricsApiService) GetDatabaseMetricHealthCheckExecute(r ApiGetDatabaseMetricHealthCheckRequest) (*MetricGenericResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MetricGenericResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseMetricsApiService.GetDatabaseMetricHealthCheck")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/database/{databaseId}/metric/healthCheck"
	localVarPath = strings.Replace(localVarPath, "{"+"databaseId"+"}", url.PathEscape(parameterToString(r.databaseId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lastSeconds == nil {
		return localVarReturnValue, nil, reportError("lastSeconds is required and must be specified")
	}

	localVarQueryParams.Add("lastSeconds", parameterToString(*r.lastSeconds, ""))
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

type ApiGetDatabaseMetricMemoryRequest struct {
	ctx         context.Context
	ApiService  *DatabaseMetricsApiService
	databaseId  string
	lastSeconds *float32
}

// Up to how many seconds in the past to ask analytics results
func (r ApiGetDatabaseMetricMemoryRequest) LastSeconds(lastSeconds float32) ApiGetDatabaseMetricMemoryRequest {
	r.lastSeconds = &lastSeconds
	return r
}

func (r ApiGetDatabaseMetricMemoryRequest) Execute() (*MetricMemoryDatapointResponseList, *http.Response, error) {
	return r.ApiService.GetDatabaseMetricMemoryExecute(r)
}

/*
GetDatabaseMetricMemory Get Memory consumption metric over time for the database

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseId Database ID
 @return ApiGetDatabaseMetricMemoryRequest
*/
func (a *DatabaseMetricsApiService) GetDatabaseMetricMemory(ctx context.Context, databaseId string) ApiGetDatabaseMetricMemoryRequest {
	return ApiGetDatabaseMetricMemoryRequest{
		ApiService: a,
		ctx:        ctx,
		databaseId: databaseId,
	}
}

// Execute executes the request
//  @return MetricMemoryDatapointResponseList
func (a *DatabaseMetricsApiService) GetDatabaseMetricMemoryExecute(r ApiGetDatabaseMetricMemoryRequest) (*MetricMemoryDatapointResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MetricMemoryDatapointResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseMetricsApiService.GetDatabaseMetricMemory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/database/{databaseId}/metric/memory"
	localVarPath = strings.Replace(localVarPath, "{"+"databaseId"+"}", url.PathEscape(parameterToString(r.databaseId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lastSeconds == nil {
		return localVarReturnValue, nil, reportError("lastSeconds is required and must be specified")
	}

	localVarQueryParams.Add("lastSeconds", parameterToString(*r.lastSeconds, ""))
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

type ApiGetDatabaseMetricRestartRequest struct {
	ctx         context.Context
	ApiService  *DatabaseMetricsApiService
	databaseId  string
	lastSeconds *float32
}

// Up to how many seconds in the past to ask analytics results
func (r ApiGetDatabaseMetricRestartRequest) LastSeconds(lastSeconds float32) ApiGetDatabaseMetricRestartRequest {
	r.lastSeconds = &lastSeconds
	return r
}

func (r ApiGetDatabaseMetricRestartRequest) Execute() (*MetricRestartResponse, *http.Response, error) {
	return r.ApiService.GetDatabaseMetricRestartExecute(r)
}

/*
GetDatabaseMetricRestart List database restarts

Get database restart message and timestamp.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseId Database ID
 @return ApiGetDatabaseMetricRestartRequest
*/
func (a *DatabaseMetricsApiService) GetDatabaseMetricRestart(ctx context.Context, databaseId string) ApiGetDatabaseMetricRestartRequest {
	return ApiGetDatabaseMetricRestartRequest{
		ApiService: a,
		ctx:        ctx,
		databaseId: databaseId,
	}
}

// Execute executes the request
//  @return MetricRestartResponse
func (a *DatabaseMetricsApiService) GetDatabaseMetricRestartExecute(r ApiGetDatabaseMetricRestartRequest) (*MetricRestartResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MetricRestartResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseMetricsApiService.GetDatabaseMetricRestart")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/database/{databaseId}/metric/restart"
	localVarPath = strings.Replace(localVarPath, "{"+"databaseId"+"}", url.PathEscape(parameterToString(r.databaseId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lastSeconds == nil {
		return localVarReturnValue, nil, reportError("lastSeconds is required and must be specified")
	}

	localVarQueryParams.Add("lastSeconds", parameterToString(*r.lastSeconds, ""))
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

type ApiGetDatabaseMetricStorageRequest struct {
	ctx         context.Context
	ApiService  *DatabaseMetricsApiService
	databaseId  string
	lastSeconds *float32
}

// Up to how many seconds in the past to ask analytics results
func (r ApiGetDatabaseMetricStorageRequest) LastSeconds(lastSeconds float32) ApiGetDatabaseMetricStorageRequest {
	r.lastSeconds = &lastSeconds
	return r
}

func (r ApiGetDatabaseMetricStorageRequest) Execute() (*MetricStorageDatapointResponseList, *http.Response, error) {
	return r.ApiService.GetDatabaseMetricStorageExecute(r)
}

/*
GetDatabaseMetricStorage Get Storage consumption metric over time for the database

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseId Database ID
 @return ApiGetDatabaseMetricStorageRequest
*/
func (a *DatabaseMetricsApiService) GetDatabaseMetricStorage(ctx context.Context, databaseId string) ApiGetDatabaseMetricStorageRequest {
	return ApiGetDatabaseMetricStorageRequest{
		ApiService: a,
		ctx:        ctx,
		databaseId: databaseId,
	}
}

// Execute executes the request
//  @return MetricStorageDatapointResponseList
func (a *DatabaseMetricsApiService) GetDatabaseMetricStorageExecute(r ApiGetDatabaseMetricStorageRequest) (*MetricStorageDatapointResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MetricStorageDatapointResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseMetricsApiService.GetDatabaseMetricStorage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/database/{databaseId}/metric/storage"
	localVarPath = strings.Replace(localVarPath, "{"+"databaseId"+"}", url.PathEscape(parameterToString(r.databaseId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.lastSeconds == nil {
		return localVarReturnValue, nil, reportError("lastSeconds is required and must be specified")
	}

	localVarQueryParams.Add("lastSeconds", parameterToString(*r.lastSeconds, ""))
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
