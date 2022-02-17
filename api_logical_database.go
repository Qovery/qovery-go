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

// LogicalDatabaseApiService LogicalDatabaseApi service
type LogicalDatabaseApiService service

type ApiCreateLogicalDatabaseOnDatabaseRequest struct {
	ctx                    context.Context
	ApiService             *LogicalDatabaseApiService
	databaseId             string
	logicalDatabaseRequest *LogicalDatabaseRequest
}

func (r ApiCreateLogicalDatabaseOnDatabaseRequest) LogicalDatabaseRequest(logicalDatabaseRequest LogicalDatabaseRequest) ApiCreateLogicalDatabaseOnDatabaseRequest {
	r.logicalDatabaseRequest = &logicalDatabaseRequest
	return r
}

func (r ApiCreateLogicalDatabaseOnDatabaseRequest) Execute() (*LogicalDatabaseResponse, *http.Response, error) {
	return r.ApiService.CreateLogicalDatabaseOnDatabaseExecute(r)
}

/*
CreateLogicalDatabaseOnDatabase Create a logical database on the database

If you don't specify credentials, Qovery will autogenerate them.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseId Database ID
 @return ApiCreateLogicalDatabaseOnDatabaseRequest
*/
func (a *LogicalDatabaseApiService) CreateLogicalDatabaseOnDatabase(ctx context.Context, databaseId string) ApiCreateLogicalDatabaseOnDatabaseRequest {
	return ApiCreateLogicalDatabaseOnDatabaseRequest{
		ApiService: a,
		ctx:        ctx,
		databaseId: databaseId,
	}
}

