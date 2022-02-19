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

// Linger please
var (
	_ context.Context
)

// DatabaseMainCallsApiService DatabaseMainCallsApi service
type DatabaseMainCallsApiService service

type ApiDeleteDatabaseRequest struct {
	ctx        context.Context
	ApiService *DatabaseMainCallsApiService
	databaseId string
}

func (r ApiDeleteDatabaseRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteDatabaseExecute(r)
}

/*
DeleteDatabase Delete a database

To delete a database you must have the admin permission

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseId Database ID
 @return ApiDeleteDatabaseRequest
*/
func (a *DatabaseMainCallsApiService) DeleteDatabase(ctx context.Context, databaseId string) ApiDeleteDatabaseRequest {
	return ApiDeleteDatabaseRequest{
		ApiService: a,
		ctx:        ctx,
		databaseId: databaseId,
	}
}

// Execute executes the request
func (a *DatabaseMainCallsApiService) DeleteDatabaseExecute(r ApiDeleteDatabaseRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseMainCallsApiService.DeleteDatabase")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/database/{databaseId}"
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

type ApiEditDatabaseRequest struct {
	ctx                 context.Context
	ApiService          *DatabaseMainCallsApiService
	databaseId          string
	databaseEditRequest *DatabaseEditRequest
}

func (r ApiEditDatabaseRequest) DatabaseEditRequest(databaseEditRequest DatabaseEditRequest) ApiEditDatabaseRequest {
	r.databaseEditRequest = &databaseEditRequest
	return r
}

func (r ApiEditDatabaseRequest) Execute() (*DatabaseResponse, *http.Response, error) {
	return r.ApiService.EditDatabaseExecute(r)
}

/*
EditDatabase Edit a database

To edit a database  you must have the admin permission

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseId Database ID
 @return ApiEditDatabaseRequest
*/
func (a *DatabaseMainCallsApiService) EditDatabase(ctx context.Context, databaseId string) ApiEditDatabaseRequest {
	return ApiEditDatabaseRequest{
		ApiService: a,
		ctx:        ctx,
		databaseId: databaseId,
	}
}

// Execute executes the request
//  @return DatabaseResponse
func (a *DatabaseMainCallsApiService) EditDatabaseExecute(r ApiEditDatabaseRequest) (*DatabaseResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DatabaseResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseMainCallsApiService.EditDatabase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/database/{databaseId}"
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
	localVarPostBody = r.databaseEditRequest
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

type ApiEditDatabaseCredentialsRequest struct {
	ctx                context.Context
	ApiService         *DatabaseMainCallsApiService
	databaseId         string
	credentialsRequest *CredentialsRequest
}

func (r ApiEditDatabaseCredentialsRequest) CredentialsRequest(credentialsRequest CredentialsRequest) ApiEditDatabaseCredentialsRequest {
	r.credentialsRequest = &credentialsRequest
	return r
}

func (r ApiEditDatabaseCredentialsRequest) Execute() (*CredentialsResponse, *http.Response, error) {
	return r.ApiService.EditDatabaseCredentialsExecute(r)
}

/*
EditDatabaseCredentials Edit database  master credentials

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseId Database ID
 @return ApiEditDatabaseCredentialsRequest
*/
func (a *DatabaseMainCallsApiService) EditDatabaseCredentials(ctx context.Context, databaseId string) ApiEditDatabaseCredentialsRequest {
	return ApiEditDatabaseCredentialsRequest{
		ApiService: a,
		ctx:        ctx,
		databaseId: databaseId,
	}
}

// Execute executes the request
//  @return CredentialsResponse
func (a *DatabaseMainCallsApiService) EditDatabaseCredentialsExecute(r ApiEditDatabaseCredentialsRequest) (*CredentialsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CredentialsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseMainCallsApiService.EditDatabaseCredentials")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/database/{databaseId}/masterCredentials"
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

type ApiGetDatabaseRequest struct {
	ctx        context.Context
	ApiService *DatabaseMainCallsApiService
	databaseId string
}

func (r ApiGetDatabaseRequest) Execute() (*DatabaseResponse, *http.Response, error) {
	return r.ApiService.GetDatabaseExecute(r)
}

/*
GetDatabase Get database by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseId Database ID
 @return ApiGetDatabaseRequest
*/
func (a *DatabaseMainCallsApiService) GetDatabase(ctx context.Context, databaseId string) ApiGetDatabaseRequest {
	return ApiGetDatabaseRequest{
		ApiService: a,
		ctx:        ctx,
		databaseId: databaseId,
	}
}

// Execute executes the request
//  @return DatabaseResponse
func (a *DatabaseMainCallsApiService) GetDatabaseExecute(r ApiGetDatabaseRequest) (*DatabaseResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DatabaseResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseMainCallsApiService.GetDatabase")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/database/{databaseId}"
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

type ApiGetDatabaseMasterCredentialsRequest struct {
	ctx        context.Context
	ApiService *DatabaseMainCallsApiService
	databaseId string
}

func (r ApiGetDatabaseMasterCredentialsRequest) Execute() (*CredentialsResponse, *http.Response, error) {
	return r.ApiService.GetDatabaseMasterCredentialsExecute(r)
}

/*
GetDatabaseMasterCredentials Get master credentials of the database

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseId Database ID
 @return ApiGetDatabaseMasterCredentialsRequest
*/
func (a *DatabaseMainCallsApiService) GetDatabaseMasterCredentials(ctx context.Context, databaseId string) ApiGetDatabaseMasterCredentialsRequest {
	return ApiGetDatabaseMasterCredentialsRequest{
		ApiService: a,
		ctx:        ctx,
		databaseId: databaseId,
	}
}

// Execute executes the request
//  @return CredentialsResponse
func (a *DatabaseMainCallsApiService) GetDatabaseMasterCredentialsExecute(r ApiGetDatabaseMasterCredentialsRequest) (*CredentialsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CredentialsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseMainCallsApiService.GetDatabaseMasterCredentials")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/database/{databaseId}/masterCredentials"
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

type ApiGetDatabaseStatusRequest struct {
	ctx        context.Context
	ApiService *DatabaseMainCallsApiService
	databaseId string
}

func (r ApiGetDatabaseStatusRequest) Execute() (*Status, *http.Response, error) {
	return r.ApiService.GetDatabaseStatusExecute(r)
}

/*
GetDatabaseStatus Get database status

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseId Database ID
 @return ApiGetDatabaseStatusRequest
*/
func (a *DatabaseMainCallsApiService) GetDatabaseStatus(ctx context.Context, databaseId string) ApiGetDatabaseStatusRequest {
	return ApiGetDatabaseStatusRequest{
		ApiService: a,
		ctx:        ctx,
		databaseId: databaseId,
	}
}

// Execute executes the request
//  @return Status
func (a *DatabaseMainCallsApiService) GetDatabaseStatusExecute(r ApiGetDatabaseStatusRequest) (*Status, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Status
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseMainCallsApiService.GetDatabaseStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/database/{databaseId}/status"
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

type ApiListDatabaseVersionRequest struct {
	ctx        context.Context
	ApiService *DatabaseMainCallsApiService
	databaseId string
}

func (r ApiListDatabaseVersionRequest) Execute() (*VersionResponseList, *http.Response, error) {
	return r.ApiService.ListDatabaseVersionExecute(r)
}

/*
ListDatabaseVersion List eligible versions for the database

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param databaseId Database ID
 @return ApiListDatabaseVersionRequest
*/
func (a *DatabaseMainCallsApiService) ListDatabaseVersion(ctx context.Context, databaseId string) ApiListDatabaseVersionRequest {
	return ApiListDatabaseVersionRequest{
		ApiService: a,
		ctx:        ctx,
		databaseId: databaseId,
	}
}

// Execute executes the request
//  @return VersionResponseList
func (a *DatabaseMainCallsApiService) ListDatabaseVersionExecute(r ApiListDatabaseVersionRequest) (*VersionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VersionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DatabaseMainCallsApiService.ListDatabaseVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/database/{databaseId}/version"
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