// Execute executes the request
//  @return LogicalDatabaseResponse
func (a *LogicalDatabaseApiService) CreateLogicalDatabaseOnDatabaseExecute(r ApiCreateLogicalDatabaseOnDatabaseRequest) (*LogicalDatabaseResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LogicalDatabaseResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogicalDatabaseApiService.CreateLogicalDatabaseOnDatabase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/database/{databaseId}/logicalDatabase"
	localVarPath = strings.Replace(localVarPath, "{"+"databaseId"+"}", url.PathEscape(parameterToString(r.databaseId, "")), -1)

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
	localVarPostBody = r.logicalDatabaseRequest
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

type ApiDeleteLogicalDatabaseRequest struct {
	ctx               context.Context
	ApiService        *LogicalDatabaseApiService
	logicalDatabaseId string
}

func (r ApiDeleteLogicalDatabaseRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteLogicalDatabaseExecute(r)
}

/*
DeleteLogicalDatabase Delete a Logical database

To delete a logical database you must have the project user permission

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param logicalDatabaseId Logical Database ID
 @return ApiDeleteLogicalDatabaseRequest
*/
func (a *LogicalDatabaseApiService) DeleteLogicalDatabase(ctx context.Context, logicalDatabaseId string) ApiDeleteLogicalDatabaseRequest {
	return ApiDeleteLogicalDatabaseRequest{
		ApiService:        a,
		ctx:               ctx,
		logicalDatabaseId: logicalDatabaseId,
	}
}

// Execute executes the request
func (a *LogicalDatabaseApiService) DeleteLogicalDatabaseExecute(r ApiDeleteLogicalDatabaseRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogicalDatabaseApiService.DeleteLogicalDatabase")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/logicalDatabase/{logicalDatabaseId}"
	localVarPath = strings.Replace(localVarPath, "{"+"logicalDatabaseId"+"}", url.PathEscape(parameterToString(r.logicalDatabaseId, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiEditLogicalDatabaseRequest struct {
	ctx                    context.Context
	ApiService             *LogicalDatabaseApiService
	logicalDatabaseId      string
	logicalDatabaseRequest *LogicalDatabaseRequest
}

func (r ApiEditLogicalDatabaseRequest) LogicalDatabaseRequest(logicalDatabaseRequest LogicalDatabaseRequest) ApiEditLogicalDatabaseRequest {
	r.logicalDatabaseRequest = &logicalDatabaseRequest
	return r
}

func (r ApiEditLogicalDatabaseRequest) Execute() (*LogicalDatabaseResponse, *http.Response, error) {
	return r.ApiService.EditLogicalDatabaseExecute(r)
}

/*
EditLogicalDatabase Edit a logical database

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param logicalDatabaseId Logical Database ID
 @return ApiEditLogicalDatabaseRequest
*/
func (a *LogicalDatabaseApiService) EditLogicalDatabase(ctx context.Context, logicalDatabaseId string) ApiEditLogicalDatabaseRequest {
	return ApiEditLogicalDatabaseRequest{
		ApiService:        a,
		ctx:               ctx,
		logicalDatabaseId: logicalDatabaseId,
	}
}

// Execute executes the request
//  @return LogicalDatabaseResponse
func (a *LogicalDatabaseApiService) EditLogicalDatabaseExecute(r ApiEditLogicalDatabaseRequest) (*LogicalDatabaseResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LogicalDatabaseResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogicalDatabaseApiService.EditLogicalDatabase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/logicalDatabase/{logicalDatabaseId}"
	localVarPath = strings.Replace(localVarPath, "{"+"logicalDatabaseId"+"}", url.PathEscape(parameterToString(r.logicalDatabaseId, "")), -1)

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
	localVarPostBody = r.logicalDatabaseRequest
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

type ApiEditLogicalDatabaseCredentialsRequest struct {
	ctx                context.Context
	ApiService         *LogicalDatabaseApiService
	logicalDatabaseId  string
	credentialsRequest *CredentialsRequest
}

func (r ApiEditLogicalDatabaseCredentialsRequest) CredentialsRequest(credentialsRequest CredentialsRequest) ApiEditLogicalDatabaseCredentialsRequest {
	r.credentialsRequest = &credentialsRequest
	return r
}

func (r ApiEditLogicalDatabaseCredentialsRequest) Execute() (*CredentialsResponse, *http.Response, error) {
	return r.ApiService.EditLogicalDatabaseCredentialsExecute(r)
}

/*
EditLogicalDatabaseCredentials Edit logical database credentials

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param logicalDatabaseId Logical Database ID
 @return ApiEditLogicalDatabaseCredentialsRequest
*/
func (a *LogicalDatabaseApiService) EditLogicalDatabaseCredentials(ctx context.Context, logicalDatabaseId string) ApiEditLogicalDatabaseCredentialsRequest {
	return ApiEditLogicalDatabaseCredentialsRequest{
		ApiService:        a,
		ctx:               ctx,
		logicalDatabaseId: logicalDatabaseId,
	}
}

// Execute executes the request
//  @return CredentialsResponse
func (a *LogicalDatabaseApiService) EditLogicalDatabaseCredentialsExecute(r ApiEditLogicalDatabaseCredentialsRequest) (*CredentialsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CredentialsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogicalDatabaseApiService.EditLogicalDatabaseCredentials")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/logicalDatabase/{logicalDatabaseId}/credentials"
	localVarPath = strings.Replace(localVarPath, "{"+"logicalDatabaseId"+"}", url.PathEscape(parameterToString(r.logicalDatabaseId, "")), -1)

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
	localVarPostBody = r.credentialsRequest
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

type ApiGetLogicalDatabaseRequest struct {
	ctx               context.Context
	ApiService        *LogicalDatabaseApiService
	logicalDatabaseId string
}

func (r ApiGetLogicalDatabaseRequest) Execute() (*LogicalDatabaseResponse, *http.Response, error) {
	return r.ApiService.GetLogicalDatabaseExecute(r)
}

/*
GetLogicalDatabase Get logical database by ID

A logical database exists inside a database. The database is a service that exists within an environment, that you can deploy, and that has allocated resources. A database can have several logical databases

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param logicalDatabaseId Logical Database ID
 @return ApiGetLogicalDatabaseRequest
*/
func (a *LogicalDatabaseApiService) GetLogicalDatabase(ctx context.Context, logicalDatabaseId string) ApiGetLogicalDatabaseRequest {
	return ApiGetLogicalDatabaseRequest{
		ApiService:        a,
		ctx:               ctx,
		logicalDatabaseId: logicalDatabaseId,
	}
}

// Execute executes the request
//  @return LogicalDatabaseResponse
func (a *LogicalDatabaseApiService) GetLogicalDatabaseExecute(r ApiGetLogicalDatabaseRequest) (*LogicalDatabaseResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LogicalDatabaseResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogicalDatabaseApiService.GetLogicalDatabase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/logicalDatabase/{logicalDatabaseId}"
	localVarPath = strings.Replace(localVarPath, "{"+"logicalDatabaseId"+"}", url.PathEscape(parameterToString(r.logicalDatabaseId, "")), -1)

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

type ApiGetLogicalDatabaseCredentialsRequest struct {
	ctx               context.Context
	ApiService        *LogicalDatabaseApiService
	logicalDatabaseId string
}

func (r ApiGetLogicalDatabaseCredentialsRequest) Execute() (*CredentialsResponse, *http.Response, error) {
	return r.ApiService.GetLogicalDatabaseCredentialsExecute(r)
}

/*
GetLogicalDatabaseCredentials Get  credentials of the logical database

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param logicalDatabaseId Logical Database ID
 @return ApiGetLogicalDatabaseCredentialsRequest
*/
func (a *LogicalDatabaseApiService) GetLogicalDatabaseCredentials(ctx context.Context, logicalDatabaseId string) ApiGetLogicalDatabaseCredentialsRequest {
	return ApiGetLogicalDatabaseCredentialsRequest{
		ApiService:        a,
		ctx:               ctx,
		logicalDatabaseId: logicalDatabaseId,
	}
}

// Execute executes the request
//  @return CredentialsResponse
func (a *LogicalDatabaseApiService) GetLogicalDatabaseCredentialsExecute(r ApiGetLogicalDatabaseCredentialsRequest) (*CredentialsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CredentialsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogicalDatabaseApiService.GetLogicalDatabaseCredentials")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/logicalDatabase/{logicalDatabaseId}/credentials"
	localVarPath = strings.Replace(localVarPath, "{"+"logicalDatabaseId"+"}", url.PathEscape(parameterToString(r.logicalDatabaseId, "")), -1)

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

type ApiListLogicalDatabaseApplicationRequest struct {
	ctx               context.Context
	ApiService        *LogicalDatabaseApiService
	logicalDatabaseId string
}

func (r ApiListLogicalDatabaseApplicationRequest) Execute() (*ApplicationResponseList, *http.Response, error) {
	return r.ApiService.ListLogicalDatabaseApplicationExecute(r)
}

/*
ListLogicalDatabaseApplication List linked applications

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param logicalDatabaseId Logical Database ID
 @return ApiListLogicalDatabaseApplicationRequest
*/
func (a *LogicalDatabaseApiService) ListLogicalDatabaseApplication(ctx context.Context, logicalDatabaseId string) ApiListLogicalDatabaseApplicationRequest {
	return ApiListLogicalDatabaseApplicationRequest{
		ApiService:        a,
		ctx:               ctx,
		logicalDatabaseId: logicalDatabaseId,
	}
}

// Execute executes the request
//  @return ApplicationResponseList
func (a *LogicalDatabaseApiService) ListLogicalDatabaseApplicationExecute(r ApiListLogicalDatabaseApplicationRequest) (*ApplicationResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ApplicationResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogicalDatabaseApiService.ListLogicalDatabaseApplication")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/logicalDatabase/{logicalDatabaseId}/application"
	localVarPath = strings.Replace(localVarPath, "{"+"logicalDatabaseId"+"}", url.PathEscape(parameterToString(r.logicalDatabaseId, "")), -1)

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

type ApiListLogicalDatabaseDatabaseRequest struct {
	ctx        context.Context
	ApiService *LogicalDatabaseApiService
	databaseId string
}

func (r ApiListLogicalDatabaseDatabaseRequest) Execute() (*LogicalDatabaseResponseList, *http.Response, error) {
	return r.ApiService.ListLogicalDatabaseDatabaseExecute(r)
}

/*
ListLogicalDatabaseDatabase List logical databases of a database

A logical database exists inside a database. The database is a service that exists within an environment, that you can deploy, and that has allocated resources. A database can have several logical databases

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseId Database ID
 @return ApiListLogicalDatabaseDatabaseRequest
*/
func (a *LogicalDatabaseApiService) ListLogicalDatabaseDatabase(ctx context.Context, databaseId string) ApiListLogicalDatabaseDatabaseRequest {
	return ApiListLogicalDatabaseDatabaseRequest{
		ApiService: a,
		ctx:        ctx,
		databaseId: databaseId,
	}
}

// Execute executes the request
//  @return LogicalDatabaseResponseList
func (a *LogicalDatabaseApiService) ListLogicalDatabaseDatabaseExecute(r ApiListLogicalDatabaseDatabaseRequest) (*LogicalDatabaseResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *LogicalDatabaseResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogicalDatabaseApiService.ListLogicalDatabaseDatabase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/database/{databaseId}/logicalDatabase"
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
